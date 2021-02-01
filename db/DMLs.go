package db

const (
	GetUserByLogin = `SELECT * FROM users WHERE login = ($1)`
	SignUp         = `INSERT INTO users(name, surname, login, password, email, phone, status) VALUES($1, $2, $3, $4, $5, $6, $7)`
	GetLogin       = `select u.id, u.name, u.surname, u.login, u.password, u.phone, u.email, r.name, u.status, u.created_at from userrole ur, users u, roles r where u.id = ur.user_id and r.id = ur.role_id and u.login = ($1)`
	AddUserRole = `INSERT INTO userRole(role_id, user_id) SELECT ($1), id FROM users WHERE login = (&2)`
)
