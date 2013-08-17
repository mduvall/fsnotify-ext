package fsnotify2

import (
	"errors"
	"github.com/howeyc/fsnotify"
	"os"
	"path/filepath"
)

type Watcher struct {
	fsnotify.Watcher
}

func (w *Watcher) WatchAll(path string) error {
	watchFunc := func(path string, fi os.FileInfo, err error) error {
		return w.Watch(path)
	}

	if err := filepath.Walk(path, watchFunc); err != nil {
		return errors.New("Path given is not valid.")
	}

	return nil
}

// Return a new watcher of the extended type
func NewWatcher() (*Watcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	return &Watcher{*watcher}, nil
}
