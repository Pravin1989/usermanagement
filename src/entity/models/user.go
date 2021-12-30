package models

import (
	"ecommerce/user/usermanagement/src/utilities"
	"time"
)

type User struct {
	UserId
	UserRequest
	ChangeTrackingSupport
}

type UserRequest struct {
	FirstName   string  `json:"firstName" db:"first_name"`
	LastName    string  `json:"lastName" db:"last_name"`
	Salary      float32 `json:"salary" db:"salary"`
	Address     string  `json:"address" db:"address"`
	Designation string  `json:"designation" db:"designation"`
}

type ChangeTrackingSupportUser struct {
	CreatedBy string `json:"createdBy" db:"created_by"`
	UpdatedBy string `json:"updatedBy" db:"updated_by"`
}

type ChangeTrackingSupportDate struct {
	CreatedTime time.Time `json:"createdAt" db:"created_at"`
	UpdatedTime time.Time `json:"updatedAt" db:"updated_at"`
}

type UserId struct {
	ID string `json:"userId" db:"user_id"`
}

type ChangeTrackingSupport struct {
	ChangeTrackingSupportUser
	ChangeTrackingSupportDate
}

func (request *UserRequest) ToUser() User {
	user := User{}
	user.UserId = UserId{utilities.GenerateUUID().String()}
	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.Salary = request.Salary
	user.Address = request.Address
	user.Designation = request.Designation
	user.CreatedBy = "SYSTEM"
	user.UpdatedBy = "SYSTEM"
	return user
}
