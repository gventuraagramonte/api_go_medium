# 🔧 Etapa 1: build
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copiar go.mod y go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código
COPY . .

# Compilar el binario estático
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api_go_medium -ldflags="-s -w"

# 🐳 Etapa 2: imagen final mínima
FROM scratch

# Copiar el binario desde el builder
COPY --from=builder /app/api_go_medium /api_go_medium

# Puerto de la aplicación
EXPOSE 8080

# Ejecutar el binario
ENTRYPOINT [ "/api_go_medium" ]