package allParams

// UploadVideoParams 视频上传参数
type UploadVideoParams struct {
	UserId   int    `form:"user_id" json:"user_id" binding:"required"`     //上传者id
	UserName string `form:"user_name" json:"user_name" binding:"required"` //上传者名
	Title    string `form:"title" json:"title" binding:"required"`         //视频名称
	Dsc      string `form:"dsc" json:"dsc"`                                //视频简短描述
}
