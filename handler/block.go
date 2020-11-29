package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mcbattirola/naivecoin-go/blockchain"
)

type BlockDTO struct {
	Data string `json:"data" binding:"required"`
}

func GetBlocks(c *gin.Context) {
	c.JSON(200, gin.H{
		"blockchain": blockchain.GetBlockchain(),
	})
}

func MineBlock(c *gin.Context) {
	var input BlockDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"block": blockchain.GenerateNextBlock(input.Data),
	})
}
