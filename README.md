# Go-Rest-Playground

Microsservices development in GO(Golang)

## Services
  - auth-service: Autentication and validate users
  - product-service: Api return products

## Skils
  - Hexagonal architecture
  - Framework GIN (Gin-gonic)
  - API Restfull
  - gRPC
  - PostgreSQL
  - Gorm(Database ORM)
  - Migrations
  - RabbitMQ
  - JWT
  - Bcrypt
  - Gin Middleware
  - OpenAPI(Swagger)
  - Docker
  - Enviroment variables
  - DTOs
  - Assemblers or Model Mappers
  - Wirings(Modulate app)
 
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
