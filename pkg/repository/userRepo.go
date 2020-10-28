package repository

import (
	"log"

	"github.com/nikitalier/tenderMonitoring/pkg/models"
)

//GetUserByLogin ...
func (r *Repository) GetUserByLogin(login string) (u models.User) {
	err := r.db.Get(&u, "select * from \"User\" where login=$1", login)
	if err != nil {
		log.Println(err)
	}
	return u
}

//GetUserRolesByID ...
func (r *Repository) GetUserRolesByID(id int) (roles []models.Role) {
	err := r.db.Select(&roles, "select r.name from \"User_role\" ur join \"Role\" r	on ur.role_id  = r.id where ur.user_id = $1", id)
	if err != nil {
		log.Println(err)
	}

	return roles
}

//GetAllUsers ...
func (r *Repository) GetAllUsers() (users []models.User) {
	err := r.db.Select(&users, "select id, login, full_name from \"User\" u ")
	if err != nil {
		log.Println(err)
	}
	return users
}
