# Nombre del binario e imagen
APP_NAME = api_go_medium
PORT = 8080

# 🧱 Build local (binario)
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(APP_NAME) -ldflags="-s -w"

# 🧼 Limpiar binarios
clean:
	rm -f $(APP_NAME)
	
# 🚀 Ejecutar localmente con go run
run:
	go run .

# 🐳 Build Docker usando Dockerfile optimizado
docker-build:
	docker build -t $(APP_NAME):latest .

# 🐳 Run Docker (requiere .env para DB externa)
docker-run:
	docker run -p $(PORT):8080 --env-file $(CURDIR)/.env $(APP_NAME):latest

# 🧪 Ejecutar contenedor en background
docker-run-detached:
	docker run -d -p $(PORT):8080 --env-file .env --name $(APP_NAME) $(APP_NAME):latest

# 🛑 Parar y eliminar contenedor
docker-stop:
	docker rm -f $(APP_NAME)