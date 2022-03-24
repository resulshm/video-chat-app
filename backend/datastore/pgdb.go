package datastore

import (
	"context"
	"video-chat-app/utils"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

type PgAccess struct {
	Access,
	pool *pgxpool.Pool
}

type pgxWithTx func(tx pgx.Tx) (rollback bool, err error)
type pgxQuery func(conn *pgxpool.Conn) (err error)

func NewPgAccess(conf *utils.Config) (pg *PgAccess, err error) {
	var pool *pgxpool.Pool
	pool, err = pgxpool.Connect(context.Background(), utils.Conf.DbConn)
	if err != nil {
		eMsg := "error creating connection pool"
		log.WithError(err).Error(eMsg)
		err = errors.Wrap(err, eMsg)
		return
	}
	pg = &PgAccess{pool: pool}
	return
}

func (d *PgAccess) runInTx(ctx context.Context, pTx pgx.Tx, clog *log.Entry, f pgxWithTx) (err error) {
	var conn *pgxpool.Conn
	defer func() {
		if conn != nil {
			conn.Release()
		}
	}()
	rollback := true
	var tx pgx.Tx
	defer func() {
		if rollback && tx != nil {
			rErr := tx.Rollback(ctx)
			if rErr != nil && err != pgx.ErrTxClosed {
				clog.WithError(rErr).Error("Error in tx.Rollback")
			}
		}
	}()
	if pTx == nil {
		conn, err = d.pool.Acquire(ctx)
		if err != nil {
			clog.WithError(err).Error("error acquiring connection")
			return
		}
		tx, err = conn.Begin(ctx)
		if err != nil {
			eMsg := "Error in conn.Begin"
			clog.WithError(err).Error(eMsg)
			return errors.Wrap(err, eMsg)
		}
	} else {
		tx, err = pTx.Begin(ctx)
		if err != nil {
			eMsg := "Error in tx.Begin"
			clog.WithError(err).Error(eMsg)
			return errors.Wrap(err, eMsg)
		}
	}
	rollback, err = f(tx)
	if err != nil {
		eMsg := "error in executing f"
		clog.WithError(err).Error(eMsg)
		return errors.Wrap(err, eMsg)
	}
	if !rollback {
		err = tx.Commit(ctx)
		if err != nil {
			eMsg := "Error in tx.Commit"
			clog.WithError(err).Error(eMsg)
			return errors.Wrap(err, eMsg)
		}
	}
	return
}

func (d *PgAccess) runQuery(ctx context.Context, clog *log.Entry, f pgxQuery) (err error) {
	var conn *pgxpool.Conn
	defer func() {
		if conn != nil {
			conn.Release()
		}
	}()
	conn, err = d.pool.Acquire(ctx)
	if err != nil {
		clog.WithError(err).Error("error acquiring connection")
		return
	}

	err = f(conn)
	if err != nil {
		eMsg := "error in executing f"
		clog.WithError(err).Error(eMsg)
		return errors.Wrap(err, eMsg)
	}
	return
}
