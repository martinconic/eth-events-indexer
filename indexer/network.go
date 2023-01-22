package indexer

import (
	"context"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/martinconic/eth-events-indexer/config"
)

const definition = `[{"constant":true,"inputs":[],"name":"name","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"newAdmin","type":"address"}],"name":"changeExecutionAdmin","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"spender","type":"address"},{"name":"amount","type":"uint256"}],"name":"approve","outputs":[{"name":"success","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"owner","type":"address"},{"name":"amount","type":"uint256"}],"name":"burnFor","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"from","type":"address"},{"name":"to","type":"address"},{"name":"amount","type":"uint256"}],"name":"transferFrom","outputs":[{"name":"success","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"owner","type":"address"},{"name":"spender","type":"address"},{"name":"amount","type":"uint256"}],"name":"approveFor","outputs":[{"name":"success","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"decimals","outputs":[{"name":"","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"owner","type":"address"},{"name":"spender","type":"address"},{"name":"amountNeeded","type":"uint256"}],"name":"addAllowanceIfNeeded","outputs":[{"name":"success","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"amount","type":"uint256"}],"name":"burn","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"who","type":"address"}],"name":"isExecutionOperator","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"who","type":"address"}],"name":"isSuperOperator","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"executionOperator","type":"address"},{"name":"enabled","type":"bool"}],"name":"setExecutionOperator","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"getAdmin","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"from","type":"address"},{"name":"to","type":"address"},{"name":"amount","type":"uint256"},{"name":"gasLimit","type":"uint256"},{"name":"data","type":"bytes"}],"name":"approveAndExecuteWithSpecificGas","outputs":[{"name":"success","type":"bool"},{"name":"returnData","type":"bytes"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"newAdmin","type":"address"}],"name":"changeAdmin","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"symbol","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"to","type":"address"},{"name":"amount","type":"uint256"}],"name":"transfer","outputs":[{"name":"success","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"superOperator","type":"address"},{"name":"enabled","type":"bool"}],"name":"setSuperOperator","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"getExecutionAdmin","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"target","type":"address"},{"name":"amount","type":"uint256"},{"name":"data","type":"bytes"}],"name":"paidCall","outputs":[{"name":"","type":"bytes"}],"payable":true,"stateMutability":"payable","type":"function"},{"constant":false,"inputs":[{"name":"target","type":"address"},{"name":"amount","type":"uint256"},{"name":"data","type":"bytes"}],"name":"approveAndCall","outputs":[{"name":"","type":"bytes"}],"payable":true,"stateMutability":"payable","type":"function"},{"constant":false,"inputs":[{"name":"from","type":"address"},{"name":"to","type":"address"},{"name":"amount","type":"uint256"},{"name":"gasLimit","type":"uint256"},{"name":"tokenGasPrice","type":"uint256"},{"name":"baseGasCharge","type":"uint256"},{"name":"tokenReceiver","type":"address"},{"name":"data","type":"bytes"}],"name":"approveAndExecuteWithSpecificGasAndChargeForIt","outputs":[{"name":"success","type":"bool"},{"name":"returnData","type":"bytes"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"to","type":"address"},{"name":"gasLimit","type":"uint256"},{"name":"data","type":"bytes"}],"name":"executeWithSpecificGas","outputs":[{"name":"success","type":"bool"},{"name":"returnData","type":"bytes"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"owner","type":"address"},{"name":"spender","type":"address"}],"name":"allowance","outputs":[{"name":"remaining","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"from","type":"address"},{"name":"to","type":"address"},{"name":"amount","type":"uint256"},{"name":"gasLimit","type":"uint256"},{"name":"tokenGasPrice","type":"uint256"},{"name":"baseGasCharge","type":"uint256"},{"name":"tokenReceiver","type":"address"}],"name":"transferAndChargeForGas","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"inputs":[{"name":"sandAdmin","type":"address"},{"name":"executionAdmin","type":"address"},{"name":"beneficiary","type":"address"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":true,"name":"to","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Transfer","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"owner","type":"address"},{"indexed":true,"name":"spender","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Approval","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"superOperator","type":"address"},{"indexed":false,"name":"enabled","type":"bool"}],"name":"SuperOperator","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"oldAdmin","type":"address"},{"indexed":false,"name":"newAdmin","type":"address"}],"name":"AdminChanged","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"oldAdmin","type":"address"},{"indexed":false,"name":"newAdmin","type":"address"}],"name":"ExecutionAdminAdminChanged","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"executionOperator","type":"address"},{"indexed":false,"name":"enabled","type":"bool"}],"name":"ExecutionOperator","type":"event"}]`

type NetworkClient struct {
	Config *config.NetworkConfig

	client *ethclient.Client

	Logs chan types.Log
}

type LogTransfer struct {
	from  common.Address
	to    common.Address
	value *big.Int
}

type LogApproval struct {
	owner   common.Address
	spender common.Address
	value   *big.Int
}

func NewNetworkClient(c *config.NetworkConfig) *NetworkClient {
	return &NetworkClient{
		Config: c,
	}
}

type Event struct {
	Key   string
	Value *big.Int
}

func readLogEvents(contractAbi abi.ABI, vLog types.Log, logName string) {
	var event Event

	err := contractAbi.UnpackIntoInterface(&event, logName, vLog.Data)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Key: ", string(event.Key[:]))
	log.Println("Value: ", event.Value)

	for i := range vLog.Topics {
		log.Println(vLog.Topics[i].Hex())
	}

}

func GetERC20Events(contractAbi abi.ABI, vLog types.Log) {
	logTransferSig := []byte("Transfer(address,address,uint256)")
	LogApprovalSig := []byte("Approval(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	logApprovalSigHash := crypto.Keccak256Hash(LogApprovalSig)

	log.Println(vLog.TxIndex)
	log.Println(vLog.Topics)
	log.Printf("Log Block Number: %d\n", vLog.BlockNumber)
	log.Printf("Log Index: %d\n", vLog.Index)

	switch vLog.Topics[0].Hex() {
	case logTransferSigHash.Hex():
		log.Printf("Log Name: Transfer\n")

		var transferEvent LogTransfer
		var event Event

		err := contractAbi.UnpackIntoInterface(&event, "Transfer", vLog.Data)
		if err != nil {
			log.Println(err)
		}
		log.Println(event)

		transferEvent.from = common.HexToAddress(vLog.Topics[1].Hex())
		transferEvent.to = common.HexToAddress(vLog.Topics[2].Hex())
		transferEvent.value = event.Value

		log.Printf("From: %s\n", transferEvent.from.Hex())
		log.Printf("To: %s\n", transferEvent.to.Hex())
		log.Printf("Tokens: %s\n", transferEvent.value.String())

	case logApprovalSigHash.Hex():
		log.Printf("Log Name: Approval\n")

		var approvalEvent LogApproval
		var event Event

		err := contractAbi.UnpackIntoInterface(&event, "Approval", vLog.Data)
		if err != nil {
			log.Println(err)
		}

		approvalEvent.owner = common.HexToAddress(vLog.Topics[1].Hex())
		approvalEvent.spender = common.HexToAddress(vLog.Topics[2].Hex())

		log.Printf("Token Owner: %s\n", approvalEvent.owner.Hex())
		log.Printf("Spender: %s\n", approvalEvent.spender.Hex())
		log.Printf("Tokens: %s\n", approvalEvent.value.String())
	}

	log.Printf("\n\n")

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

	contractAbi, err := abi.JSON(strings.NewReader(definition))
	if err != nil {
		log.Println(err)
		return err
	}

	go func() {
		for {
			select {
			case err = <-sub.Err():
				log.Println(err)
			case vLog := <-nc.Logs:
				log.Println("---------------")
				// readLogEvents(contractAbi, vLog, "Transfer")
				GetERC20Events(contractAbi, vLog)
			}
		}
	}()

	return err
}
