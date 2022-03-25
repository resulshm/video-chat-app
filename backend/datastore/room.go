package datastore

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

const (
	sqlCreateRoom = `INSERT INTO tbl_room(name) VALUES ($1) RETURNING id`
)

func (d *PgAccess) CreateRoom(ctx context.Context, name string) (roomID string, err error) {

	err = d.runInTx(ctx, nil, nil, func(tx pgx.Tx) (rollback bool, err error) {
		rollback = true
		row := tx.QueryRow(ctx, sqlCreateRoom, name)
		err = row.Scan(&roomID)
		if err != nil {
			eMsg := "Couldn't insert room to database"
			logrus.WithError(err).Error(eMsg)
			return
		}

		return false, nil
	})
	return
}
