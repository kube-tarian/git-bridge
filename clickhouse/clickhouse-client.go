package clickhouse

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ClickHouse/clickhouse-go"
	"github.com/vijeyash1/gitevent/models"
)

func GetClickHouseConnection(url string) (*sql.DB, error) {
	connect, err := sql.Open("clickhouse", url)
	if err != nil {
		log.Fatal(err)
	}
	if err := connect.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			fmt.Println(err)
		}
		return nil, err
	}
	return connect, nil
}

func CreateGitSchema(connect *sql.DB) {
	_, err := connect.Exec(`
		CREATE TABLE IF NOT EXISTS vij (
			id                UUID,
			event        			String,
			eventid        		String,
			branch        		String,
			url        				String,
			authorname 				String,
			authormail     		String,
			doneat						String,
			repository    		String,
			addedfiles    		String,
			modifiedfiles 		String,
			removedfiles  		String,
			message       		String
		) engine=File(TabSeparated)
	`)
	if err != nil {
		log.Fatal(err)
	}
}
func InsertGitEvent(connect *sql.DB, metrics models.Gitevent) {
	var (
		tx, _   = connect.Begin()
		stmt, _ = tx.Prepare("INSERT INTO vij (id, event, eventid, branch, url, authorname, authormail, doneat, repository, addedfiles, modifiedfiles, removedfiles, message) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,)")
	)

	defer stmt.Close()
	if _, err := stmt.Exec(
		metrics.Uuid,
		metrics.Event,
		metrics.Eventid,
		metrics.Branch,
		metrics.Url,
		metrics.Authorname,
		metrics.Authormail,
		metrics.DoneAt,
		metrics.Repository,
		metrics.Addedfiles,
		metrics.Modifiedfiles,
		metrics.Removedfiles,
		metrics.Message,
	); err != nil {
		log.Fatal(err)
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}
