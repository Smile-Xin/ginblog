package dao

import (
	"fmt"
	. "ginbblog/model"
	"ginbblog/utils"
	"ginbblog/utils/errmsg"
)

// QueryUser id 查询用户
//func QueryUser(id int) (user User, code uint) {
//	result := db.First(&user, id)
//	err := result.Error
//	if err != nil {
//		fmt.Printf("query user fail: %s", err)
//		return
//	}
//	// 不存在用户
//	if result.RowsAffected < 1 {
//		return user, errmsg.INEXISTENCE_USER
//	}
//	return user, errmsg.SUCCESS
//}

// QueryUser name 查询用户
func QueryUser(name string) (user User, code uint) {
	result := db.Where("user_name=?", name).Find(&user)
	err := result.Error
	if err != nil {
		fmt.Printf("query user fail: %s", err)
		return
	}
	// 不存在用户
	if result.RowsAffected < 1 {
		return user, errmsg.INEXISTENCE_USER
	}
	return user, errmsg.SUCCESS
}

//func QueryUserByModel(modelUser User) (userList []User, code uint) {
//	db.Model(&modelUser).Find(userList)
//}

// GetUser 分页获取用户信息
func GetUser(userName string, pageSize int, pageNum int) (user []User, total int64, code uint) {

	if userName != "" {
		// 总人数
		err := db.Where("user_name like ?", "%"+userName+"%").Find(&user).Count(&total).Error
		if err != nil {
			fmt.Printf("get total fail: %s", err)
			return
		}

		err = db.Where("user_name like ?", "%"+userName+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&user).Error
		if err != nil {
			fmt.Printf("query user fail: %s", err)
			return
		}
	} else {
		if pageSize == 0 && pageNum == 0 {
			pageSize = -1
			pageNum = 2
		}
		// 总人数
		err := db.Find(&user).Count(&total).Error
		if err != nil {
			fmt.Printf("get total fail: %s", err)
			return
		}

		// 按页数查
		result := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&user)
		err = result.Error
		if err != nil {
			fmt.Printf("query user fail: %s", err)
			return
		}
	}

	return user, total, errmsg.SUCCESS
}

// name 查重
func ExistUser(name string) bool {
	var user User
	result := db.Where("user_name = ?", name).Find(&user)
	if result.Error != nil {
		fmt.Printf("find user fail%s", result.Error)
	}
	if user.ID <= 0 {
		return false
	} else {
		return true
	}
}

// id 查重
//func ExistUser(id int) bool {
//	var user User
//	result := db.Where("id = ?", id).Find(&user)
//	if result.Error != nil {
//		fmt.Printf("find user fail%s", err)
//	}
//	if result.RowsAffected <= 0 {
//		return false
//	} else {
//		return true
//	}
//}

func AddUser(user *User) (code uint) {
	if ExistUser(user.UserName) {
		code = errmsg.EXIST_USER
	} else {
		err := db.Create(&user).Error
		if err != nil {
			fmt.Printf("creat user fail %s", err)
			code = errmsg.DATABASE_WRITE_FAIL
			return
		}
		code = errmsg.SUCCESS
	}

	return code
}

func EditUser(user *User) (code uint) {
	code = errmsg.SUCCESS
	// 判断是否改用户名
	var uuser User
	db.Select("user_name").Where("id = ?", user.ID).Find(&uuser)
	if uuser.UserName != user.UserName && ExistUser(user.UserName) {
		fmt.Printf("uuser=%s", uuser.UserName)
		fmt.Printf("user=%s", user.UserName)
		fmt.Println(user.UserName != uuser.UserName)
		// 更改了用户名
		// 查重
		code = errmsg.EXIST_USER
		return code
	}
	// 更改数据库
	result := db.Model(&user).Updates(map[string]interface{}{
		"user_name": user.UserName,
		"role":      user.Role,
	})
	if result.Error != nil {
		fmt.Printf("edit user fail:%s", result.Error)
		code = errmsg.DATABASE_WRITE_FAIL
		return
	}
	return code
}

func DeleteUser(name string) (code uint) {
	code = errmsg.SUCCESS
	fmt.Printf("name=%s", name)

	// 判断用户是否存在
	if !ExistUser(name) {
		// 不存在
		code = errmsg.INEXISTENCE_USER
		return
	}

	// 删除用户
	result := db.Where("user_name=?", name).Delete(&User{})
	if result.Error != nil {
		fmt.Printf("delete user fail%s", err)
		code = errmsg.DATABASE_WRITE_FAIL
		return
	}

	return
}

func LoginUser(u *User) (ok bool, code uint) {
	var user User
	result := db.Where("user_name = ?", u.UserName).Find(&user)
	if result.Error != nil {
		fmt.Printf("query user fail:%s", err)
		ok = false
		code = errmsg.DATABASE_WRITE_FAIL
		return
	}
	if !(user.ID > 0) {
		fmt.Printf("can not found user :%s", u.UserName)
		ok = false
		code = errmsg.INEXISTENCE_USER
		return
	}
	if user.Role != 0 {
		ok = false
		code = errmsg.INSUFFICIENT_ROLE
		return
	}
	pw, err := utils.ScryptPW(u.Password)
	if err != nil {
		fmt.Printf("Scrypt password fail:%s", err)
		ok = false
		return
	}
	if user.Password != pw {
		ok = false
		code = errmsg.PASSWORD_ERROR
		return
	}

	// 登陆成功将密码清空防止被窃取
	u.Password = ""
	ok = true
	code = errmsg.SUCCESS
	return

}

// 后台登录
func FrontLogin(u *User) (ok bool, code uint) {
	var user User
	result := db.Where("user_name = ?", u.UserName).Find(&user)
	if result.Error != nil {
		fmt.Printf("query user fail:%s", err)
		ok = false
		code = errmsg.DATABASE_WRITE_FAIL
		return
	}
	if !(user.ID > 0) {
		fmt.Printf("can not found user :%s", u.UserName)
		ok = false
		code = errmsg.INEXISTENCE_USER
		return
	}
	pw, err := utils.ScryptPW(u.Password)
	if err != nil {
		fmt.Printf("Scrypt password fail:%s", err)
		ok = false
		return
	}
	if user.Password != pw {
		ok = false
		code = errmsg.PASSWORD_ERROR
		return
	}

	// 登陆成功将密码清空防止被窃取
	u.Password = ""
	ok = true
	code = errmsg.SUCCESS
	return
}

// 验证权限
func ExamineRole(name string, role int) bool {
	var user User
	role = 0
	if db.Where("user_name = ?", name).Find(&user).Error != nil {
		fmt.Println("query user role fail")
		return false
	}
	if role != user.Role {
		return false
	}
	return true
}
