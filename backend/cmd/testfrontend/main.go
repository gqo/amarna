package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
)

type Packet struct {
	Action   string `json:"action"`
	CompName string `json:"compName"`
}

var port = flag.Int("listen", 8080, "the port number")          //default port is 8080
var backendPort = flag.Int("backend", 8081, "backend port num") //default backend port is 8081

func main() {
	c, e := net.Dial("tcp", ":"+strconv.Itoa(*backendPort))
	if e != nil {
		fmt.Println("There was a connection error")
		log.Panicln(e)
	}
	defer c.Close()

	encode := json.NewEncoder(c)

	inputJSON := Packet{
		Action:   "GET",
		CompName: "TEST",
	}

	jsonEncodeErr := encode.Encode(inputJSON)

	if jsonEncodeErr != nil {
		fmt.Printf("Ertror is: %v", jsonEncodeErr)
	}

}
