package datastore

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

const (
	sqlCreateRoom = `INSERT INTO tbl_room(name) VALUES ($1) RETURNING id`
	sqlGetRoom    = `SELECT id, name FROM tbl_room WHERE id = $1`
)

type Room struct {
	ID   string
	Name string
}

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

func (d *PgAccess) GetRoom(ctx context.Context, roomID string) (room *Room, err error) {
	err = d.runQuery(ctx, nil, func(conn *pgxpool.Conn) (err error) {
		defer func() {
			if err != nil {
				room = nil
			}
		}()
		room = &Room{}

		row := conn.QueryRow(ctx, sqlGetRoom, roomID)
		err = row.Scan(
			&room.ID,
			&room.Name,
		)
		if err != nil {
			if err == pgx.ErrNoRows {
				err = nil
				room = nil
				return
			}
			eMsg := "Couldn't get room from database"
			logrus.WithError(err).Error(eMsg)
			return
		}
		return

	})
	return
}
