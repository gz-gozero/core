package mysq

import (
	"context"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

func GenQuery(query sq.SelectBuilder, param map[string]interface{}) sq.SelectBuilder {
	for key, value := range param {
		parts := strings.Split(key, "__")
		if len(parts) == 1 {
			query = query.Where(sq.Eq{key: value})
			continue
		}

		key = parts[0]
		op := parts[1]

		// GE
		if op == "Min" || op == "GE" {
			query = query.Where(sq.GtOrEq{key: value})
			continue
		}

		// LE
		if op == "Max" || op == "LE" {
			query = query.Where(sq.LtOrEq{key: value})
			continue
		}

		// GT
		if op == "GT" {
			query = query.Where(sq.Gt{key: value})
			continue
		}

		// LT
		if op == "LT" {
			query = query.Where(sq.Lt{key: value})
			continue
		}

		// LK
		if op == "LK" {
			query = query.Where(sq.Like{key: fmt.Sprint("%", value, "%")})
			continue
		}

		// LR
		if op == "LR" {
			query = query.Where(sq.Like{key: fmt.Sprint(value, "%")})
			continue
		}

		// LL
		if op == "LL" {
			query = query.Where(sq.Like{key: fmt.Sprint("%", value)})
			continue
		}

		query = query.Where(sq.Eq{key: value})
	}
	return query
}

func SelectList[T any](ctx context.Context, cc *sqlc.CachedConn, table string, param map[string]interface{}) ([]*T, error) {
	query := sq.Select("t.*").From(fmt.Sprintf("%s AS t", table))
	query = GenQuery(query, param)

	sql, values, err := query.ToSql()

	if err != nil {
		return nil, err
	}

	var resp []*T
	err = cc.QueryRowsNoCacheCtx(ctx, &resp, sql, values...)
	if len(resp) == 0 {
		return nil, err
	}

	return resp, err
}
