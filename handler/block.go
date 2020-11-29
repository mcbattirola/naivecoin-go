package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mcbattirola/naivecoin-go/blockchain"
)

type blockDTO struct {
	Data  string `json:"data" binding:"required"`
	Nonce string `json:"nonce" binding:"required"`
}

// GetBlocks returns all the blocks in the blockchain
func GetBlocks(c *gin.Context) {
	c.JSON(200, gin.H{
		"blockchain": blockchain.GetBlockchain(),
	})
}

// MineBlock creates a block from input data
func MineBlock(c *gin.Context) {
	var input blockDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"block": blockchain.GenerateNextBlock(input.Data, input.Nonce),
	})
}
