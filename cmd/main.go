package main

import (
	campaigns "campaigns/pkg"
	"flag"
)

var (
	serverPort = flag.String("server.port", ":8080", "http api server port")
)

func main() {
	flag.Parse()
	// start campaigns server
	campaigns.StartCampaignsServer(*serverPort)
}
