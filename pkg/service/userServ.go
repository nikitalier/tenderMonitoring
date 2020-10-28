package service

import (
	"github.com/nikitalier/tenderMonitoring/pkg/models"
)

//FindUserByLogin ...
func (s *Service) FindUserByLogin(login string) (user models.User) {
	user = s.repository.GetUserByLogin(login)
	user.Password = ""
	return user
}

//GetUserRoles ...
func (s *Service) GetUserRoles(id int) []string {
	roleStruct := s.repository.GetUserRolesByID(id)
	n := len(roleStruct)

	roles := make([]string, n)

	for i := 0; i < n; i++ {
		roles[i] = roleStruct[i].Name
	}

	return roles
}

//GetAllUsers ...
func (s *Service) GetAllUsers() (users []models.User) {
	users = s.repository.GetAllUsers()
	return users
}
