package service

import (
	"ecommerce/user/usermanagement/src/db"
	"ecommerce/user/usermanagement/src/entity/models"
)

const (
	SelectQuery = "SELECT * FROM user_details WHERE user_id= $1"
)

func GetUserDetailsById(userId models.UserId) (models.User, error) {
	dbCon := db.GetPostgresDBConnection()
	user := models.User{}
	fetchErr := dbCon.DB.Get(&user, SelectQuery, userId.ID)
	if fetchErr != nil {
		return models.User{}, fetchErr
	}
	return user, nil
}
