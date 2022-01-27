package utils

import "sync"

var (
	sendMultipleMessageMutex = &sync.Mutex{}
)
