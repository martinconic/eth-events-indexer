package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetSmartContractEvents(c *gin.Context) {
	address := c.Param("address")

	addrId, err := server.db.Get(address)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = server.NetworkClient.GetContractEvents(address)
	if err != nil {
		log.Println("Failed to get events for: ", address)
	}
	contract := server.NetworkClient.Contracts[address]
	message, err := server.db.UpdateIndexing(address, true)
	if err != nil {
		log.Println(err)
	}
	log.Println(message)

	go func() {
		for {
			select {
			case err = <-contract.Sub.Err():
				log.Println(err)
			case vLog := <-contract.Logs:
				log.Println("---------------")
				// readLogEvents(contractAbi, vLog, "Transfer")
				tx := contract.GetERC20Events(&vLog)
				tx.ID = addrId
				server.db.InsertEvent(tx)
			}
		}
	}()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Started indexing events for Contract: " + address,
	})

}

func GetIndexedEvents(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	tx, err := server.db.GetEvents(id)
	log.Println(tx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": err.Error(),
		})
		return
	}
	if tx == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": "empty, not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Tx[0] Address: " + tx[0].TxAddr,
	})
}

func AddSmartContract(c *gin.Context) {
	address := c.Param("address")
	_, err := server.db.Insert(address)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Success inserting contract: " + address,
	})
}
