/* !!
 * Code generated by fkit
 * If you need to write additional functions, should create a separate file.
 * File Created: Sunday, 27 April 2025 16:23:20 +07:00
 * Author: VietLD (leduyviet2612@gmail.com)
 * -----
 * Last Modified: Sunday, 27 April 2025 16:23:20 +07:00
 * Modified By: HIEUBV\hieubv
 * -----
 * Copyright 2025. All rights reserved.
 */

package postgres_dao

import (
	sq "github.com/Masterminds/squirrel"

	"github.com/bachhieu/fountain/baselib/sql_client"
	"github.com/bachhieu/fountain/baselib/v_log"
	"github.com/bachhieu/test/biz/dal/models"
	"github.com/jmoiron/sqlx"
)

var (
	product_cates_table  = "product_categories"
	product_cates_fields = []string{}
)

func productCatesToMap(record *models.ProductCateMD) map[string]interface{} {
	return map[string]interface{}{
		"product_id":  record.ProductID,
		"category_id": record.CateID,
	}
}

type ProductCatesDAO struct {
	*sqlx.DB
}

func NewProductCatesDAO(DB *sqlx.DB) *ProductCatesDAO {
	return &ProductCatesDAO{DB}
}

func (dao *ProductCatesDAO) Insert(record *models.ProductCateMD, opts ...Option) error {
	builder := sq.Insert(product_cates_table).SetMap(productCatesToMap(record))

	sqlQuery, args, err := builder.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		v_log.V(1).WithError(err).Errorf("ProductCatesDAO::Insert - Lỗi xây dựng truy vấn: %+v", err)
		return err
	}

	_, err = ExecWithTx(dao.DB, opts, sqlQuery, args...)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("ProductCatesDAO::Insert - Lỗi thực thi truy vấn: %+v", err)
		return err
	}

	return nil
}

func (dao *ProductCatesDAO) BatchInsert(records []*models.ProductCateMD, opts ...Option) error {
	builder := sq.Insert(product_cates_table).
		Columns(product_cates_fields...)
	for _, record := range records {
		builder = builder.SetMap(productCatesToMap(record))
	}

	sqlQuery, args, err := builder.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		v_log.V(1).WithError(err).Errorf("ProductCatesDAO::BatchInsert - Lỗi xây dựng truy vấn: %+v", err)
		return err
	}

	if _, err := ExecWithTx(dao.DB, opts, sqlQuery, args...); err != nil {
		v_log.V(1).WithError(err).Errorf("ProductCatesDAO::BatchInsert - Lỗi thực thi truy vấn: %+v", err)
		return err
	}

	return nil
}

// func (dao *ProductCatesDAO) BatchUpsert(records []*models.ProductCateMD, opts ...Option) error {
// 	builder := sq.Insert(product_cates_table).
//         Columns(product_cates_fields...)

// 	for _, record := range records {
// 		builder = builder.Values(
// 			record.ID
// 		)
// 	}

// 	//TODO: Sửa lại ON CONFLICT sẽ dựa trên các trường cần thiết
// 	builder = builder.Suffix(`ON CONFLICT (id) DO UPDATE SET
// 		name = EXCLUDED.name`)

// 	sqlQuery, args, err := builder.PlaceholderFormat(sq.Dollar).ToSql()
// 	if err != nil {
// 		v_log.V(1).WithError(err).Errorf("ProductCatesDAO::BatchUpsert - Lỗi xây dựng truy vấn: %+v", err)
// 		return err
// 	}

// 	if _, err := ExecWithTx(dao.DB, opts, sqlQuery, args...); err != nil {
// 		v_log.V(1).WithError(err).Errorf("ProductCatesDAO::BatchUpsert - Lỗi thực thi truy vấn: %+v", err)
// 		return err
// 	}

// 	return nil
// }

func (dao *ProductCatesDAO) Upsert(record *models.ProductCateMD, opts ...Option) error {
	builder := sq.Insert(product_cates_table).SetMap(productCatesToMap(record)).
		Suffix(`ON CONFLICT (id) DO UPDATE SET
		name = EXCLUDED.name`)

	sqlQuery, args, err := builder.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		v_log.V(1).WithError(err).Errorf("ProductCatesDAO::Upsert - Lỗi xây dựng truy vấn: %+v", err)
		return err
	}

	_, err = ExecWithTx(dao.DB, opts, sqlQuery, args...)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("ProductCatesDAO::Upsert - Lỗi thực thi truy vấn: %+v", err)
		return err
	}

	return nil
}

func (dao *ProductCatesDAO) Get(id string, opts ...Option) *models.ProductCateMD {
	queryBuilder := sq.Select(product_cates_fields...).
		From(product_cates_table).
		Where(sq.Eq{"id": id})

	query, args, err := queryBuilder.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		v_log.V(1).WithError(err).Errorf("ProductCatesDAO::Get - Lỗi xây dựng truy vấn: %+v", err)
		return nil
	}

	do, err := sql_client.QueryDataParser[models.ProductCateMD](dao.DB, query, nil, args...)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("ProductCatesDAO::Get - Lỗi thực thi truy vấn: %+v", err)
		return nil
	}

	return do
}

func (dao *ProductCatesDAO) GetMany(ids []string, opts ...Option) ([]*models.ProductCateMD, error) {
	queryBuilder := sq.Select(product_cates_fields...).
		From(product_cates_table).
		Where(sq.Eq{"id": ids})

	query, args, err := queryBuilder.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		v_log.V(1).WithError(err).Errorf("ProductCatesDAO::GetMany - Error building query: %+v", err)
		return nil, err
	}

	res, err := sql_client.QueryListDataParser[models.ProductCateMD](dao.DB, query, nil, args...)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("ProductCatesDAO::GetMany - Error executing query: %+v", err)
		return nil, err
	}

	return res, nil
}

func (dao *ProductCatesDAO) GetAll(opts ...Option) []*models.ProductCateMD {
	queryBuilder := sq.Select(product_cates_fields...).
		From(product_cates_table).
		OrderBy("created_time ASC")

	query, args, err := queryBuilder.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		v_log.V(1).WithError(err).Errorf("ProductCatesDAO::GetAll - Lỗi xây dựng truy vấn: %+v", err)
		return nil
	}

	res, err := sql_client.QueryListDataParser[models.ProductCateMD](dao.DB, query, nil, args...)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("ProductCatesDAO::GetAll - Lỗi thực thi truy vấn: %+v", err)
		return nil
	}

	return res
}

// func (dao *ProductCatesDAO) Gets(params *product_cate_do.ProductCateQueryParams, offset, limit int, opts ...Option) ([]*models.ProductCateMD, error) {
// 	queryBuilder := sq.Select(product_cates_fields...).
// 		From(product_cates_table)

// 	if params.Q != nil && *params.Q != "" {
// 		queryBuilder = queryBuilder.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(*params.Q)+"%")
// 	}

// 	queryBuilder = queryBuilder.
// 		OrderBy("created_time ASC").
// 		Offset(uint64(offset)).
// 		Limit(uint64(limit))

// 	query, args, err := queryBuilder.PlaceholderFormat(sq.Dollar).ToSql()
// 	if err != nil {
// 		v_log.V(1).WithError(err).Errorf("ProductCatesDAO::Gets - Error building query: %+v", err)
// 		return nil, err
// 	}

// 	res, err := sql_client.QueryListDataParser[models.ProductCateMD](dao.DB, query, nil, args...)
// 	return res, err
// }

// func (dao *ProductCatesDAO) Update(md *models.ProductCateMD, opts ...Option) error {
// 	queryBuilder := sq.Update(product_cates_table).
// 		SetMap(productCatesToMap(md)).
// 		Where(sq.Eq{"id": md.ID})

// 	sqlQuery, args, err := queryBuilder.PlaceholderFormat(sq.Dollar).ToSql()
// 	if err != nil {
// 		v_log.V(1).WithError(err).Errorf("ProductCatesDAO::Update - Lỗi xây dựng truy vấn: %+v", err)
// 		return err
// 	}

// 	_, err = ExecWithTx(dao.DB, opts, sqlQuery, args...)
// 	if err != nil {
// 		v_log.V(1).WithError(err).Errorf("ProductCatesDAO::Update - Lỗi thực thi truy vấn: %+v", err)
// 		return err
// 	}

// 	return nil
// }
// func (dao *ProductCatesDAO) Delete(id string, opts ...Option) error {
// 	builder := sq.Delete(product_cates_table).
// 		Where(sq.Eq{"id": id})

// 	query, args, err := builder.PlaceholderFormat(sq.Dollar).ToSql()
// 	if err != nil {
// 		v_log.V(1).WithError(err).Errorf("ProductCatesDAO::Delete - Lỗi xây dựng truy vấn: %+v", err)
// 		return err
// 	}

// 	_, err = ExecWithTx(dao.DB, opts, query, args...)
// 	if err != nil {
// 		v_log.V(1).WithError(err).Errorf("ProductCatesDAO::Delete - Lỗi thực thi truy vấn: %+v", err)
// 		return err
// 	}

// 	return nil
// }

// func (dao *ProductCatesDAO) CountTotal(params *product_cate_do.ProductCateQueryParams, opts ...Option) int {
// 	queryBuilder := sq.Select("COUNT(1)").
// 		From(product_cates_table)

// 	if params.Q != nil && *params.Q != "" {
// 		queryBuilder = queryBuilder.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(*params.Q)+"%")
// 	}

// 	query, args, err := queryBuilder.PlaceholderFormat(sq.Dollar).ToSql()
// 	if err != nil {
// 		v_log.V(1).WithError(err).Errorf("ProductCatesDAO::CountTotal - Error building query: %+v", err)
// 		return 0
// 	}

// 	var count int
// 	err = dao.DB.QueryRow(query, args...).Scan(&count)
// 	if err != nil {
// 		v_log.V(1).WithError(err).Errorf("ProductCatesDAO::CountTotal - Error executing query: %+v", err)
// 		return 0
// 	}

// 	return count
// }
