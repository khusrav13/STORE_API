package db

const UserTable = `CREATE TABLE IF NOT EXISTS users (
	id BIGSERIAL PRIMARY KEY,
	name VARCHAR(70) NOT NULL,
	surname VARCHAR(70) NOT NULL,
	login VARCHAR(70) NOT NULL,
	password TEXT NOT NULL,
	email VARCHAR(70) NOT NULL UNIQUE,
	phone VARCHAR(70) NOT NULL UNIQUE,
	status BOOLEAN NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);`

const CardTable = `CREATE TABLE IF NOT EXISTS accounts (
	id BIGSERIAL PRIMARY KEY,
	userId BIGSERIAL REFERENCES users(id),
	name VARCHAR(70) NOT NULL,
	number INTEGER NOT NULL,
	balance INTEGER NOT NULL
);`

const CategoryTable = `CREATE TABLE IF NOT EXISTS categories (
	id BIGSERIAL PRIMARY KEY,
	name TEXT UNIQUE NOT NULL,
	status BOOLEAN NOT NULL DEFAULT FALSE
);`

const ProductTable = `CREATE TABLE IF NOT EXISTS products (
	id BIGSERIAL PRIMARY KEY,
	category_id INTEGER REFERENCES categories(id) NOT NULL,
	name TEXT NOT NULL,
	photo TEXT,
	price INTEGER NOT NULL,
	status TEXT NOT NULL
);`

const OrderTable = `CREATE TABLE IF NOT EXISTS orders (
	id BIGSERIAL PRIMARY KEY,
	user_id INTEGER NOT NULL,
	ordered_date TEXT NOT NULL,
	updated_date TEXT,
	product_name TEXT NOT NULL,
	address TEXT NOT NULL,
	status TEXT NOT NULL
);`

const CreateRolesTable = `CREATE TABLE IF NOT EXISTS roles (
	id BIGSERIAL PRIMARY KEY,
	name VARCHAR(70) NOT NULL UNIQUE,
	display_name VARCHAR(60) NOT NULL UNIQUE,
	description VARCHAR(191) NOT NULL UNIQUE
)`

const CreateUserRoleTable = `Create table if not exists userRole (
	role_id INTEGER REFERENCES roles,
	user_id INTEGER REFERENCES users UNIQUE
);`
