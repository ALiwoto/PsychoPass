package database

import (
	"fmt"
	"sync"
	"time"

	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylConfig"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/utils/logging"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var SESSION *gorm.DB

func StartDatabase() {
	// check if `SESSION` variable is already established or not.
	// if yes, check if we have got any error from it or not.
	// if there is an error in the session, it mean we have to establish
	// a new connection again.
	if SESSION != nil && SESSION.Error == nil {
		return
	}

	var db *gorm.DB
	var err error
	var conf *gorm.Config
	if sibylConfig.IsDebug() {
		conf = &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		}
	} else {
		conf = &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		}
	}
	if sibylConfig.SibylConfig.UseSqlite {
		db, err = gorm.Open(sqlite.Open(
			fmt.Sprintf("%s.db", sibylConfig.SibylConfig.DbName)), conf)
	} else {
		db, err = gorm.Open(postgres.Open(sibylConfig.SibylConfig.DbUrl), conf)
	}

	if err != nil {
		logging.Fatal("failed to connect to the database:", err)
	}

	SESSION = db
	logging.Info("Database connected")

	// Create tables if they don't exist
	err = SESSION.AutoMigrate(
		modelUser,
		modelToken,
		modelScan,
	)
	if err != nil {
		logging.Fatal(err)
	}

	if sibylConfig.SibylConfig.UseSqlite {
		dbMutex = &sync.Mutex{}
	}

	go cleanMaps()

	migrateOwners(sibylConfig.GetOwnersID())

	logging.Info("Auto-migrated database schema")
}

func cleanMaps() {
	mtime := sibylConfig.GetMaxCacheTime()
	tokenDbMap.SetExpiration(mtime)
	userDbMap.SetExpiration(mtime)
	scanDbMap.SetExpiration(mtime)
	associationScanMap.SetExpiration(mtime)
	for {
		time.Sleep(mtime)

		// please don't use len() function here, as it may return
		// `true` in some situations, but the maps may actually be
		// healthy, but they are only unused and so their caches are
		// completely deleted by cleaner.
		if tokenDbMap == nil || userDbMap == nil {
			return
		}

		tokenDbMap.DoCheck()
		userDbMap.DoCheck()
		scanDbMap.DoCheck()
		associationScanMap.DoCheck()
	}
}

func IsFirstTime() bool {
	return SESSION.Find(modelToken).RowsAffected == 0
}
