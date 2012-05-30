How to Run Go on OpenShift
==========================

OpenShift is a great PaaS offering by Red Hat. At the moment 
Go is not currently on the supported platform list. However Red Hat 
have kindly provided a DIY cartridge to run whatever you want on the 
platform. As Go is one of my favourite languages I thought I'd put 
together a template that can be used to deploy a basic Go application. 

0. Create a OpenShift account. Go to http://openshift.redhat.com to create an account.
    
1. Install the OpenShift client tools. 
      
    cd /etc/yum.repos.d  
    sudo wget https://openshift.redhat.com/app/repo/openshift.repo  
    sudo yum install rhc

2. Create a new domain

    rhc-create-domain --namespace *your-namespace* --rhlogin *your-login*

3. Create a new application 

    rhc-create-app --app *your-appname* --type diy-0.1 --rhlogin *your-login*

4. Add this as an upstream repository

    git remote add upstream -m master git://github.com/gcmurphy/golang-openshift.git  
    git pull -s recursive -X theirs upstream master

5. Build the source and copy the output to your bin directory (updated for Go version 1)
    
    cd server  
    go build   
    mv server ../bin  

5. Push to your OpenShift repository
   
    git commit -a -m "My first go + openshift application"   
    git push


6. You are good to Go! You can read more about the OpenShift DIY cartridge [here](https://www.redhat.com/openshift/community/blogs/a-paas-that-runs-anything-http-getting-started-with-diy-applications-on-openshift). 

A note on cross compiling
-------------------------

This skeleton project depends on the fact that the OpenShift images that 
are deployed are running on Linux x86_64 virtual machines. However the 
OpenShift toolkit is available for both Windows and Mac OSX. In order 
to use this toolkit on these platforms the following additional steps need 
to be used to create a cross compiler for the non Linux development environments: 


1. Get the latest release of the Go source code. 

    wget http://go.googlecode.com/files/go1.0.1.src.tar.gz

2. Extract the source code and compile it, targeting the Linux platform and architecture
    
    tar xzvf go*.tar.gz  
    cd go/src  
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ./make.bash  

3. If everything worked you should see:

    Installed Go for linux/amd64 in _Your build directory_   
    Installed commands in _Your build directory/bin_
    
4. Try to build the project now using the go compiler we just created 

    cd golang-openshift/server  
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ~/xcompile/go/bin/go build  
    
The output file will run on a Linux AMD64 box but not locally which is 
a bit of an annoyance. The alternative would be to create a copy of the 
Go runtime on the Linux machine and compile the code remotely. 