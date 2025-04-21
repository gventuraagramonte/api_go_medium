<p align="center">
  <img src="assets/logo.png" width="150" alt="Logo Giorgio">
</p>

<h3 align="center">Giorgio API Backend</h3>
<p align="center">
  Backend Developer · Gopher 🐹 · Clean API Architect
</p>

---

## ✈️ API de Pasajeros en Go

Una API REST desarrollada en Go que permite gestionar pasajeros mediante operaciones CRUD. Utiliza PostgreSQL (Neon) como base de datos y GORM como ORM.

---

## 🚀 Características

- Listado de pasajeros activos
- Registro de nuevos pasajeros
- Consulta por número de asiento
- Soft delete (desactivación de pasajeros)
- Conexión segura usando variables de entorno
- Validación segura de tokens JWT usando JWKS (Auth0)
- Conexión segura usando variables de entorno
- Protecciones estáticas con `gosec` (sin `InsecureSkipVerify`)

---

## 🛠 Tecnologías

- [Go](https://golang.org/) 1.20+
- [GORM](https://gorm.io/)
- [PostgreSQL (Neon)](https://neon.tech/)
- [godotenv](https://github.com/joho/godotenv)
- [Auth0](https://auth0.com/)
- [MicahParks/keyfunc](https://github.com/MicahParks/keyfunc) para JWKS
- `crypto/x509` y `tls.Config` para validación de certificados

---

## 📁 Estructura del proyecto

```
api_go_medium/
├── main.go              # Punto de entrada
├── go.mod / go.sum      # Configuración de dependencias
├── .env                 # Variables sensibles (NO versionar)

├── models/              # Modelos y conexión a PostgreSQL
│   ├── pasajero.go
│   └── database.go

├── controllers/         # Lógica de negocio (handlers)
│   └── pasajero_controller.go

└── routes/              # Registro centralizado de rutas
    └── routes.go
```

---

## 🧪 Endpoints

| Método | Ruta                              | Descripción                |
| ------ | --------------------------------- | -------------------------- |
| GET    | `/pasajeros`                      | Listar pasajeros activos   |
| GET    | `/pasajeros/{asiento}`            | Buscar pasajero por número |
| POST   | `/pasajeros`                      | Agregar un nuevo pasajero  |
| PUT    | `/pasajeros/{asiento}/desactivar` | Soft delete (desactivar)   |

---

## ⚙️ Configuración

1. Clona el repositorio:

```bash
git clone https://github.com/tuusuario/api_go_medium.git
cd api_go_medium
```

2. Crea el archivo `.env` con tu conexión a PostgreSQL:

```env
DB_HOST=yourhost.neon.tech
DB_USER=youruser
DB_PASS=yourpassword
DB_NAME=yourdbname
DB_PORT=5432
```

3. Instala dependencias:

```bash
go mod tidy
```

4. Corre el proyecto:

```bash
go run .
```

---

---

## 🔐 Seguridad con GOSEC

**Gosec** es una herramienta para escanear código Go en busca de vulnerabilidades de seguridad.

### 📥 Instalación

```bash
go install github.com/securego/gosec/v2/cmd/gosec@latest
```

> ✅ Asegúrate de tener `$GOPATH/bin` en tu `PATH`. Por defecto en Windows: `C:\Users\TU_USUARIO\go\bin`

### 🚦 Ejecutar escaneo

```bash
# Escaneo rápido
$ make gosec

# Generar reporte HTML
$ make gosec-report

# Revisa el archivo generado:
$ open gosec-report.html
```

---

## 📦 Archivos versionables

- `go.mod`, `go.sum`
- `main.go`
- `/models`, `/controllers`, `/routes`
- `README.md`

## 🚫 Ignorar (añadir a .gitignore)

```
.env
*.log
*.tmp
*.exe
.vscode/
.idea/
```

---

## 👨‍💻 Autor

**Giorgio**  
Backend Developer | Go Enthusiast | API Architect  
Construido con 💡 propósito, ⚙️ precisión y ☕ pasión.

---
