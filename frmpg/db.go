package frmpg

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

func DoTxDefault(db *pgxpool.Pool, f func(tx pgx.Tx) error) (err error) {
	var tx pgx.Tx
	tx, err = db.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return
	}
	defer func() {
		p := recover()
		if p != nil {
			err = tx.Rollback(context.Background())
			if err != nil {
				log.Println("[ERROR] rollback db transaction failed.", err)
			}
			return
		} else {
			err = tx.Commit(context.Background())
			if err != nil {
				log.Println("[ERROR] commit db transaction failed.", err)
			}
			return
		}
	}()

	err = f(tx)
	return
}
