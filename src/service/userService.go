package service

import (
	"ecommerce/user/usermanagement/src/db"
	"ecommerce/user/usermanagement/src/entity/models"
	"ecommerce/user/usermanagement/src/utilities"
	"fmt"
	"time"
)

const (
	SelectQuery = "SELECT * FROM user_details WHERE user_id= $1"
	insertUser  = `INSERT INTO user_details("user_id", "first_name", "last_name", "salary",
	 "address", "designation", "created_by", "updated_by", "created_at", "updated_at")
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10);`
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

func SaveUserDetails(user models.User) (models.User, error) {
	dbCon := db.GetPostgresDBConnection()
	user.UserId = models.UserId{ID: utilities.GenerateUUID().String()}
	user.CreatedTime = time.Now()
	user.UpdatedTime = time.Now()
	tx := dbCon.DB.MustBegin()
	result, saveErr := tx.Exec(insertUser,
		user.UserId.ID,
		user.FirstName,
		user.LastName,
		user.Salary,
		user.Address,
		user.Designation,
		user.CreatedBy,
		user.UpdatedBy,
		user.CreatedTime,
		user.UpdatedTime)
	if saveErr != nil {
		return models.User{}, saveErr
	}
	tx.Commit()
	Id, _ := result.LastInsertId()
	fmt.Println("Last inserted index :", Id)
	return user, nil
}
