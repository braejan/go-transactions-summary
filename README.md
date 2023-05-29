# Transaction Summary - A Golang DDD implementation.

## Historia
En un mundo de una especie consciente de sí misma ha decidido crear un sistema financiero global en el cuál para ser parte del sistema financiero sólo debe proporcionar su número de identifiación única (es un Integer 😉). Las transacciones son simples. Si el ciudadano necesita hacer un pago, se registra en un archivo con un signo positivo (➕) de lo contrario, será con un sigo negativo (➖). Cada registro tiene una fecha en formado MM/AA (el número del mes y los dos últimos digitos del año).

## Objetivo
Reunir la información de las compras y ventas realizadas y almacenadas en un archivo de formato texto plano (CSV) y enviar un correo electrónico al usuario con un resumen de sus transacciones.

## Solución Planteada
Se ha decidido utilizar una arquitectura orientada a dominio (DDD) para resolver este problema. Separando los dominios de cuentas, usuarios y todo el manejo de archivo por aparte. La aplicación está desacoplada, de tal manera que es muy fácil añadirle un REST (cmd/api/file/api.go) o una función lamdba en AWS (cmd/aws/process/process.go) para que pueda ser consumida por otros sistemas. 

Al ser un motivo de ejemplo, el archivo CSV debe cumplir la siguiente estructura:
```csv
Id,Date,Transaction
0,7/5,+60.5
1,7/28,-10.3
2,8/2,-20.46
3,8/13,+10
4,8/20,-15.75
```
### Consideraciones y restricciones:
- El archivo debe estar en formato CSV.
- El archivo debe tener una cabecera.
- El archivo debe tener 3 columnas: Id, Date, Transaction.
- La columna Id debe ser un número entero.
- La columna Date debe ser una fecha en formato MM/AA.
- La columna Transaction debe ser un número decimal.
- La columna Transaction debe tener un signo positivo (➕) o negativo (➖).
- El archivo debe tener al menos un registro.
- Los Id no necesariamente deben ser únicos. El sistema considera que el userID exista garantizando su creación. También aplica para la cuenta interna que maneja el sistema.

## Requerimientos

- Golang 1.19 o superior (https://golang.org/dl/)
- Docker (https://docs.docker.com/get-docker/)
- Docker Compose (https://docs.docker.com/compose/install/)

## Instalación

1. Clona este repositorio en tu máquina local.
2. Ubica tu terminal en el directorio del repositorio `go-transactions-summary`.
### Construir el proyecto localmente usando docker compose:

```console
docker-compose up
```
#### Puede tomar un tiempo mientras descarga las dependencias y construye el proyecto.
    Wait for it.
#### No te gustan las cosas por defecto?
Puedes entrar a este link oficial de [postgres en docker](https://hub.docker.com/_/postgres) para construir tu base de datos personalizada.

## Uso
El servicio `/loadfile`  es un servicio REST que está publicado en el puerto `8080`. Permite cargar un archivo CSV en el servidor. Para cargar un archivo, puedes utilizar el siguiente comando curl:

```shell
curl -X POST -F "file=@/ruta/al/repositorio/samples/file/csv/txns.csv" -F "filename=txns.csv" http://localhost:8080/loadfile
```
Asegúrate de reemplazar ruta/al/repositorio/samples/file/csv/txns.csv con la ruta completa de tu archivo CSV y txns.csv con el nombre deseado para el archivo.

Es importante tener en cuenta que ambos parámetros son requeridos:

    El parámetro file debe especificar la ubicación del archivo a cargar utilizando el prefijo @.
    El parámetro filename debe contener el nombre que deseas asignar al archivo.

Recuerda que el servicio `/loadfile` está diseñado para aceptar archivos CSV y realizar el procesamiento correspondiente. Asegúrate de proporcionar un archivo válido en formato CSV para obtener los resultados esperados.

## Pruebas

Para ejecutar las pruebas unitarias, debes ejecutar el siguiente comando:
```bash
go test -v ./... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
```
Esta ejecución generará en consola el resultado de la ejecución del set de pruebas de todos los archivos *_test.go. Además, generará un archivo coverage.html que puedes abrir en tu navegador para ver el porcentaje de cobertura de las pruebas.

## Deuda técnica.
El proyecto debe enviar un email con el resumen del resultado, sin embargo queda pendiente su implementación usando un servicio de email cómo AWS SES.
## Contribución

Si deseas contribuir a este proyecto, por favor sigue los siguientes pasos:

1. Haz un fork de este repositorio.
2. Crea una rama con una nueva funcionalidad o corrección de errores.
3. Realiza tus cambios y realiza commit.
4. Haz un push de tus cambios a tu repositorio fork.
5. Crea un Pull Request en este repositorio.

## Licencia

Este proyecto está bajo la licencia [MIT License](LICENSE).