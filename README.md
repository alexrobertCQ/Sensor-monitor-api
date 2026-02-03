# Industrial Sensor Monitor API (Go + PostgreSQL + Docker)

Sistema Backend desarrollado por un **Mecatrónico Industrial** enfocado en el monitoreo de variables críticas (pH, Presión, Temperatura) en entornos de calderas y procesos industriales.

## 🛠️ Tecnologías
- **Lenguaje:** Go (Golang) con arquitectura limpia.
- **Framework:** Gin Gonic (API REST).
- **Base de Datos:** PostgreSQL para persistencia de datos.
- **Infraestructura:** Docker & Docker Compose.

## 🚀 Características
- **Lógica de Alertas:** Detección automática de valores críticos (ej. pH > 12.0 para evitar incrustaciones).
- **Persistencia:** Almacenamiento seguro en base de datos relacional.
- **Portabilidad:** Despliegue inmediato mediante contenedores.

## 📦 Ejecución
```bash
docker-compose up --build