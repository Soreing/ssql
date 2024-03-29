package ssql

import (
	"context"
	"database/sql"

	"github.com/Soreing/easyscan"
)

type DB interface {
	Get(
		ctx context.Context,
		dest easyscan.ScanOne,
		query string,
		params ...interface{},
	) (err error)
	Select(
		ctx context.Context,
		dest easyscan.ScanMany,
		query string,
		params ...interface{},
	) (err error)
	Exec(
		ctx context.Context,
		query string,
		params ...interface{},
	) (res sql.Result, err error)
	Begin(
		ctx context.Context,
	) (TX, error)
	Beginx(
		ctx context.Context,
		opt *sql.TxOptions,
	) (TX, error)
	UseHook(
		fn func(context.Context, QueryContext, error),
	)
}

type TX interface {
	Get(
		ctx context.Context,
		dest easyscan.ScanOne,
		query string,
		params ...interface{},
	) (err error)
	Select(
		ctx context.Context,
		dest easyscan.ScanMany,
		query string,
		params ...interface{},
	) (err error)
	Exec(
		ctx context.Context,
		query string,
		params ...interface{},
	) (res sql.Result, err error)
	Commit(
		ctx context.Context,
	) error
	Rollback(
		ctx context.Context,
	) error
}
