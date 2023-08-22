package main

import (
	"fmt"
	"net/http"

	contractModule "august-wasm/contract"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})

	r.GET("/:contract_id/state", func(c *gin.Context) {
		contractId := c.Param("contract_id")
		contractExec, err := contractModule.NewContractExecution(contractId)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Not found",
			})
		}
		stateStr := contractExec.ReadStateFile()
		c.JSON(http.StatusOK, stateStr)
	})

	r.POST("/:contract_id/events", func(c *gin.Context) {
		contractId := c.Param("contract_id")
		contractExec, err := contractModule.NewContractExecution(contractId)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Not found",
			})
		}
		var body []contractModule.Action

		// Marshall body into Action array
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Bad request",
			})
			return
		}

		contractExec.ProcessActions(body)
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	// Log requests
	r.Use(gin.Logger())

	r.Run(":8304")
}
