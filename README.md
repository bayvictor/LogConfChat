# LogConfChat
A chat server in golang, which log all activities for convenience for later analytics,   and with config file boost.
LogConfChat architecture: -- A chatserver wich log allactivities for convenience for later analytics,
                        and config file support in golang.
<pre>
1. Architecture<br>

├── cmd_chat_client.go        # Chat with vibriant GUI-like text consule UI<br>
├── cmd_chat_server_7777.log  # log for exec "cmd_chat_server" when running on port# 7777<br>
├── cmd_chat_server.go        #<br>
├── config                    # persistence config like port#, when absent or error, using cmd os.Args[] default.<br>
├── live_test.sh              # minute-to-minute tweak script, test out all features.<br>
├── pkg                       # when to "go get xxxx" external libraries put.<br>
├── run_me_once.sh            # init scripts, before any scripts run.<br>
├── src                       #  <br>
├── test.results.log.txt      # worksheets for development<br>
└── write_config_file.go      # which reveals how config structs<br>
</pre>Where                      
2 open source golang libraries gocui, gonfig are used.
=========================
how to use:
##step 1. git clone https://www.github.com/LogConfChat.git 
##Step 2. cd  LogConfChat; ./run_me_once.sh
##step 3. run servers. 
      ## you can run multiple chatserver but on different port number.
./cmd_chat_server 2345
./cmd_chat_server 2346 
...
./cmd_chat_server  ##which default to "5000"
 server init order: check confile file first, then os.Args[], if all missing or fail using default "5000" port.


## Step 4. running client: connecting to right hostname, port#, with flexible initialization.

./cmd_chat_client localhost 2346 
./cmd_chat_client some_where_over_the_rainbow 3456 
./cmd_chat_client remote_host 5678 
./cmd_chat_client  ##which default to "localhost" "5000"
The above server as it's name pointing out, can log all activities, can read/write config file using gonfig.

Licenses. like all its dependency here, this code and it's future change here, also totally free, but provided as it is and deny all liabilities.

================================================================================


Test Results,  snapshots:

One of the client "victor" view, during "daniel" and "lambda" come and go!
<pre>
┌─ messages: ───────────────────────────────────────────────┌─ 2 users: ───────┐<br>
│[10:49:48] + daniel connected                              │victor            │<br>
│[10:49:55] daniel: i am the saint                          │victor            │<br>
│[10:50:02] daniel: daniel the one                          │                  │<br>
│[10:50:09] daniel: young winner.                           │                  │<br>
│[11:09:53] + victor connected                              │                  │<br>
│[11:09:54] victor: skdfjska                                │                  │<br>
│[11:09:55] victor: sdfksalk;                               │                  │<br>
│[11:09:57] victor: 12342143                                │                  │<br>
│[11:09:59] victor: 34555566                                │                  │<br>
│[11:10:00] victor: 445                                     │                  │<br>
│[11:10:02] victor: hello                                   │                  │<br>
│[11:10:31] + lambda connected                              │                  │<br>
│[11:10:35] lambda: kaka                                    │                  │<br>
│[11:10:37] lambda: yaya                                    │                  │<br>
│[11:10:40] lambda: never!                                  │                  │<br>
│[11:10:45] - lambda disconnected                           │                  │<br>
│[11:17:32] - daniel disconnected                           │                  │<br>
│                                                           │                  │<br>
┌─ send: ───────────────────────────────────────────────────┐                  │<br>
│                                                           │                  │<br>
│                                                           │                  │<br>
│                                                           │                  │<br>
└───────────────────────────────────────────────────────────┘──────────────────┘<br>

</pre>
Sample server log of above events:
vhuang@dell7710:/sdb3/src/golang/victorchat$ cat victor.log 
+ victor connected
2020/09/13 10:48:08 + victor connected
victor: ye,ye, n\ya, yum,yum
2020/09/13 10:48:22 victor: ye,ye, n\ya, yum,yum
victor: hello
2020/09/13 10:48:24 victor: hello
+ daniel connected
2020/09/13 10:49:48 + daniel connected
...
- daniel disconnected
2020/09/13 11:17:32 - daniel disconnected

/==================================================================
Client interface:
full command line usage is "cmd_chat_client <host> <port>", with argv[1],default  is "cmd_chat_client localhost 5555".

===================================================================
Case:
when multiple server starts in localhost, they log differently. filename: "os.Args[0]+"_"+os.Args[1]+".log", basically server-exec name+ portnumber which it running on.


===================================================================
manually check config file in "config/config.development.json", (because ENV default is development).
change port to one never used before "5678", 
then go to batch test file run2.sh, change all command line port to "5678", find it read out and use ok.
If remove "config/" dir then run, all fails because without configfile, server port default to 5000 without argument. 
found it behave correctly.
