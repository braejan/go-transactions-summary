-- Conectar a la base de datos
\c stori-challenge-db;

-- Eliminar la tabla "users" si existe
DROP TABLE IF EXISTS users;

-- Crear la tabla "users"
CREATE TABLE users (
    id    BIGINT PRIMARY KEY,
    name  TEXT, 
    email TEXT UNIQUE
);
COMMENT ON TABLE users IS 'Tabla de usuario';
COMMENT ON COLUMN users.id IS 'Nombre del usuario';
COMMENT ON COLUMN users.id IS 'Dirección de correo electrónico del usuario';


CREATE TABLE accounts (
    id      UUID PRIMARY KEY,
    balance BIGINT,
    userid  BIGINT UNIQUE,
    active  BOOLEAN
);
COMMENT ON TABLE accounts IS 'Tabla de cuentas de usuario';
COMMENT ON COLUMN accounts.id IS 'Identificador único de la cuenta';
COMMENT ON COLUMN accounts.balance IS 'Saldo de la cuenta';
COMMENT ON COLUMN accounts.userid IS 'ID de usuario asociado a la cuenta';
COMMENT ON COLUMN accounts.active IS 'Indica si la cuenta está activa o no';

ALTER TABLE accounts
ADD CONSTRAINT fk_account_user
FOREIGN KEY (userid)
REFERENCES users (id)
ON DELETE CASCADE;
