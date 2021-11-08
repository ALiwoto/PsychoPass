package database

import (
	"errors"
	"sync"

	sv "github.com/AnimeKaizoku/PsychoPass/sibyl/core/sibylValues"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrNoSession    = errors.New("database session is not initialized")
)

var (
	dbMutex       *sync.Mutex
	tokenMapMutex *sync.Mutex
	userMapMutex  *sync.Mutex
	tokenDbMap    map[int64]*sv.Token
	userDbMap     map[int64]*sv.User
	modelUser     *sv.User  = &sv.User{}
	modelToken    *sv.Token = &sv.Token{}
	lastStats     *sv.StatValue
)
