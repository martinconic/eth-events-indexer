package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSmartContractEvents(c *gin.Context) {
	//get smart contract events
	address := "0x3845badAde8e6dFF049820680d1F14bD3903a5d0"

	err := server.NetworkClient.GetContractEvents(address)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"Status": "Started indexing events for Contract: " + address,
	})

}
