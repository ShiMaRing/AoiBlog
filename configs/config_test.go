package configs

import (
	"github.com/fsnotify/fsnotify"
	"testing"
)

func TestFileWatcher(t *testing.T) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		t.Fatal(err)
	}
	defer watcher.Close()
	done := make(chan bool)
	go func() {
		for true {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				t.Logf("event %v", event)
				t.Logf("op %v", event.Op.String())
				if event.Op&fsnotify.Write == fsnotify.Write {
					t.Logf("modifed file %s", event.String())
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				t.Logf("err %v", err)
			}
		}
	}()
	path := "test.txt"
	_ = watcher.Add(path)
	<-done
}
