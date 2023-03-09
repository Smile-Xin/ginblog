package errmsg

const (
	SUCCESS = 200 // 访问成功

	// user 业务代码

	// TOKEN_FAIL token错误
	TOKEN_FAIL = 1010
	// token不存在
	INEXISTENCE_TOKEN = 1011
	// token格式错误
	TOKEN_TYPE_ERROR = 1012
	// token解析错误
	TOKEN_PARSE_ERROR = 1013

	// EXIST_USER 已存在用户
	EXIST_USER = 1001
	// INEXISTENCE_USER 不存在用户
	INEXISTENCE_USER = 1002
	// 密码错误
	PASSWORD_ERROR = 1003
	// 用户权限不足
	INSUFFICIENT_ROLE = 1004

	// category 业务代码

	// EXIST_CATEGORY 已存在分类
	EXIST_CATEGORY = 2001
	// INEXISTENCE_CATEGORY 不存在分类
	INEXISTENCE_CATEGORY = 2002

	// article 业务代码

	// EXIST_ARTICLE 已存在
	EXIST_ARTICLE = 3001
	// INEXISTENCE_ARTICLE 不存在
	INEXISTENCE_ARTICLE = 3002

	// DATABASE_WRITE_FAIL 操作数据库错误
	DATABASE_WRITE_FAIL = 4396
	// TRANSPORT_ERR 网络传输错误
	TRANSPORT_ERR = 4004
	// 七牛云上传错误
	QN_UPLOAD_ERROR = 5001
)

var codemsg = map[uint]string{
	SUCCESS: "OK", // 200

	TOKEN_FAIL:        "token错误",   // 1010
	INEXISTENCE_TOKEN: "token不存在",  // 1011
	TOKEN_TYPE_ERROR:  "token格式错误", // 1012
	TOKEN_PARSE_ERROR: "token解析错误", // 1013
	EXIST_USER:        "已存在用户",     // 1001
	INEXISTENCE_USER:  "不存在用户",     // 1002
	PASSWORD_ERROR:    "密码错误",      // 1003
	INSUFFICIENT_ROLE: "权限不足",      // 1004

	EXIST_CATEGORY:       "已存在分类", // 2001
	INEXISTENCE_CATEGORY: "不存在分类", // 2002

	EXIST_ARTICLE:       "重复文章",  // 3001
	INEXISTENCE_ARTICLE: "不存在文章", // 3002

	DATABASE_WRITE_FAIL: "操作数据库错误",  // 4396
	TRANSPORT_ERR:       "网络传输错误",   // 4004
	QN_UPLOAD_ERROR:     "七牛上传文件失败", // 5001
}

func GetErrMsg(code uint) string {
	return codemsg[code]
}
