package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mcbattirola/naivecoin-go/p2p"
)

type peerDTO struct {
	URL string `json:"url"`
}

// GetPeers return all current available peers
func GetPeers(c *gin.Context) {
	sockets := p2p.GetSockets()

	c.JSON(200, gin.H{
		"peers": sockets,
	})
}

// AddPeer initiates a connection with the peers
func AddPeer(c *gin.Context) {
	var input peerDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := p2p.ConnectToPeers(input.URL)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
}
