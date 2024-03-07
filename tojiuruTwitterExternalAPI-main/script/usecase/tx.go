package usecase

import (
	"context"
	"database/sql"
	"log"
)

// wraper transaction
func Transaction(ctx context.Context, conn *sql.DB, txFunc func(*sql.Tx) (interface{}, error)) (interface{}, error) {
	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer func() {
		if p := recover(); p != nil {
			if err := tx.Rollback(); err != nil {
				panic(p)
			}
		} else if err != nil {
			if err := tx.Rollback(); err != nil {
				log.Println(err)
			}
			return
		}
		err = tx.Commit()
		if err != nil {
			log.Println(err)
		}
	}()
	return txFunc(tx)
}
