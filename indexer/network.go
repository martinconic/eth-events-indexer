package indexer

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/martinconic/eth-events-indexer/config"
)

type NetworkClient struct {
	Config *config.NetworkConfig

	client *ethclient.Client
}

func NewNetworkClient(c *config.NetworkConfig) *NetworkClient {
	return &NetworkClient{
		Config: c,
	}
}

func (nc *NetworkClient) GetContractEvents(address string) {
	var err error

	rawUrl := nc.Config.Wss + nc.Config.Key
	nc.client, err = ethclient.Dial(rawUrl)

	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(address)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := nc.client.SubscribeFilterLogs(context.Background(), query, logs)
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
