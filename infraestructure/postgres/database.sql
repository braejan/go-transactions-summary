-- Eliminar la base de datos si existe
DROP DATABASE IF EXISTS stori-challenge-db;

-- Crear la base de datos
CREATE DATABASE stori-challenge-db;

-- Conectar a la base de datos
\c stori-challenge-db;

-- Crear la tabla "users"
CREATE TABLE users (
    id    BIGINT PRIMARY KEY,
    name  TEXT,
    email TEXT
);
