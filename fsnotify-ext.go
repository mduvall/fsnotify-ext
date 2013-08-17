package fsnotify2

import (
	"github.com/howeyc/fsnotify"
)

type Watcher struct {
	fsnotify.Watcher
}

// Return a new watcher of the extended type
func NewWatcher() (*Watcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	return &Watcher{*watcher}, nil
}
