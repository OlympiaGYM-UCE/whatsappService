# WhatsApp Service

A Go-based microservice for sending WhatsApp messages using the Twilio API. Part of the OlympiaGYM-UCE system.

## Features

- Send WhatsApp messages via Twilio
- Environment-based configuration
- Error handling and logging
- Simple API integration

## Prerequisites

- Go 1.16 or higher
- Twilio Account
- WhatsApp Business Account
- Environment variables configured

## Installation

1. Clone the repository:
```bash
git clone https://github.com/your-username/whatsappService.git
cd whatsappService
```

2. Install dependencies:
```bash
go mod download
```

3. Configure environment variables:
```bash
export TWILIO_ACCOUNT_SID=your_account_sid
export TWILIO_AUTH_TOKEN=your_auth_token
```

## Project Structure

```
whatsappService/
├── main.go
├── whatsapp.go      # WhatsApp service implementation
├── go.mod          # Go modules file
├── go.sum          # Go modules checksums
└── Dockerfile      # Docker configuration
```

## Usage

Run the service:
```bash
go run main.go
```

### SendWhatsAppMessage Function

```go
func SendWhatsAppMessage(to string, message string) error
```

Parameters:
- `to`: Recipient's phone number
- `message`: Message content to send

Returns:
- `error`: Any error that occurred during sending

## Environment Variables

- `TWILIO_ACCOUNT_SID`: Your Twilio Account SID
- `TWILIO_AUTH_TOKEN`: Your Twilio Auth Token

## Docker Support

Build the Docker image:
```bash
docker build -t olympiagym/whatsapp-service .
```

Run the container:
```bash
docker run -e TWILIO_ACCOUNT_SID=xxx -e TWILIO_AUTH_TOKEN=xxx olympiagym/whatsapp-service
```

## Dependencies

- github.com/twilio/twilio-go
- os package for environment variables
- fmt package for logging

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
