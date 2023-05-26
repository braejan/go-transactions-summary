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

CREATE TABLE transactions (
    id         UUID PRIMARY KEY,
    accountid UUID NOT NULL,
    amount     FLOAT NOT NULL,
    operation  VARCHAR(255) NOT NULL,
    date       TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL,
    origin     VARCHAR(255) NOT NULL
);

ALTER TABLE transactions
ADD CONSTRAINT fk_transaction_account
FOREIGN KEY (accountid)
REFERENCES accounts (id)
ON DELETE CASCADE;

CREATE INDEX idx_transactions_account_operation
ON transactions(accountid, operation);

CREATE INDEX idx_transactions_origin
ON transactions(origin);

COMMENT ON TABLE transactions IS 'Table to store transactions data';

COMMENT ON COLUMN transactions.id IS 'Transaction ID';
COMMENT ON COLUMN transactions.accountid IS 'Account ID of the transaction';
COMMENT ON COLUMN transactions.amount IS 'Amount of the transaction';
COMMENT ON COLUMN transactions.operation IS 'Operation of the transaction';
COMMENT ON COLUMN transactions.date IS 'Date of the transaction';
COMMENT ON COLUMN transactions.created_at IS 'Date and time when the transaction was created';
COMMENT ON COLUMN transactions.origin IS 'Origin of the transaction';
