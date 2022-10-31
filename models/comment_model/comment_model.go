package commentmodel

type CommentModel struct {
	CommentId string `json:"_id" bson:"_id"`

	VideoId      string `json:"videoId" bson:"videoId" binding:"required"`
	UserName     string `json:"userName" bson:"userName" binding:"required"`
	Comment      string `json:"comment" bson:"comment" binding:"required"`
	UserId       string `json:"userId" bson:"userId" binding:"required"`
	CreatedAt    string `json:"createdAt" bson:"createdAt" binding:"required"`
	IsReply      bool   `json:"isReply" bson:"isReply"`
	TotalReplies int64  `json:"totalReplies" bson:"totalReplies"`
	Likes        int64  `json:"likes" bson:"likes"`
	Dislikes     int64  `json:"dislikes" bson:"dislikes"`
}

func (c CommentModel) ToMap() map[string]interface{} {

	return map[string]interface{}{
		"videoId":      c.VideoId,
		"userName":     c.UserName,
		"comment":      c.Comment,
		"userId":       c.UserId,
		"createdAt":    c.CreatedAt,
		"isReply":      c.IsReply,
		"totalReplies": c.TotalReplies,
		"likes":        c.IsReply,
		"dislikes":     c.Dislikes,
	}

}
