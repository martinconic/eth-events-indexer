package indexer

import (
	"context"
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

	Logs chan types.Log
}

func NewNetworkClient(c *config.NetworkConfig) *NetworkClient {
	return &NetworkClient{
		Config: c,
	}
}

func (nc *NetworkClient) GetContractEvents(address string) error {
	var err error

	rawUrl := nc.Config.Wss + nc.Config.Key
	nc.client, err = ethclient.Dial(rawUrl)

	if err != nil {
		log.Fatal(err)
		return err
	}

	contractAddress := common.HexToAddress(address)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	nc.Logs = make(chan types.Log)
	sub, err := nc.client.SubscribeFilterLogs(context.Background(), query, nc.Logs)
	if err != nil {
		log.Fatal(err)
		return err
	}

	go func() {
		for {
			select {
			case err = <-sub.Err():
				log.Println(err)
			case vLog := <-nc.Logs:
				log.Println(vLog)
			}
		}
	}()

	return err
}
