// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: c48ffb9f04
// Version Date: Wed Aug 12 00:08:17 UTC 2020

package main

import (
	"flag"

	// This Service
	"github.com/thebinarytrio/weighted-lottery/backend/lotteryservice-service/svc/server"
)

func main() {
	// Update addresses if they have been overwritten by flags
	flag.Parse()

	server.Run(server.DefaultConfig)
}
