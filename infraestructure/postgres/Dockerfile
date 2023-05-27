# This is the Dockerfile for the postgres image to run the database locally
# and create the database and tables needed for the project

FROM postgres

# Copiar el archivo de creación de la base de datos
COPY database.sql /docker-entrypoint-initdb.d/

# Variables de entorno para la configuración
ENV POSTGRES_HOST=localhost
ENV POSTGRES_PORT=5432
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=postgres
ENV POSTGRES_DB=stori-challenge-db

# Puerto expuesto por el contenedor
EXPOSE 5432

# Comando de inicio de la base de datos
CMD ["postgres"]