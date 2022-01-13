# Go-Rest-Playground

Microsservices development in GO(Golang)

## Services
  - auth-service: Autentication and validate users
  - product-service: Api return products

## Skils

### Architecture/DDD
  - Hexagonal architecture
### Framework
- GIN (Gin-gonic)

### HTTP
- API Restfull
- gRPC

### DATABASE
- Gorm(Database ORM)
- Migrations

### Messaging
  - RabbitMQ

### Autentication
  - JWT
  - Bcrypt

### DOCS
  - OpenAPI(Swagger)

### OTHERS
  - Docker
  - Enviroment variables
  - DTOs
  - Assemblers or Model Mappers
 
## Bootstrap

    docker-compose up -d

## Start Auth Service
  - gRPC:
      
        go run auth-service/grpc/main.go

  - API:

        go run auth-service/api/main.go

## Start Products Service
  - API:

        go run products-service/api/main.go
