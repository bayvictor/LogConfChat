#!/bin/bash -x 
#mkdir -p  ~/victorchat 
#cd ~/src/victorchat
export GOPATH=$PWD
go get "github.com/tkanos/gonfig"
go get "github.com/jroimartin/gocui"                                                                                                                 
go run write_config_file.go & # generate pesistence port# , to read fram file 
go run cmd_chat_server.go &  
go run cmd_chat_client.go ## you can in diff window run muitiple clients 


ls 
