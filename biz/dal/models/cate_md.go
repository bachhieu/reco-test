package models

import "time"

type CateMD struct {
	ID          string `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
	CreatedBy   string `db:"created_by" json:"created_by"`
	CreatedTime int32  `db:"created_time" json:"created_time"`
	UpdatedTime int32  `db:"updated_time" json:"updated_time"`
}

func (current *CateMD) AllowedUpdateFields(req *CateMD) *CateMD {
	return &CateMD{
		ID:          current.ID,
		Name:        req.Name,
		Description: req.Description,
		CreatedBy:   current.CreatedBy,
		CreatedTime: current.CreatedTime,
		UpdatedTime: int32(time.Now().Unix()),
	}
}
