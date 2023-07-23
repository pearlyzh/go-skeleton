package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go-skeleton/generated/grpc/go_skeleton"
)

type RequestRepository interface {
	SaveRequest(context.Context, *string, *go_skeleton.PingRequest) error
}

type requestRepoImpl struct {
	db *sqlx.DB
}

func NewRequestRepository(db *sqlx.DB) RequestRepository {
	return &requestRepoImpl{
		db,
	}
}

const requestTable = "request"
const insertRequestStr = `INSERT INTO %s (uuid, name) VALUES (?,?)`

func (r *requestRepoImpl) SaveRequest(ctx context.Context, uid *string, request *go_skeleton.PingRequest) error {
	stmt := fmt.Sprintf(insertRequestStr, requestTable)
	_, err := r.db.ExecContext(ctx, stmt, uid, request.GetName())
	return err
}
