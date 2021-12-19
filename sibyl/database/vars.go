package database

import (
	"errors"
	"sync"

	sv "github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"
)

var (
	ErrInvalidToken   = errors.New("token is invalid")
	ErrNoSession      = errors.New("database session is not initialized")
	ErrTooManyRevokes = errors.New("token has been revoked too many times")
)

var (
	dbMutex              *sync.Mutex                     = &sync.Mutex{}
	tokenMapMutex        *sync.Mutex                     = &sync.Mutex{}
	userMapMutex         *sync.Mutex                     = &sync.Mutex{}
	scanMapMutex         *sync.Mutex                     = &sync.Mutex{}
	associationScanMutex *sync.Mutex                     = &sync.Mutex{}
	tokenDbMap           map[int64]*sv.Token             = make(map[int64]*sv.Token)
	userDbMap            map[int64]*sv.User              = make(map[int64]*sv.User)
	scanDbMap            map[string]*sv.Report           = make(map[string]*sv.Report)
	associationScanMap   map[string]*sv.MultiScanRawData = make(map[string]*sv.MultiScanRawData)
	modelUser            *sv.User                        = &sv.User{}
	modelToken           *sv.Token                       = &sv.Token{}
	modelScan            *sv.Report                      = &sv.Report{}
	lastStats            *sv.StatValue
)
