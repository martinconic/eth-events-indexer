package main

import (
	"fmt"
	"os"

	"github.com/martinconic/eth-events-indexer/config"
	"github.com/martinconic/eth-events-indexer/indexer"
)

func main() {

	v, err := config.NewViper()
	if err != nil {
		fmt.Println("Error on init")
		os.Exit(1)
	}

	network := indexer.NewNetworkClient(config.GetNetworkConfigs(v))
	network.GetContractEvents(v.GetString("network.address"))

}
