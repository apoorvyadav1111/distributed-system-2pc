package main

import (
	"log"

	delivery "github.com/apoorvyadav1111/distributed-systems-2pc/delivery/svc"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new delivery service
	delivery.Clean()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.POST("/delivery/agent/reserve", func(c *gin.Context) {
		agent, err := delivery.ReserveAgent(c)
		if err != nil {
			c.JSON(429, err)
			return
		}
		c.JSON(200, delivery.ReserveAgentResponse{AgentID: agent.ID})
	})

	r.POST("/delivery/agent/assign", func(c *gin.Context) {
		var req delivery.AssignAgentRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithStatus(429)
		}
		agent, err := delivery.AssignAgent(req.OrderId)
		if err != nil {
			c.JSON(429, err)
			return
		}
		c.JSON(200, delivery.AssignAgentResponse{AgentID: agent.ID, OrderID: req.OrderId})
	})

	log.Println("Server is running on port 8080")
	r.Run(":8082")
}
