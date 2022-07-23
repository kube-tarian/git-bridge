package clickhouse

import (
	"github.com/kube-tarian/git-bridge/models"

	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

type DBClient struct {
	db *gorm.DB
}

func Initialize(url string) (*DBClient, error) {
	db, err := gorm.Open(clickhouse.Open(url), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.Gitevent{})
	// Set table options
	db.Set("gorm:table_options", "ENGINE=File(cluster, default, hits)").AutoMigrate(&models.Gitevent{})

	// Set table cluster options
	db.Set("gorm:table_cluster_options", "on cluster default").AutoMigrate(&models.Gitevent{})

	return &DBClient{db: db}, nil

}

func (c *DBClient) InsertEvent(metrics *models.Gitevent) {

	c.db.Create(metrics)

}
