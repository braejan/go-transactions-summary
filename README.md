# go-transactions-summary
Este proyecto es una aplicación en Golang que procesa un archivo CSV que contiene transacciones y las almacena en una base de datos. Además, envía un correo electrónico al usuario con un resumen de sus transacciones.

## Requerimientos

- Golang
- Postgrest
- REST
- Docker
- AWS

## Instalación

1. Clona este repositorio en tu máquina local.
2. Configura la conexión a la base de datos en el archivo `config.go`.
3. Configura las credenciales de envío de correo electrónico en el archivo `config.go`.
4. Ejecuta el comando `go run main.go` para iniciar la aplicación.

## Uso

1. Coloca el archivo CSV con las transacciones en la carpeta `data`.
2. Ejecuta la aplicación y espera a que procese el archivo.
3. Una vez completado, recibirás un correo electrónico con el resumen de tus transacciones.

## Contribución

Si deseas contribuir a este proyecto, por favor sigue los siguientes pasos:

1. Haz un fork de este repositorio.
2. Crea una rama con una nueva funcionalidad o corrección de errores.
3. Realiza tus cambios y realiza commit.
4. Haz un push de tus cambios a tu repositorio fork.
5. Crea un Pull Request en este repositorio.

## Licencia

Este proyecto está bajo la licencia [MIT License](LICENSE).