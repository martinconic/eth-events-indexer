package main

import (
	"fmt"
	"os"

	"github.com/martinconic/eth-events-indexer/api"
	"github.com/martinconic/eth-events-indexer/config"
)

func main() {

	v, err := config.NewViper()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	api.StartServer(v)

}
