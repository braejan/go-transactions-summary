# Etapa de compilación para el servicio de archivos
FROM golang:1.19-alpine AS build-file-rest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o rest-file ./cmd/api/file

# Imagen final para el servicio de archivos
FROM alpine:latest

WORKDIR /app

COPY --from=build-file-rest /app/rest-file .

EXPOSE 8080

CMD ["./rest-file"]
