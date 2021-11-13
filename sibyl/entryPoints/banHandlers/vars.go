package banHandlers

import "sync"

var (
	multiBanMutex   *sync.Mutex
	multiUnBanMutex *sync.Mutex
)
