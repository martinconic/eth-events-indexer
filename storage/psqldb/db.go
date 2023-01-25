package psqldb

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/martinconic/eth-events-indexer/config"
	"github.com/martinconic/eth-events-indexer/data"
	"github.com/martinconic/eth-events-indexer/storage"

	"github.com/lib/pq"
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

func (p *PostgresDB) GetSmartContracts() ([]data.Contract, error) {
	sc := []data.Contract{}
	rows, err := p.DB.Query("SELECT id, sc_addr, is_indexing, last_tx_db, last_indx_date FROM contracts")
	if err != nil {
		return sc, err
	}

	for rows.Next() {
		var c data.Contract
		err = rows.Scan(
			&c.ID,
			&c.ScAddr,
			&c.IsIndexing,
			&c.LastTxDb,
			&c.LastIndxDate,
		)
		if err != nil {
			log.Println(err)
		}

		sc = append(sc, c)
	}

	return sc, nil
}

func (p *PostgresDB) UpdateIndexing(contract string, isIndexing bool) (string, error) {
	sql := `UPDATE contracts SET is_indexing = $1 where sc_addr = $2`
	_, err := p.DB.Exec(sql, isIndexing, contract)
	return getSqlResponse("Success update index", err)
}

func (p *PostgresDB) Insert(contract string) (int64, error) {
	sql := ` INSERT into contracts (sc_addr) VALUES($1)`
	result, err := p.DB.Exec(sql, contract)
	if err != nil {
		log.Println(err)
	}
	return result.LastInsertId()

}

func (p *PostgresDB) InsertEvent(tx *data.Transaction) (string, error) {
	sql := ` INSERT into transactions (sc_id, tx_addr, from_addr, to_addr, tokens, block_nr, tx_index,
		removed, log_index, log_name) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err := p.DB.Exec(sql, tx.ID, tx.TxAddr, tx.FromAddr, tx.ToAddr, tx.Tokens, tx.BlockNr, tx.TxIndex,
		tx.Removed, tx.LogIndex, tx.LogName)
	return getSqlResponse("Success inserting event", err)
}

func (p *PostgresDB) Get(contract string) (int, error) {
	var id int
	row := p.DB.QueryRow("SELECT id FROM contracts where sc_addr= ?", contract).Scan(&id)
	if row == nil {
		return -1, fmt.Errorf("error getting contract")
	}
	log.Println(id)
	return id, nil
}

func (p *PostgresDB) GetEvents(cid int) ([]data.Transaction, error) {
	var tx []data.Transaction
	rows, err := p.DB.Query("SELECT id, sc_id, tx_addr, from_addr, to_addr, tokens, block_nr, tx_index, removed, log_index, log_name FROM transactions where sc_id=$1", cid)
	if err != nil {
		return tx, err
	}

	for rows.Next() {
		var t data.Transaction
		err = rows.Scan(
			&t.ID,
			&t.ScId,
			&t.TxAddr,
			&t.FromAddr,
			&t.ToAddr,
			&t.Tokens,
			&t.BlockNr,
			&t.TxIndex,
			&t.Removed,
			&t.LogIndex,
			&t.LogName,
		)
		if err != nil {
			log.Println(err)
		}

		tx = append(tx, t)
	}

	return tx, nil
}

func (p *PostgresDB) Update(contract *data.Contract) (string, error) {
	return "Success", nil
}

func getSqlResponse(success string, err error) (string, error) {
	if err != nil {
		pqErr := err.(*pq.Error)
		return string(pqErr.Code), err
	}
	return success, err
}
