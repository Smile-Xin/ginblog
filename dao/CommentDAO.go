package dao

import (
	. "ginbblog/model"
	"ginbblog/utils/errmsg"
	"gorm.io/gorm"
)

// AddComment 新增评论
func AddComment(data *Comment) uint {
	err = db.Create(&data).Error
	if err != nil {
		return errmsg.DATABASE_WRITE_FAIL
	}
	return errmsg.SUCCESS
}

// GetComment 查询单个评论
func GetComment(id int) (Comment, uint) {
	var comment Comment
	err = db.Where("id = ?", id).First(&comment).Error
	if err != nil {
		return comment, errmsg.DATABASE_WRITE_FAIL
	}
	return comment, errmsg.SUCCESS
}

// GetCommentList 后台所有获取评论列表
func GetCommentList(pageSize int, pageNum int) ([]Comment, int64, uint) {

	var commentList []Comment
	var total int64
	db.Find(&commentList).Count(&total)
	err = db.Model(&commentList).Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Select("comment.id, article.title,user_id,article_id, user.user_name, comment.content, comment.status,comment.created_at,comment.deleted_at").Joins("LEFT JOIN article ON comment.article_id = article.id").Joins("LEFT JOIN user ON comment.user_id = user.id").Scan(&commentList).Error
	if err != nil {
		return commentList, 0, errmsg.DATABASE_WRITE_FAIL
	}
	return commentList, total, errmsg.SUCCESS
}

// GetCommentCount 获取评论数量
func GetCommentCount(id int) int64 {
	var comment Comment
	var total int64
	db.Find(&comment).Where("article_id = ?", id).Where("status = ?", 1).Count(&total)
	return total
}

// GetCommentListFront 展示页面获取评论列表
func GetCommentListFront(id int, pageSize int, pageNum int) ([]Comment, int64, uint) {
	var commentList []Comment
	var total int64
	db.Find(&Comment{}).Where("article_id = ?", id).Where("status = ?", 1).Count(&total)
	err = db.Model(&Comment{}).Limit(pageSize).Offset((pageNum-1)*pageSize).Order("Created_At DESC").Select("comment.id, article.title, user_id, article_id, user.user_name, comment.content, comment.status,comment.created_at,comment.deleted_at").Joins("LEFT JOIN article ON comment.article_id = article.id").Joins("LEFT JOIN user ON comment.user_id = user.id").Where("article_id = ?",
		id).Where("status = ?", 1).Scan(&commentList).Error
	if err != nil {
		return commentList, 0, errmsg.DATABASE_WRITE_FAIL
	}
	return commentList, total, errmsg.SUCCESS
}

// 编辑评论（暂不允许编辑评论）

// DeleteComment 删除评论
func DeleteComment(id uint) uint {
	var comment Comment
	err = db.Where("id = ?", id).Delete(&comment).Error
	if err != nil {
		return errmsg.DATABASE_WRITE_FAIL
	}
	return errmsg.SUCCESS
}

// CheckComment 通过评论
func CheckComment(id int, data *Comment) uint {
	var comment Comment
	var res Comment
	var article Article
	var maps = make(map[string]interface{})
	maps["status"] = data.Status

	err = db.Model(&comment).Where("id = ?", id).Updates(maps).First(&res).Error
	db.Model(&article).Where("id = ?", res.ArticleId).UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1))
	if err != nil {
		return errmsg.DATABASE_WRITE_FAIL
	}
	return errmsg.SUCCESS
}

// UncheckComment 撤下评论
func UncheckComment(id int, data *Comment) uint {
	var comment Comment
	var res Comment
	var article Article
	var maps = make(map[string]interface{})
	maps["status"] = data.Status

	err = db.Model(&comment).Where("id = ?", id).Updates(maps).First(&res).Error
	db.Model(&article).Where("id = ?", res.ArticleId).UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1))
	if err != nil {
		return errmsg.DATABASE_WRITE_FAIL
	}
	return errmsg.SUCCESS
}
