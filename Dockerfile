# Etapa 1: Construcción (Build)
FROM golang:1.25-alpine AS builder

# Directorio de trabajo
WORKDIR /app

# Copiar archivos de dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente
COPY . .

# Compilar la aplicación (binario estático)
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server/main.go

# Etapa 2: Ejecución (Runtime)
FROM alpine:latest
WORKDIR /root/
# Copiar solo el binario desde la etapa anterior
COPY --from=builder /app/main .

# Exponer el puerto de la API
EXPOSE 8080

# Comando para iniciar
CMD ["./main"]