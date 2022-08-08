package clickhouse

import (
	"context"
	"fmt"
	"log"

	"github.com/kube-tarian/git-bridge/client/pkg/config"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"

	"github.com/ClickHouse/clickhouse-go/v2"
	// "gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

type DBClient struct {
	db   *gorm.DB
	conn driver.Conn
	conf *config.Config
}

func NewDBClient(conf *config.Config) (*DBClient, error) {
	// db, err := gorm.Open(clickhouse.Open(conf.DBAddress), &gorm.Config{})
	// if err != nil {
	// 	return nil, err
	// }
	// db.AutoMigrate(&models.Gitevent{})
	// // Set table options
	// db.Set("gorm:table_options", "ENGINE=File(cluster, default, hits)").AutoMigrate(&models.Gitevent{})

	// // Set table cluster options
	// db.Set("gorm:table_cluster_options", "on cluster default").AutoMigrate(&models.Gitevent{})

	log.Println("Initializing DB client")
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{conf.DBAddress},
		Auth: clickhouse.Auth{
			Database: "test",
		},
		// Auth: clickhouse.Auth{
		// 	Database: "default",
		// 	Username: "default",
		// 	Password: "",
		// },
		// Compression: &clickhouse.Compression{
		// 	Method: clickhouse.CompressionLZ4,
		// },
		// Settings: clickhouse.Settings{
		// 	"max_execution_time": 60,
		// },
		//Debug: true,
	})
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(context.Background()); err != nil {
		return nil, err
	}

	var settings []struct {
		Name        string  `ch:"name"`
		Value       string  `ch:"value"`
		Changed     uint8   `ch:"changed"`
		Description string  `ch:"description"`
		Min         *string `ch:"min"`
		Max         *string `ch:"max"`
		Readonly    uint8   `ch:"readonly"`
		Type        string  `ch:"type"`
	}
	if err = conn.Select(context.Background(), &settings, "SELECT * FROM system.settings WHERE name LIKE $1 ORDER BY length(name) LIMIT 5", "%max%"); err != nil {
		conn.Close()
		return nil, err
	}
	for _, s := range settings {
		fmt.Printf("name: %s, value: %s, type=%s\n", s.Name, s.Value, s.Type)
	}

	const dbCreate = `CREATE DATABASE IF NOT EXISTS test;`
	if err := conn.Exec(context.Background(), dbCreate); err != nil {
		return nil, err
	}

	const ddlSetExperimental = `SET allow_experimental_object_type=1;`
	if err := conn.Exec(context.Background(), ddlSetExperimental); err != nil {
		return nil, err
	}

	const ddl = `CREATE table IF NOT EXISTS git_json(event JSON) ENGINE = MergeTree ORDER BY tuple();`
	if err := conn.Exec(context.Background(), ddl); err != nil {
		return nil, err
	}

	return &DBClient{conn: conn}, nil
}

func (c *DBClient) InsertEvent(metrics string) {
	log.Printf("Inserting event: %v", metrics)
	insertStmt := fmt.Sprintf("INSERT INTO git_json FORMAT JSONAsObject %v", metrics)
	if err := c.conn.Exec(context.Background(), insertStmt); err != nil {
		log.Printf("Insert failed, %v", err)
	}
	// c.db.Create(metrics)
}
