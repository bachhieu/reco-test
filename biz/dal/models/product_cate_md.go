package models

type ProductCateMD struct {
	ProductID string `db:"product_id" json:"product_id"`
	CateID    string `db:"category_id" json:"category_id"`
}

func (current *ProductCateMD) AllowedUpdateFields(req *ProductCateMD) *ProductCateMD {
	return &ProductCateMD{
		ProductID: req.ProductID,
		CateID:    req.CateID,
	}
}
