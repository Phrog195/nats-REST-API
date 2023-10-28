package postgresql

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dsn := "url"
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal(err)
	}
}
