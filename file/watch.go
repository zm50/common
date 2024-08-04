package file

import (
	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
)

type IWatcher interface {
	Monitor() error
	Watch(filePath string) error
	EventHandler(fn func(Op) error) error
	ErrorHandler(fn func(error)) error
	Stop() error
}

type Op uint

const (
	Create Op = 1 << iota
	Write
	Remove
	Rename
	Chmod
)

type watcher struct {
	*fsnotify.Watcher
	eventHandler func(Op) error
	errorHandler func(error)
}

func (w *watcher) Monitor() error {
	var resErr error
	for {
		var err error
		select {
		case event, ok := <-w.Events:
			if !ok {
				resErr = errors.New("event channel closed")
				goto EXIT
			}
			err = w.eventHandler(Op(event.Op))
		case e, ok := <-w.Errors:
			if !ok {
				resErr = errors.New("error channel closed")
				goto EXIT
			}
			err = e
		}

		w.errorHandler(err)
	}

	EXIT:

	return resErr
}

func (w *watcher) Watch(filePath string) error {
	return w.Watcher.Add(filePath)
}

func (w *watcher) EventHandler(fn func(Op) error) error {
	if fn != nil {
		w.eventHandler = fn
	}

	return nil
}

func (w *watcher) ErrorHandler(fn func(error)) error {
	if fn != nil {
		w.errorHandler = fn
	}

	return nil
}

func (w *watcher) Stop() error {
	return w.Watcher.Close()
}

func NewWatcher(bufferSize uint) (IWatcher, error) {
	bw, err := fsnotify.NewBufferedWatcher(bufferSize)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to create watcher")
	}

	w := &watcher{
		Watcher: bw,
		eventHandler: func(o Op) error {return nil},
		errorHandler: func(e error) {},
	}

	return w, err
}
