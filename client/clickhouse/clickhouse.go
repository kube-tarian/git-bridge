package clickhouse

import (
	"github.com/kube-tarian/git-bridge/models"

	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

func InsertEvent(url string, metrics models.Gitevent) {
	db, err := gorm.Open(clickhouse.Open(url), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Auto Migrate
	db.AutoMigrate(&models.Gitevent{})
	// Set table options
	db.Set("gorm:table_options", "ENGINE=File(cluster, default, hits)").AutoMigrate(&models.Gitevent{})

	// Set table cluster options
	db.Set("gorm:table_cluster_options", "on cluster default").AutoMigrate(&models.Gitevent{})

	// Insert
	db.Create(&models.Gitevent{Uuid: string(metrics.Uuid), Event: metrics.Event, Eventid: metrics.Eventid, Branch: metrics.Branch, Url: metrics.Url, Authorname: metrics.Authorname, Authormail: metrics.Authormail, DoneAt: metrics.DoneAt, Repository: metrics.Repository, Addedfiles: metrics.Addedfiles, Modifiedfiles: metrics.Modifiedfiles, Removedfiles: metrics.Removedfiles, Message: metrics.Message})
}
