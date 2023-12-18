// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: queries.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const deleteAllProducts = `-- name: DeleteAllProducts :exec
DELETE FROM product_meta
`

func (q *Queries) DeleteAllProducts(ctx context.Context) error {
	_, err := q.db.Exec(ctx, deleteAllProducts)
	return err
}

const getAllProducts = `-- name: GetAllProducts :many
SELECT id, platform, symbol, locale, market, name, description FROM product_meta
`

type GetAllProductsRow struct {
	ID          string
	Platform    Platform
	Symbol      string
	Locale      Locale
	Market      Market
	Name        pgtype.Text
	Description pgtype.Text
}

func (q *Queries) GetAllProducts(ctx context.Context) ([]GetAllProductsRow, error) {
	rows, err := q.db.Query(ctx, getAllProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllProductsRow
	for rows.Next() {
		var i GetAllProductsRow
		if err := rows.Scan(
			&i.ID,
			&i.Platform,
			&i.Symbol,
			&i.Locale,
			&i.Market,
			&i.Name,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProductById = `-- name: GetProductById :one
SELECT id, platform, symbol, locale, market, name, description FROM product_meta
WHERE id = $1
`

type GetProductByIdRow struct {
	ID          string
	Platform    Platform
	Symbol      string
	Locale      Locale
	Market      Market
	Name        pgtype.Text
	Description pgtype.Text
}

func (q *Queries) GetProductById(ctx context.Context, id string) (GetProductByIdRow, error) {
	row := q.db.QueryRow(ctx, getProductById, id)
	var i GetProductByIdRow
	err := row.Scan(
		&i.ID,
		&i.Platform,
		&i.Symbol,
		&i.Locale,
		&i.Market,
		&i.Name,
		&i.Description,
	)
	return i, err
}

const getProductsByCondition = `-- name: GetProductsByCondition :many
SELECT id, platform, symbol, locale, market, name, description FROM product_meta
WHERE platform = $1 AND market = $2
`

type GetProductsByConditionParams struct {
	Platform Platform
	Market   Market
}

type GetProductsByConditionRow struct {
	ID          string
	Platform    Platform
	Symbol      string
	Locale      Locale
	Market      Market
	Name        pgtype.Text
	Description pgtype.Text
}

func (q *Queries) GetProductsByCondition(ctx context.Context, arg GetProductsByConditionParams) ([]GetProductsByConditionRow, error) {
	rows, err := q.db.Query(ctx, getProductsByCondition, arg.Platform, arg.Market)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetProductsByConditionRow
	for rows.Next() {
		var i GetProductsByConditionRow
		if err := rows.Scan(
			&i.ID,
			&i.Platform,
			&i.Symbol,
			&i.Locale,
			&i.Market,
			&i.Name,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

type InsertProductsParams struct {
	ID          string
	Platform    Platform
	Symbol      string
	Locale      Locale
	Market      Market
	Name        pgtype.Text
	Description pgtype.Text
}
