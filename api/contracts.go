package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSmartContracts(c *gin.Context) {
	sc, err := server.db.GetSmartContracts()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": err.Error(),
		})
		return
	}
	if len(sc) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": "empty, not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Smart contracts in db: " + fmt.Sprint(sc),
	})
}
