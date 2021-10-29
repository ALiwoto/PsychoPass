package database

import (
	"sync"

	sv "github.com/AnimeKaizoku/sibylapi-go/sibyl/core/sibylValues"
)

var (
	dbMutex       *sync.Mutex
	tokenMapMutex *sync.Mutex
	userMapMutex  *sync.Mutex
	tokenDbMap    map[int64]*sv.Token
	userDbMap     map[int64]*sv.User
)
