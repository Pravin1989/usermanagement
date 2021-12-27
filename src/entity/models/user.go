package models

import "time"

type User struct {
	UserId
	UserRequest
	ChangeTrackingSupport
}

type UserRequest struct {
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	Salary      float32 `json:"salary"`
	Address     string  `json:"address"`
	Designation string  `json:"designation"`
}

type ChangeTrackingSupportUser struct {
	CreatedBy string
	UpdatedBy string
}

type ChangeTrackingSupportDate struct {
	CreatedTime time.Time `json:"createdTime"`
	UpdatedTime time.Time `json:"updatedTime"`
}

type UserId struct {
	ID string
}

type ChangeTrackingSupport struct {
	ChangeTrackingSupportUser
	ChangeTrackingSupportDate
}
