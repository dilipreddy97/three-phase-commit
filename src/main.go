package main

import (
	"os"
	"strconv"
	//"fmt"
)

func main() {

	args := os.Args[1:4]

	serverId := args[0]
	n, _ := strconv.Atoi(args[1])
	masterFacingPort := args[2]
	id_num, _ := strconv.Atoi(serverId)

	peerFacingPort := strconv.Itoa(20000 + id_num)

	var peers []string
	var server Server

	for i := 0; i < n; i++ {
		peerStr := strconv.Itoa(20000 + i)
		peers = append(peers, peerStr)
	}

	if masterFacingPort == "10002" { // this is only on first startup
		server = Server{pid: serverId, peers: peers, masterFacingPort: masterFacingPort,
			peerFacingPort: peerFacingPort, is_coord: true, playlist: make(map[string]string), crashStage: ""}
	} else {
		server = Server{pid: serverId, peers: peers, masterFacingPort: masterFacingPort,
			peerFacingPort: peerFacingPort, is_coord: false, playlist: make(map[string]string), crashStage: ""}
	}

	server.run()

	os.Exit(0)

}
