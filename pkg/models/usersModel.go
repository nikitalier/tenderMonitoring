package models

//User ...
type User struct {
	ID       int    `db:"id"`
	Login    string `db:"login"`
	FullName string `db:"full_name"`
	Password string `db:"password"`
}
