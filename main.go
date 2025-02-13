package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

// Estructura del request
type WhatsAppRequest struct {
	To      string `json:"to" binding:"required"`
	Message string `json:"message" binding:"required"`
}

func main() {
	// Cargar variables de entorno desde .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando el archivo .env")
	}

	// Verificar que las variables de entorno se cargaron correctamente
	log.Println("TWILIO_ACCOUNT_SID:", os.Getenv("TWILIO_ACCOUNT_SID"))

	r := gin.Default()

	// Endpoint para enviar mensaje de WhatsApp
	r.POST("/send-whatsapp", func(c *gin.Context) {
		var req WhatsAppRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := SendWhatsAppMessage(req.To, req.Message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error enviando el mensaje"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Mensaje enviado con éxito"})
	})

	// Iniciar servidor en el puerto 8080
	r.Run(":9090")
}
