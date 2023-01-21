package psqldb

import (
	"database/sql"
	"fmt"

	"github.com/martinconic/eth-events-indexer/config"
)

type Database struct {
	Config *config.PostgresConfig
	DB     *sql.DB
}

const psqlInfo = "host=%s port=%d user=%s " +
	"password=%s dbname=%s sslmode=disable"

func NewDatabase(config *config.PostgresConfig) (*Database, error) {
	db := &Database{
		Config: config,
	}
	error := db.getDBConnection()

	return db, error
}

func (d *Database) getDBConnection() error {
	var err error

	if d.DB != nil {
		return err
	}

	dataSource := fmt.Sprintf(psqlInfo, d.Config.Host, d.Config.Port, d.Config.User,
		d.Config.Password, d.Config.Name)

	d.DB, err = sql.Open("postgres", dataSource)
	if err != nil {
		panic(err)
	}

	return err

}
