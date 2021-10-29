package database

import "github.com/AnimeKaizoku/sibylapi-go/sibyl/core/sibylConfig"

func lockdb() {
	if sibylConfig.SibylConfig.UseSqlite {
		dbMutex.Lock()
	}
}
func unlockdb() {
	if sibylConfig.SibylConfig.UseSqlite {
		dbMutex.Unlock()
	}
}
