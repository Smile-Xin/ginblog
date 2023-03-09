package dao

import (
	"fmt"
	"ginbblog/model"
	"ginbblog/utils/errmsg"
)

func GetCategory(pageSize int, pageNum int, cateName string) (categoryList []model.Category, total int64, code uint) {
	if cateName != "" {
		// 请求人数
		err := db.Where("name like ?", "%"+cateName+"%").Find(&categoryList).Count(&total).Error
		if err != nil {
			fmt.Printf("get total fail: %s", err)
			code = errmsg.DATABASE_WRITE_FAIL
			return
		}

		// 分页查找数据
		err = db.Where("name like ?", "%"+cateName+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&categoryList).Error
		if err != nil {
			fmt.Printf("get category fail: %s", err)
			code = errmsg.DATABASE_WRITE_FAIL
			return
		}
	} else {
		// 总分类数
		db.Find(&categoryList).Count(&total)
		// 无筛选获取分类列表
		err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&categoryList).Error
		if err != nil {
			fmt.Printf("get category fail:%s", err)
			code = errmsg.DATABASE_WRITE_FAIL
			return
		}
	}
	code = errmsg.SUCCESS
	return
}

func QueryCategory(name string) (category model.Category, code uint) {
	result := db.Where("name = ?", name).Find(&category)
	// 处理数据库读取错误
	if result.Error != nil {
		code = errmsg.DATABASE_WRITE_FAIL
		fmt.Printf("query category fail %s", err)
		return
	}
	// 不存在用户
	if result.RowsAffected < 1 {
		code = errmsg.INEXISTENCE_CATEGORY
		return
	}
	// 正常查询
	code = errmsg.SUCCESS

	return
}

func AddCategory(category model.Category) (code uint) {
	// 查重
	if ExistCategory(category.Name) {
		code = errmsg.EXIST_CATEGORY
		return
	}

	// 操作数据库
	result := db.Create(&category)
	if result.Error != nil {
		fmt.Printf("add category fail%s", err)
		code = errmsg.DATABASE_WRITE_FAIL
		return
	}
	if result.RowsAffected < 1 {
		fmt.Println("change 0 rows (add category)")
		code = errmsg.DATABASE_WRITE_FAIL
		return
	}

	code = errmsg.SUCCESS

	return
}

func ExistCategory(name string) bool {
	var category model.Category
	result := db.Where("name = ?", name).Find(&category)
	if result.Error != nil {
		fmt.Printf("find user fail%s", err)
	}
	if result.RowsAffected <= 0 {
		return false
	} else {
		return true
	}
}

// EditCategory 修改分类
func EditCategory(category model.Category) (code uint) {
	// 是否存在分类
	//if !ExistCategory(category.Name) {
	//	code = errmsg.INEXISTENCE_CATEGORY
	//	return
	//}

	// 操作数据库
	reult := db.Model(&model.Category{}).Where("id=?", category.ID).Updates(&model.Category{
		Name: category.Name,
	})

	if reult.Error != nil {
		fmt.Printf("edit category fail :%s", err)
		code = errmsg.DATABASE_WRITE_FAIL
		return
	}

	if reult.RowsAffected < 1 {
		fmt.Println("change 0 row (edit category)")
		code = errmsg.DATABASE_WRITE_FAIL
		return
	}
	code = errmsg.SUCCESS
	return
}

func DeleteCategory(name string) (code uint) {
	if !ExistCategory(name) {
		code = errmsg.INEXISTENCE_CATEGORY
	} else {
		// 操作数据库
		result := db.Where("name = ?", name).Delete(&model.Category{})
		// 处理错误
		if result.Error != nil {
			code = errmsg.DATABASE_WRITE_FAIL
			fmt.Printf("delete category fail:%s", err)
			return
		}
		code = errmsg.SUCCESS
	}
	return
}
