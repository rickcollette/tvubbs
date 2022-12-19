package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"tvubbs/bbsconfig"
	"tvubbs/connection"
	"tvubbs/dbstruct"
	"tvubbs/room"

	"gopkg.in/yaml.v3"

	supportscolor "github.com/johnaoss/supports-color"
)

var HasAnsi bool = false

type Server struct {
	Running  bool
	Listener net.Listener
	Rooms    []*room.Room
	LogFile  *os.File
}

// Return a stringified list of rooms
func (s *Server) ListRooms() string {
	str := "Available rooms:\n"
	for i, room := range s.Rooms {
		str += fmt.Sprintf("\t%d: %s\n", i, room.Name)
	}
	return str
}

// Allow the user to select a room, sets the connection's index to the index of
// the room in the server struct's list of rooms
func (s *Server) SelectRoom(c *connection.Connection) error {
	roomList := s.ListRooms()
	if err := c.SendMessage("Select a room to join by its number\n"); err != nil {
		return fmt.Errorf("Failed to send message to %q (%s): %s",
			c.UserName, c.Conn.RemoteAddr(), err.Error())
	}
	room, err := c.SendWithResponse(roomList)
	if err != nil {
		return fmt.Errorf("Failed to send message to %q (%s): %s\n",
			c.UserName, c.Conn.RemoteAddr(), err.Error())
	}

	if room == "" {
		c.SendError("Must enter room name to join")
		return fmt.Errorf("User %q (%s) failed to choose room\n", c.UserName, c.Conn.RemoteAddr())
	}

	roomIndex, err := strconv.Atoi(room)
	if err != nil {
		return fmt.Errorf("Error choosing room for user %s: %s\n", c.String(), err.Error())
	} else if roomIndex > len(s.Rooms) || roomIndex < 0 {
		return fmt.Errorf("User %s selected invalid room\n", c.String())
	}

	s.Rooms[roomIndex].AddUser(c)
	c.Room = roomIndex

	return nil
}

// Handle various user chat commands. Only available when in a room
func (s *Server) HandleChatCommands(message string, c *connection.Connection) bool {
	switch message {
	case "/help":
		if err := c.SendMessage("Help Message"); err != nil {
			log.Println(err)
			return true
		}
	case "/name":
		newName, err := c.SendWithResponse("New name: ")
		if err != nil {
			log.Println(err)
			return true
		}
		log.Printf("User %s changed name to %s\n", c.String(), newName)
		s.Rooms[c.Room].WriteMessage(fmt.Sprintf("User %s changed name to %s\n", c.UserName, newName))
		c.UserName = newName

		return true
	case "/leave":
		room := s.Rooms[c.Room]
		room.RemoveUser(c)
		s.SelectRoom(c)
		return true
	case "/quit":
		c.Close()
		return true
	}
	return false
}

// Handle user messages to a room as well as commands. Exits when the user disconnects
func (s *Server) HandleMessages(c *connection.Connection) {
	for c.Open == true {
		text, err := c.SendWithResponse(">> ")
		if err != nil {
			log.Printf("Failed to read message from %s: %s", c.String(), err.Error())
			return
		}

		if s.HandleChatCommands(text, c) == true {
			continue
		}

		message := fmt.Sprintf("<%s> (%s): %s\n", time.Now().Format(time.Kitchen), c.UserName, text)
		room := s.Rooms[c.Room]
		room.WriteMessage(message)

		logStr := fmt.Sprintf("%s: %s", room.Name, message)
		_, err = s.LogFile.WriteString(logStr)
		if err != nil {
			log.Printf("Failed to log message from user %s: %s", c.String(), err.Error())
		}

		log.Printf("User %s sent message %q to room %q\n", c.String(), text, room.Name)
	}
}

func (s *Server) MainMenu(c *connection.Connection) {
	if err := c.SendMessage("Welcome to the TVU BBS!\n"); err != nil {
		log.Println(err)
		return
	}

	if err := s.SelectRoom(c); err != nil {
		log.Println(err)
		return
	}

	s.doChat(c)
}

func (s *Server) doChat(c *connection.Connection) {
	if err := s.SelectRoom(c); err != nil {
		log.Println(err)
		return
	}

	go s.HandleMessages(c)
}
// Initialize the connection object and start a go routine to handle messaging with the client
func (s *Server) HandleConnection(c *connection.Connection) {
	ansi := supportscolor.GetSupportLevel()
	if ansi.Has1m {
		HasAnsi = true
	} else if ansi.Has256 {
		HasAnsi = true
	} else if ansi.HasBasic {
		HasAnsi = true
	} else {
		HasAnsi = false
	}

	prelog, err := os.ReadFile("ascii/prelog.txt")
	if err != nil {
		log.Println(err)
		return
	} else {
		c.SendMessage(string(prelog))
	}
	c.SendMessage("\n")
	c.SendMessage("If you are a new user, please enter NEW. Otherwise, enter your username.\n")
	username, err := c.SendWithResponse("Username: ")
	if err != nil || username == "" {
		c.Close()
		log.Println("User failed to enter username")
		return
	}

	c.UserName = username
	log.Printf("User %s connected\n", c.String())

	if err := s.SelectRoom(c); err != nil {
		log.Println(err)
		c.Close()
		return
	}
	s.MainMenu(c)
}

// Start the server's room go-routines, start the tcp listener and handle incoming connections
func (s *Server) Serve() {

	for _, room := range s.Rooms {
		log.Printf("Starting room %q...\n", room.Name)
		go room.Run()
	}

	for s.Running {
		conn, err := s.Listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		c := connection.NewConnection(conn)
		go s.HandleConnection(c)
	}
}

// Initialize the rooms in a server
func (s *Server) InitializeRooms() {
	for _, roomName := range bbsconfig.BbsConfig.Rooms {
		log.Printf("Initializing room %q\n", roomName)
		s.Rooms = append(s.Rooms, &room.Room{
			Name:        roomName,
			Connections: make(map[string]*connection.Connection, 0),
			WriteChan:   make(chan string),
		})
	}
}

func LoadConfig() (*dbstruct.Sysconfig, error) {
	fmt.Printf("Checking Databases...\n")
	BaseConfig := &dbstruct.Sysconfig{}
	file, err := os.Open("data/bbsconfig.yml")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&BaseConfig); err != nil {
		return nil, err
	}
	return (BaseConfig), nil
}

// Initiaize a new server with setttings read from the configuration file
func NewServer() (*Server, error) {

	BindAddr := bbsconfig.BbsConfig.BindAddr

	log.Println("Starting listener on", BindAddr)
	listener, err := net.Listen("tcp4", BindAddr)
	if err != nil {
		return nil, err
	}

	s := &Server{
		Running:  true,
		Listener: listener,
		Rooms:    make([]*room.Room, 0),
	}

	s.InitializeRooms()

	return s, nil
}
