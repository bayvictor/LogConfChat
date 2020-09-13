package main
//server2  
import (
        "bufio"
        "fmt"
        "net"
        "strings"
        "log"
        "time"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	//Port              int `env:"APP_PORT"`
	Port      string
	Remote_Host_Name string
}


func getFileName() string {
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "development"
	}
	filename := []string{"config/", "config.", env, ".json"}
	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), strings.Join(filename, ""))

	return filePath
}
func config_test_main() {

	// mock env variable
	os.Setenv("APP_PORT", "5000")

	configuration := Configuration{}
	err := gonfig.GetConf(getFileName(), &configuration)
	if err != nil {
		fmt.Println(err)
		os.Exit(500)
	}

	fmt.Println(configuration.Port)

}



// Client struct
type Client struct {
        name       string
        conn       net.Conn
        writer     *bufio.Writer
        reader     *bufio.Reader
        incoming   chan string
        outgoing   chan string
        disconnect chan bool
        status     int // 1 connected, 0 otherwise
}

// CreateClient creates new client and starts listening
// for incoming and outgoing messages.
func CreateClient(conn net.Conn) *Client {
        writer := bufio.NewWriter(conn)
        reader := bufio.NewReader(conn)

        client := &Client{
                name:       "user",
                conn:       conn,
                writer:     writer,
                outgoing:   make(chan string),
                reader:     reader,
                incoming:   make(chan string),
                disconnect: make(chan bool),
                status:     1,
        }

        go client.Write()
        go client.Read()

        return client
}

// Write writes message to the client.
func (client *Client) Write() {
        for {
                select {
                case <-client.disconnect:
                        client.status = 0
                        break
                default:
                        msg := <-client.outgoing
                        client.writer.WriteString(msg)
                        client.writer.Flush()
                }
        }
}

// Read reads message from the client.
func (client *Client) Read() {
        for {
                msg, err := client.reader.ReadString('\n')
                if err != nil {
                        client.incoming <- fmt.Sprintf("\x1b[0;31m- %s disconnected\033[0m\n", client.name)
                        client.status = 0
                        client.disconnect <- true
                        client.conn.Close()
                        break
                }
                switch {
                case strings.HasPrefix(msg, "/name>"):
                        name := strings.TrimSpace(strings.SplitAfter(msg, ">")[1])
                        client.name = name
                        client.incoming <- fmt.Sprintf("\x1b[0;32m+ %s connected\033[0m\n", name)
                default:
                        client.incoming <- fmt.Sprintf("%s: %s", client.name, msg)
                }
        }
}

// Chat struct
type Chat struct {
        clients  []*Client
        connect  chan net.Conn
        outgoing chan string
}

// CreateChat creates new chat and starts listening for connections.
func CreateChat() *Chat {
        chat := &Chat{
                clients:  make([]*Client, 0),
                connect:  make(chan net.Conn),
                outgoing: make(chan string),
        }

        chat.Listen()

        return chat
}

func log_mes_to_file(filename string, prefix string, mes string){

f, err := os.OpenFile(filename+".log",
        os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
        log.Println(err)
}
defer f.Close()

logger := log.New(f, mes, log.LstdFlags)
logger.Println(mes)
//logger.Println("mod text to append")


}


// Listen listens for connections and messages to broadcast.
func (chat *Chat) Listen() {
        go func() {
                for {
                        select {
                        case conn := <-chat.connect:
                                chat.Join(conn)
                        case msg := <-chat.outgoing:
                                chat.Broadcast(msg)
        log_mes_to_file( os.Args[0] + "_" + os.Args[1], "prefix:", msg ) //# run without diff port have diff logfile
      }
                }
        }()
}

// Connect passing connection to the chat.
func (chat *Chat) Connect(conn net.Conn) {
        chat.connect <- conn
}

// Join creates new client and starts listening for client messages.
func (chat *Chat) Join(conn net.Conn) {
        client := CreateClient(conn)
        chat.clients = append(chat.clients, client)
        go func() {
                for {
                        chat.outgoing <- <-client.incoming
                }
        }()
}

// Remove disconnected client from chat
func (chat *Chat) Remove(i int) {
        chat.clients = append(chat.clients[:i], chat.clients[i+1:]...)
}

// UpdateClientsList sends current connected users list
func (chat *Chat) UpdateClientsList() {
        connectedClients := "/clients>"
        for _, client := range chat.clients {
                connectedClients += client.name + " "
        }
        connectedClients += "\n"
        for _, client := range chat.clients {
                client.outgoing <- connectedClients
        }
}

// Broadcast sends message to all connected clients.
func (chat *Chat) Broadcast(data string) {
        currentTime := time.Now().Format("15:04:05")
        msg := fmt.Sprintf("[%s] %s", currentTime, data)
        for i, client := range chat.clients {
                if client.status == 0 {
                        chat.Remove(i)
                }
        }
        chat.UpdateClientsList()
        for _, client := range chat.clients {
                client.outgoing <- msg
        }
}



func main() {
	fmt.Println("Starting listening for connections...")
        portnum :="" 


	configuration := Configuration{}
	err := gonfig.GetConf(getFileName(), &configuration)
	if err != nil {
		fmt.Println(err)
		//os.Exit(500)
           if len(os.Args) >1 {
             portnum = os.Args[1]
           } else {
            portnum = "5000"
           }

	} else {

	fmt.Println("reading OK! from configure file, port=" + configuration.Port)
        portnum = configuration.Port
           

        };






	listener, err := net.Listen("tcp", ":" + portnum )
	if err != nil {
		fmt.Println(err)
	}

	//chat := server.CreateChat()
	chat := CreateChat()

	for { // listen for connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		chat.Connect(conn)
	}
}

