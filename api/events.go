package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSmartContractEvents(c *gin.Context) {
	a := c.Param("address")

	address, err := server.db.Get(a)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = server.NetworkClient.GetContractEvents(address)
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

func AddSmartContract(c *gin.Context) {
	address := c.Param("address")
	err := server.db.Insert(address)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Success inserting contract: " + address,
	})
}
