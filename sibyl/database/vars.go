/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
package database

import (
	"errors"
	"sync"

	ws "github.com/AnimeKaizoku/ssg/ssg"
	sv "github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"
)

var (
	ErrInvalidToken   = errors.New("token is invalid")
	ErrNoSession      = errors.New("database session is not initialized")
	ErrTooManyRevokes = errors.New("token has been revoked too many times")
)

var (
	dbMutex *sync.Mutex = &sync.Mutex{}
	/*
		tokenMapMutex        *sync.Mutex                     = &sync.Mutex{}
		userMapMutex         *sync.Mutex                     = &sync.Mutex{}
		scanMapMutex         *sync.Mutex                     = &sync.Mutex{}
		associationScanMutex *sync.Mutex                     = &sync.Mutex{}
		tokenDbMap           map[int64]*sv.Token             = make(map[int64]*sv.Token)
		userDbMap            map[int64]*sv.User              = make(map[int64]*sv.User)
		scanDbMap            map[string]*sv.Report           = make(map[string]*sv.Report)
		associationScanMap   map[string]*sv.MultiScanRawData = make(map[string]*sv.MultiScanRawData)
	*/
	tokenDbMap                                = ws.NewSafeEMap[int64, sv.Token]()
	userDbMap                                 = ws.NewSafeEMap[int64, sv.User]()
	scanDbMap                                 = ws.NewSafeEMap[string, sv.Report]()
	associationScanMap                        = ws.NewSafeEMap[string, sv.MultiScanRawData]()
	modelUser            *sv.User             = &sv.User{}
	emptyUser            *sv.User             = &sv.User{}
	emptyAssociationData *sv.MultiScanRawData = &sv.MultiScanRawData{}
	modelToken           *sv.Token            = &sv.Token{}
	modelScan            *sv.Report           = &sv.Report{}
	lastStats            *sv.StatValue
	statsMutex           *sync.Mutex = &sync.Mutex{}
)
