package allParams

// ParamLikeData 点赞参数
type ParamLikeData struct {
	VideoId   int `json:"post_id" binding:"required"`            //视频id
	Direction int `json:"direction" binding:"required, oneof=1"` //点赞 1 表示赞
}
