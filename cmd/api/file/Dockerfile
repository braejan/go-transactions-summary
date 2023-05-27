# Etapa de compilación para el servicio de cuentas
FROM golang:1.19-alpine AS build-account

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o account ./cmd/api/account

# Etapa de compilación para el servicio de archivos
FROM golang:1.19-alpine AS build-file

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o file ./cmd/api/file

# Imagen final para el servicio de cuentas
FROM alpine:latest

WORKDIR /app

COPY --from=build-account /app/account .

EXPOSE 8080

CMD ["./account"]

# Imagen final para el servicio de archivos
FROM alpine:latest

WORKDIR /app

COPY --from=build-file /app/file .

EXPOSE 8080

CMD ["./file"]
