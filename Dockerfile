# 1️⃣ Usa la imagen base de Go más reciente
FROM golang:1.24-alpine

# 2️⃣ Configurar el directorio de trabajo dentro del contenedor
WORKDIR /app

# 3️⃣ Copiar archivos de dependencias primero
COPY go.mod go.sum ./
RUN go mod download

# 4️⃣ Copiar todo el código fuente
COPY . .

# 5️⃣ Compilar todos los archivos (no solo main.go)
RUN go build -o whatsapp-service .

# 6️⃣ Exponer el puerto en el que corre la API
EXPOSE 9090

# 7️⃣ Ejecutar la aplicación
CMD ["/app/whatsapp-service"]
