How to Run Go on OpenShift
==========================

OpenShift is a great PaaS offering by Red Hat. At the moment 
Go is not currently on the supported platform list. However Red Hat 
have kindly provided a DIY cartridge to run whatever you want on the 
platform. As Go is one of my favourite languages I thought I'd put 
together a template that can be used to deploy a basic Go application. 

0. Create a OpenShift account 
    * Go to http://openshift.redhat.com to create an account
    
1. Install the OpenShift client tools. 

    cd /etc/yum.repos.d
    sudo wget https://openshift.redhat.com/app/repo/openshift.repo
    sudo yum install rhc

2. Create a new domain

    rhc-create-domain --namespace <your-namespace> --rhlogin <your-login>`

3. Create a new application 

    rhc-create-app --app <your-appname> --type diy-0.1 --rhlogin <your-login>

4. Add this as an upstream repository

    git remote add upstream -m master git://github.com/gcmurphy/golang-openshift.git

5. Push to your OpenShift repository

    git push

6. You are good to Go! You can read more about the OpenShift DIY cartridge [here](https://www.redhat.com/openshift/community/blogs/a-paas-that-runs-anything-http-getting-started-with-diy-applications-on-openshift). For 
conveinence sake you can just run: 

    make clean
    make deploy

To push it out changes to the cloud.
