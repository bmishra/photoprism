package mutex

import (
	"sync"
)

var (
	Db          = sync.Mutex{}
	MainWorker  = Busy{}
	SyncWorker  = Busy{}
	ShareWorker = Busy{}
	PrismWorker = Busy{}
)

// WorkersBusy returns true if any worker is busy.
func WorkersBusy() bool {
	return MainWorker.Busy() || SyncWorker.Busy() || ShareWorker.Busy() || PrismWorker.Busy()
}
