# Database connector

### 1.database/sql (low level standard library)

* super-fast.
* manual mapping SQL fields to object.
* errors only show up at runtime.

### 2.GORM

* CRUD already implemented.
* pretty slow on high load.

### 3.sqlx

* quite fast & easy to use.
* fields mapping is done by sqlx.
* errors only show up at runtime.

### 4.sqlc

**best option for our project**

* very fast & easy to use
* automatic code generation.
* catch SQL query errors before generating codes.

#### attention for sqlc:

* only support PostgresSQL for now.
* only support Unix environment. if you are using Windows, please use Docker.
