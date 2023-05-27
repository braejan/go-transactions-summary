# Transaction Summary - A Golang DDD implementation.
Este proyecto es una aplicación en Golang que procesa un archivo CSV que contiene transacciones y las almacena en una base de datos. Además, envía un correo electrónico al usuario con un resumen de sus transacciones.

## Requerimientos

- Golang
- Postgrest
- REST
- Docker
- AWS

## Instalación

1. Clona este repositorio en tu máquina local.
2. Ubica tu terminal en el directorio `infraestructure/postgres`.
### Construir la base de datos (default):

```console
docker build -t summary-db-image .
docker run -d --name summary-db -p 5432:5432 summary-db-image
```
 #### No te gustan las cosas por defecto?
Puedes entrar a este link oficial de [postgres en docker](https://hub.docker.com/_/postgres) para construir tu base de datos personalizada.
3. Ubica tu terminal en la raíz del proyecto.
```console
cd ../..
```
### Construir la imagen del servicio REST File:
```console
docker build -t summary-image .
docker run -d --name file-rest-service -p 8080:8080 summary-image
```

## Uso
El servicio `/loadfile` permite cargar un archivo CSV en el servidor. Para cargar un archivo, puedes utilizar el siguiente comando curl:

```shell
curl -X POST -F "file=@/ruta/al/archivo.csv" -F "filename=nombre_archivo.csv" http://localhost:8080/loadfile
```
Asegúrate de reemplazar /ruta/al/archivo.csv con la ruta completa de tu archivo CSV y nombre_archivo.csv con el nombre deseado para el archivo.

Es importante tener en cuenta que ambos parámetros son requeridos:

    El parámetro file debe especificar la ubicación del archivo a cargar utilizando el prefijo @.
    El parámetro filename debe contener el nombre que deseas asignar al archivo.

Recuerda que el servicio `/loadfile` está diseñado para aceptar archivos CSV y realizar el procesamiento correspondiente. Asegúrate de proporcionar un archivo válido en formato CSV para obtener los resultados esperados.

## Contribución

Si deseas contribuir a este proyecto, por favor sigue los siguientes pasos:

1. Haz un fork de este repositorio.
2. Crea una rama con una nueva funcionalidad o corrección de errores.
3. Realiza tus cambios y realiza commit.
4. Haz un push de tus cambios a tu repositorio fork.
5. Crea un Pull Request en este repositorio.

## Licencia

Este proyecto está bajo la licencia [MIT License](LICENSE).