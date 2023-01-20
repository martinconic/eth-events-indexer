package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSmartContractEvents(c *gin.Context) {
	//get smart contract events
	// if err != nil {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"status": err.Error(),
	// 	})
	// }

	c.JSON(http.StatusOK, gin.H{
		"Events": "Events",
	})
}
