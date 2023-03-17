package dataDump

import (
	"ScheduleJobs/initializer"
	"ScheduleJobs/model"
)

func usersData() []model.User{
	user1 := model.User{
		Name: "abc",
		Email: "abc@example.com",
		Password: "Abc_3",
	}
	user2 := model.User{
		Name: "abcd",
		Email: "abcd@example.com",
		Password: "Abcd_3",
	}
	user3 := model.User{
		Name: "xyz",
		Email: "XYZ@example.com",
		Password: "Xyz_3",
	}
	user4 := model.User{
		Name: "pqr",
		Email: "pqr@example.com",
		Password: "pqr_3",
	}
	user5 := model.User{
		Name: "pqrt",
		Email: "pqrt@example.com",
		Password: "Pqrt_3",
	}

	usersData := []model.User{user1, user2, user3, user4, user5}

	return usersData
}

func addToUsersTable(usersData []model.User){
	for _, user := range usersData{
		if initializer.Db.Where("email=?",user.Email).Find(&user).RowsAffected == 1 {continue}
		initializer.Db.Create(&user)
	}
}

func UserData() {
	addToUsersTable(usersData())
}