package main

import (
	"github.com/gin-gonic/gin"
	qrcode "github.com/skip2/go-qrcode"
	"net/http"
)

type Request struct {
	Target string `json:"target"`
}
type Response struct {
	Error string `json:"error"`
	Data  []byte `json:"data"`
}

func QrGenerator(c *gin.Context) {
	var req Request
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Error: err.Error(),
		})
		return
	}

	var png []byte
	png, err = qrcode.Encode(req.Target, qrcode.Medium, 512)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Data: png,
	})
}

func main() {
	r := gin.Default()

	r.POST("/qr", QrGenerator)

	r.Run(":8000")
}
