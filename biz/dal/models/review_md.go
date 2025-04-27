package models

import "time"

type ReviewMD struct {
	ID          string `db:"id" json:"id"`
	ProductID   string `db:"product_id" json:"product_id"`
	UserID      string `db:"user_id" json:"user_id"`
	Rating      int    `db:"rating" json:"rating"`
	Comment     string `db:"comment" json:"comment"`
	CreatedTime int32  `db:"created_time" json:"created_time"`
	UpdatedTime int32  `db:"updated_time" json:"updated_time"`
}

func (current *ReviewMD) AllowedUpdateFields(req *ReviewMD) *ReviewMD {
	return &ReviewMD{
		ID:          current.ID,
		ProductID:   req.ProductID,
		UserID:      req.UserID,
		Rating:      req.Rating,
		Comment:     req.Comment,
		CreatedTime: current.CreatedTime,
		UpdatedTime: int32(time.Now().Unix()),
	}
}
