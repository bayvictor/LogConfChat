export GOPATH=$PWD
#cd ~/src/golang/chatbot
go get "github.com/tkanos/gonfig"
go get "github.com/jroimartin/gocui" 
go run  write_config_file.go # can be comment out to watch no-config hehavior. 
go run cmd_chat_server.go 5678 & 
go run cmd_chat_client.go localhost 5678 
  
