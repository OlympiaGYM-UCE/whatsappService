package main

import (
	"fmt"
	"github.com/twilio/twilio-go"
	v2010 "github.com/twilio/twilio-go/rest/api/v2010" // Alias correcto
	"os"
)

// Función para enviar mensaje de WhatsApp
func SendWhatsAppMessage(to string, message string) error {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID") // Corregido
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")   // Corregido
	from := "whatsapp:+14155238886"               // Número de WhatsApp de Twilio

	if accountSid == "" || authToken == "" {
		fmt.Println("Error: Variables de entorno no encontradas")
		return fmt.Errorf("variables de entorno no configuradas")
	}

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &v2010.CreateMessageParams{}
	params.SetTo("whatsapp:" + to)
	params.SetFrom(from)
	params.SetBody(message)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("Error enviando mensaje:", err)
		return err
	}

	fmt.Println("Mensaje enviado con éxito, SID:", *resp.Sid)
	return nil
}
