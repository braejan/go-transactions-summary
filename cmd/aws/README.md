# README - Arquitectura en AWS

## Descripción

La arquitectura implementada en AWS consta de los siguientes componentes:

1. **API Gateway**: Actúa como la puerta de enlace hacia Internet para el sistema.
2. **VPC (Virtual Private Cloud)**: Una red privada virtual que aloja los servicios del sistema.
3. **Lambda Functions**: Dos funciones Lambda que forman parte del flujo de procesamiento del archivo.
   - **Función Lambda 1**: Recibe la solicitud proxy del API Gateway, valida el archivo y lo carga en un bucket de S3 privado.
   - **Función Lambda 2**: Se activa cuando el bucket de S3 recibe una solicitud de tipo PUT. Descarga el archivo, lo procesa según los casos de uso existentes y se conecta a la base de datos RDS para realizar operaciones.
4. **VPC Endpoints**: Se utilizan para establecer una comunicación segura y privada entre las funciones Lambda y el bucket de S3.
5. **Bucket de S3**: Almacena los archivos de manera privada y segura.
6. **Base de datos RDS**: Una instancia de base de datos PostgreSQL (postgres:10) que se encuentra dentro de la VPC y solo responde a las solicitudes provenientes del grupo de seguridad asignado a la segunda función Lambda.

## Flujo de Procesamiento del Archivo

1. El API Gateway recibe una solicitud y la envía a la Función Lambda 1.
2. La Función Lambda 1 valida el archivo recibido y lo carga en el bucket de S3 utilizando los VPC Endpoints para una comunicación segura.
3. Cuando se realiza una operación de tipo PUT en el bucket de S3, se activa la Función Lambda 2.
4. La Función Lambda 2 descarga el archivo de manera temporal y realiza el procesamiento necesario según los casos de uso implementados.
5. La Función Lambda 2 tiene acceso a las variables de entorno que contienen la información necesaria para conectarse a la base de datos RDS.
6. La Función Lambda 2 se conecta a la base de datos RDS y realiza las operaciones requeridas.
7. La base de datos RDS responde únicamente a las solicitudes provenientes del grupo de seguridad asignado a la segunda función Lambda, garantizando así la seguridad de los datos.

## Arquitectura: Diagrama

A continuación, un gráfico de la arquitectura:

![](https://33333.cdn.cke-cs.com/kSW7V9NHUXugvhoQeFaf/images/bed64221e06d702f7997fb8a2fed084e22a4d6c131fd8756.jpg)

## Consumir la API alojada en AWS:
```bash
curl -X POST -F "file=@/ruta/al/repositorio/samples/file/csv/txns.csv" -F "filename=txns.csv" https://cgwjemuul9.execute-api.us-east-2.amazonaws.com/stg-stori/loadfile
```
---

Este README proporciona una descripción y detalles sobre la arquitectura implementada en AWS para el proyecto.