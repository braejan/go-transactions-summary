-- Conectar a la base de datos
\c stori-challenge-db;

-- Eliminar la tabla "users" si existe
DROP TABLE IF EXISTS users;

-- Crear la tabla "users"
CREATE TABLE users (
    id    BIGINT PRIMARY KEY,
    name  TEXT COMMENT 'Nombre del usuario',
    email TEXT UNIQUE COMMENT 'Dirección de correo electrónico del usuario'
);

