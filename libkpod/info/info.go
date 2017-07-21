package info

import (
	"errors"
	"sync"

	"golang.org/x/net/context"
)

// InfoGiverFunc is a function that provides information back on demand. Like
// for display in `kpod info`.
//
// The name of return info is the scope of that collection.
//
// Errors inside an InfoGiverFunc are nested in the returned info object. A
// top-leverl error should be treated as a failure.
type InfoGiverFunc func(ctx context.Context) (name string, info map[string]interface{}, err error)

// RegisterInfoGiver adds an information function to the avaiable set
func RegisterInfoGiver(igf InfoGiverFunc) {
	igfMutex.Lock()
	defer igfMutex.Unlock()
	infoGivers = append(infoGivers, igf)
}

var (
	igfMutex   sync.Mutex
	infoGivers = []InfoGiverFunc{}
)

// InfoGivers is the list of registered InfoGiverFunc for usage/display
func InfoGivers() []InfoGiverFunc {
	return infoGivers[:]
}

// for testing
func resetInfoGivers() {
	infoGivers = []InfoGiverFunc{}
}

// CollectInfo processes the provided set of infogivers based on context ctx.
func CollectInfo(ctx context.Context, igfs []InfoGiverFunc) map[string]interface{} {
	info := map[string]interface{}{}
	for _, igf := range igfs {
		thisName, thisInfo, err := igf(ctx)
		if err == ErrNoInfo {
			continue
		}
		if err != nil {
			info[thisName] = Err(err)
			continue
		}
		info[thisName] = thisInfo
	}
	return info
}

// ErrNoInfo is an indicator from InfoGiverFunc that can be passed if there is
// no information avaiable.
// This should not be treated as a fatal error.
var ErrNoInfo = errors.New("no information to provide")

func Err(err error) map[string]interface{} {
	return map[string]interface{}{
		"error": err.Error(),
	}
}
