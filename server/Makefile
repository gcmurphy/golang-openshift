include $(GOROOT)/src/Make.inc

TARG=run
GOFILES=\
	main.go\

include $(GOROOT)/src/Make.cmd

deploy: all
	mv $(TARG) ../bin
	make clean
	git commit -a -m "deploy"  
	git push

