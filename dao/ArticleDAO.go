package dao

import (
	"fmt"
	"ginbblog/model"
	"ginbblog/utils/errmsg"
	"gorm.io/gorm"
)

func AddArticle(article model.Article) (code uint) {
	result := db.Create(&article)
	if result.Error != nil {
		fmt.Printf("create article fail%s", err)
		code = errmsg.DATABASE_WRITE_FAIL
	} else {
		code = errmsg.SUCCESS
	}
	return
}

func QueryArticle(id int) (code uint, article model.Article) {
	// 操作数据库
	result := db.Preload("Category").Where("id = ?", id).Find(&article)
	// 查询时的错误
	if result.Error != nil {
		code = errmsg.DATABASE_WRITE_FAIL
		fmt.Printf("find article fail%s", err)
	} else if result.RowsAffected < 1 {
		// 未查到
		code = errmsg.INEXISTENCE_ARTICLE
	} else {
		code = errmsg.SUCCESS
		db.Model(&article).Where("id = ?", id).UpdateColumn("read_count", gorm.Expr("read_count + ?", 1))

	}
	return
}

func QueryArtList(title string, cid int, pageSize int, pageNum int) (articleList []model.Article, total int64, code uint) {
	if pageSize <= 0 || pageNum <= 0 {
		pageSize = 0
		pageNum = 0
	}
	if title != "" {
		// 符合的文章数
		db.Where("title like ? and cid = ?", "%"+title+"%", cid).Find(&articleList).Count(&total)
		// 按照title查找文章
		err := db.Where("title like ? and cid = ?", "%"+title+"%", cid).Limit(pageSize).Offset((pageNum - 1) * pageNum).Find(&articleList).Error
		if err != nil {
			fmt.Printf("get article fail:%s", err)
			code = errmsg.DATABASE_WRITE_FAIL
			return
		}
		code = errmsg.SUCCESS
	} else {
		// 查询全部文章数
		db.Where("cid = ?", cid).Find(&articleList).Count(&total)
		// 查询全部文章
		err := db.Where("cid = ?", cid).Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Error
		if err != nil {
			fmt.Printf("get article fail:%s", err)
			code = errmsg.DATABASE_WRITE_FAIL
			return
		}
		code = errmsg.SUCCESS
	}
	return
}

func GetArticle(title string, pageSize int, pageNum int) (articleList []model.Article, total int64, code uint) {
	if pageSize <= 0 || pageNum <= 0 {
		pageSize = 0
		pageNum = 0
	}
	if title != "" {
		// 符合的文章数
		db.Where("title like ?", "%"+title+"%").Find(&articleList).Count(&total)
		// 按照title查找文章
		err := db.Where("title like ?", "%"+title+"%").Limit(pageSize).Offset((pageNum - 1) * pageNum).Find(&articleList).Error
		if err != nil {
			fmt.Printf("get article fail:%s", err)
			code = errmsg.DATABASE_WRITE_FAIL
			return
		}
		code = errmsg.SUCCESS
	} else {
		// 查询全部文章数
		db.Find(&articleList).Count(&total)
		// 查询全部文章
		err := db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Error
		if err != nil {
			fmt.Printf("get article fail:%s", err)
			code = errmsg.DATABASE_WRITE_FAIL
			return
		}
		code = errmsg.SUCCESS
	}
	return
}

func EditArticle(article model.Article) (code uint) {
	// 检查是否存在文章
	if !ExistArticle(int(article.ID)) {
		fmt.Printf("文章修改失败，不存在文章")
		code = errmsg.INEXISTENCE_ARTICLE
		return
	}

	fmt.Println(article)

	result := db.Model(&article).Updates(map[string]interface{}{
		"title":   article.Title,
		"cid":     article.Cid,
		"desc":    article.Desc,
		"content": article.Content,
		"img":     article.Img,
	})
	if result.Error != nil {
		fmt.Printf("article updates fail:%s", err)
		code = errmsg.DATABASE_WRITE_FAIL
	} else {
		code = errmsg.SUCCESS
	}
	return
}

func DeleteArticle(id int) (code uint) {
	// 检查是否存在文章
	if !ExistArticle(id) {
		fmt.Printf("文章删除失败，不存在文章")
		code = errmsg.INEXISTENCE_ARTICLE
		return
	}

	result := db.Delete(&model.Article{}, id)
	if result.Error != nil {
		fmt.Printf("delete article fial:%s", err)
		code = errmsg.DATABASE_WRITE_FAIL
	} else {
		code = errmsg.SUCCESS
	}
	return
}

func ExistArticle(id int) bool {
	var article model.Article
	result := db.Where("id = ?", id).Find(&article)
	if result.RowsAffected < 1 {
		// 未查到
		return false
	} else {
		return true
	}
}

//func EditArticle(article model.Article) (code uint) {
//	ExistArticle(article.)
//}
