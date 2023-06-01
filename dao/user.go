package dao

import (
	"fmt"
	"test_go_mod/global"
	"test_go_mod/models"
)

var users []models.User
var User models.User

func GetUserListDao(page int, pageSize int) (int, []interface{}) {
	//fmt.Println(int32(time.Now().Unix()))
	userList := make([]interface{}, 0, len(users))
	offset := (page - 1) * pageSize
	result := global.DB.Offset(offset).Limit(pageSize).Find(&users)
	if result.RowsAffected == 0 {
		return 0, userList
	}
	total := len(users)
	result.Offset(offset).Limit(pageSize).Find(&users)
	fmt.Println(users)
	for _, useSingle := range users {
		//birthday := ""
		//if useSingle.Birthday == nil {
		//	birthday = ""
		//} else {
		//	birthday = useSingle.Birthday.Format("2006-01-02")
		//}
		userItemMap := map[string]interface{}{
			"id":       useSingle.ID,
			"password": useSingle.Password,
			"mobile":   useSingle.Mobile,
		}
		userList = append(userList, userItemMap)
	}
	return total, userList
}

func GetUser() interface{} {
	fmt.Println(11111)
	if err := global.DB.First(&User, 1); err != nil {
		panic(err)
		fmt.Println(11)
	}
	fmt.Println(22)

	global.DB.First(&User, 1)
	//fmt.Println(User)
	//fmt.Println(User)
	return User
}

func CreateUser(user models.User) {
	global.DB.Create(&user)
	//global.DB.Commit()
}
