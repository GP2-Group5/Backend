package delivery

import (
	"github.com/GP2-Group5/Backend/feature/users"
)

type UserReq struct {
	Full_Name string `json:"full_name" form:"full_name"`
	Password  string `json:"password" form:"password"`
	Email     string `json:"email" form:"email"`
	Team_id   int    `json:"team_id" form:"team_id"`
	Status    string `json:"status" form:"status"`
	Role      string `json:"role_id" form:"role_id"`
}

func toCore(data UserReq) users.UserCore {
	return users.UserCore{
		Full_Name: data.Full_Name,
		Email:     data.Email,
		Password:  data.Password,
		Role:      data.Role,
		TeamID:    uint(data.Team_id),
	}
}
