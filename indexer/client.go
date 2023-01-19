package indexer

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getContractEvents() {
	client, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/ee163cf375034b0aacee52440071e358")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x3845badAde8e6dFF049820680d1F14bD3903a5d0")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		}
	}
}
