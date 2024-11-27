package middleware

import (
	"fmt"
	"log"
	"news-go/config"
	"news-go/models"

	"github.com/gin-gonic/gin"
	"github.com/ua-parser/uap-go/uaparser"
)

// Middleware for collecting user metadata
func GetAllClientData(c *gin.Context) {
	userAgent := c.Request.Header.Get("User-Agent")
	parser := uaparser.NewFromSaved()
	client := parser.Parse(userAgent)


	clientOS := client.Os.Family
	clientBrowser := client.UserAgent.Family
	clientDevice := client.Device.Family
	clientDeviceBrand := client.Device.Brand
	clientDeviceModel := client.Device.Model

	ip := fmt.Sprintf("Client IP: %s", c.ClientIP())

	os := fmt.Sprintf("Client OS family: %s", clientOS)
	browser := fmt.Sprintf("Client browser family: %s", clientBrowser)
	device := fmt.Sprintf("Client device family: %s, brand: %s, model: %s", clientDevice, clientDeviceBrand, clientDeviceModel)

	entry := models.Client_metadata{
		Ip: ip,
		Os: os,
		Browser: browser,
		Device: device,
	}

	if err := config.DB.Create(&entry).Error; err != nil {
		log.Printf("Insert failed: %v", err)
		c.JSON(500, gin.H{"error": "Failed to insert client metadata"})
		return
	}
}