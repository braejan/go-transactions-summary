# README - Base de Datos

## Descripción

Este esquema de base de datos corresponde a una base de datos PostgreSQL llamada "stori-challenge-db". La base de datos contiene las siguientes tablas:

1. **users**: Tabla de usuarios.
   - Columnas:
     - id (BIGINT): Identificador único del usuario.
     - name (TEXT): Nombre del usuario.
     - email (TEXT): Dirección de correo electrónico del usuario.
   - Comentario: Tabla de usuario.

2. **accounts**: Tabla de cuentas de usuario.
   - Columnas:
     - id (UUID): Identificador único de la cuenta.
     - balance (BIGINT): Saldo de la cuenta.
     - userid (BIGINT): ID de usuario asociado a la cuenta.
     - active (BOOLEAN): Indica si la cuenta está activa o no.
   - Comentario: Tabla de cuentas de usuario.

3. **transactions**: Tabla para almacenar datos de transacciones.
   - Columnas:
     - id (UUID): Identificador único de la transacción.
     - accountid (UUID): ID de la cuenta asociada a la transacción.
     - amount (FLOAT): Monto de la transacción.
     - operation (VARCHAR(255)): Operación de la transacción.
     - date (TIMESTAMP): Fecha de la transacción.
     - created_at (TIMESTAMP): Fecha y hora en que se creó la transacción.
     - origin (VARCHAR(255)): Origen de la transacción.
   - Comentario: Tabla para almacenar datos de transacciones.

## Relaciones

La base de datos tiene las siguientes relaciones:

- La tabla **accounts** tiene una relación de clave foránea con la tabla **users** mediante la columna **userid**.
  - Constraint: fk_account_user
  - Clave foránea: userid (accounts) -> id (users)
  - Acción en eliminación: ON DELETE CASCADE

- La tabla **transactions** tiene una relación de clave foránea con la tabla **accounts** mediante la columna **accountid**.
  - Constraint: fk_transaction_account
  - Clave foránea: accountid (transactions) -> id (accounts)
  - Acción en eliminación: ON DELETE CASCADE

## Índices

La base de datos tiene los siguientes índices:

- Índice en la tabla **transactions**:
  - Nombre: idx_transactions_account_operation
  - Columnas: accountid, operation

- Índice en la tabla **transactions**:
  - Nombre: idx_transactions_origin
  - Columnas: origin

## Notas

A continuación, se presentan algunas notas adicionales sobre las columnas y tablas:

- En la tabla **users**, la columna **id** representa el identificador único del usuario.
- En la tabla **accounts**, la columna **id** es el identificador único de la cuenta.
- En la tabla **transactions**, la columna **id** es el identificador único de la transacción.
- La columna **balance** en la tabla **accounts** representa el saldo de la cuenta.
- La columna **amount** en la tabla **transactions** representa el monto de la transacción.
- La columna **operation** en la tabla **transactions** representa la operación realizada en la transacción.
- La columna **date** en la tabla **transactions** representa la fecha de la transacción.
- La columna **created_at** en la tabla **transactions** representa la fecha y hora en que se creó la transacción.
- La columna **origin** en la tabla **transactions** representa el origen de la transacción.

---

Este README proporciona una descripción y detalles sobre la estructura de la base de datos utilizada en el proyecto. Para obtener más información sobre el proyecto, consulte el archivo README.md en la raíz del repositorio.