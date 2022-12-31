package main

import (
	"bytes"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/gin-gonic/gin"
	"image/png"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	router.GET("/:url", Generate)
	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal("router.Run: ", err)
	}
}

func Generate(c *gin.Context) {
	data := c.Param("url")
	if data == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Parâmetro não pode ser vazio!"})
		return
	}

	s, err := url.QueryUnescape(data)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	code, err := qr.Encode(s, qr.L, qr.Auto)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	size := 250
	code, err = barcode.Scale(code, size, size)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, code); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.Header("Content-Type", "image/png")
	c.Header("Content-Length", strconv.Itoa(len(buffer.Bytes())))

	if _, err := c.Writer.Write(buffer.Bytes()); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
}
