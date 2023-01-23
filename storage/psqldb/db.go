package psqldb

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/martinconic/eth-events-indexer/config"
	"github.com/martinconic/eth-events-indexer/storage"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	Config *config.PostgresConfig
	DB     *sql.DB
}

func NewDatabase(config *config.PostgresConfig) (storage.Database, error) {
	db := &PostgresDB{
		Config: config,
	}
	error := db.getDBConnection()

	return db, error
}

func (d *PostgresDB) getDBConnection() error {
	var err error

	if d.DB != nil {
		return err
	}

	dataSource := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", d.Config.Host, d.Config.Port, d.Config.User,
		d.Config.Password, d.Config.Name)
	d.DB, err = sql.Open("postgres", dataSource)
	if err != nil {
		log.Println(err)
	}

	return err
}

func (p *PostgresDB) Insert(contract string) error {
	sql := ` INSERT into contracts (sc_addr) VALUES($1)`
	_, err := p.DB.Exec(sql, contract)
	return err
}

func (p *PostgresDB) Get(contract string) (string, error) {
	rows, err := p.DB.Query("SELECT sc_addr FROM contracts")
	if err != nil {
		log.Println(err)
		return "", err
	}
	log.Println(rows)
	return contract, err
}

func (p *PostgresDB) Update(contract string) error {
	return nil
}
