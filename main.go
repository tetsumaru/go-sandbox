package main

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

// Customer は、顧客
type Customer struct {
	CustomerID int    `db:"customer_id"`
	Name       string `db:"name"`
	Address    string `db:"address"`
}

// CustomerKey は、顧客のキー
type CustomerKey struct {
	CustomerID int `db:"customer_id"`
}

func main() {
	dsn := os.Getenv("DSN")
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Printf("sql.Open error %s", err)
	}

	key := CustomerKey{1}
	src := Customer{key.CustomerID, "Shohei Otani", "Los Angeles Angels"}

	_, err = db.NamedExec(`
		INSERT INTO CUSTOMER (
			CUSTOMER_ID, NAME, ADDRESS
		) VALUES (
			:customer_id, :name, :address)`,
		src,
	)
	if err != nil {
		log.Printf("db.Exec error %s", err)
	}

	dst := Customer{}
	err = db.QueryRowx(`
		SELECT
			CUSTOMER_ID, NAME, ADDRESS
		FROM CUSTOMER
		WHERE CUSTOMER_ID = $1`,
		key.CustomerID,
	).StructScan(
		&dst,
	)
	if err != nil {
		log.Printf("db.QueryRow error %s", err)
	}
	log.Printf("\nsrc = %#v\ndst = %#v\n", src, dst)

	_, err = db.NamedExec(`
		DELETE FROM CUSTOMER
		WHERE CUSTOMER_ID = :customer_id`,
		key,
	)
	if err != nil {
		log.Printf("db.Exec error %s", err)
	}
}
