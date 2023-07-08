package pkg

import (
	"database/sql"
	"log"
	"st/ent"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var client *ent.Client

func InitEnt()  {
	c := open("postgresql://postgres:@127.0.0.1/starter_template")
	client = c
}

func EntClient() *ent.Client {
	return client
}


// Open new connection
func open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
			log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}
