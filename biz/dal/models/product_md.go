package models

import "time"

type ProductMD struct {
	ID            string  `db:"id" json:"id"`
	Name          string  `db:"name" json:"name"`
	Description   string  `db:"description" json:"description"`
	Price         float64 `db:"price" json:"price"`
	StockQuantity int     `db:"stock_quantity" json:"stock_quantity"`
	Display       bool    `db:"display" json:"display"`       // Return or not
	CreatedBy     string  `db:"created_by" json:"created_by"` // ID người tạo
	CreatedTime   int32   `db:"created_time" json:"created_time"`
	UpdatedTime   int32   `db:"updated_time" json:"updated_time"`
}

func (current *ProductMD) AllowedUpdateFields(req *ProductMD) *ProductMD {
	return &ProductMD{
		ID:            current.ID,
		Name:          req.Name,
		Description:   req.Description,
		Price:         req.Price,
		StockQuantity: req.StockQuantity,
		Display:       req.Display,
		CreatedBy:     req.CreatedBy,
		CreatedTime:   current.CreatedTime,
		UpdatedTime:   int32(time.Now().Unix()),
	}
}
