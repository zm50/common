package config

import (
	"io"
	"log"
	"os"

	"github.com/pkg/errors"

	"github.com/zm50/common/constant"
	"github.com/zm50/common/file"
	"github.com/zm50/common/serialize"
)

type FileConfig[T serialize.Serializer, D any] struct {
	Serializer T
	FilePath string
	Config *D
	defaultFileChangeHandler func(o file.Op) error
	defaultErrorHandler func(err error)
}

func NewFileConfig[T serialize.Serializer, D any](filePath string, config *D) *FileConfig[T, D] {
	f := &FileConfig[T, D]{
		Serializer: *new(T),
		FilePath: filePath,
		Config: config,
	}

	f.defaultErrorHandler = defaultErrorHandler[T, D](f)
	f.defaultFileChangeHandler = defaultFileChangeHandler[T, D](f)

	return f
}

func (f *FileConfig[T, D])Notify() error {
	watcher, err := file.NewWatcher(constant.Zero)
	if err != nil {
		return errors.WithMessage(err, "create watcher failed")
	}

	err = watcher.Watch(f.FilePath)
	if err != nil {
		watcher.Stop()
		return errors.WithMessage(err, "watch file failed")
	}

	watcher.EventHandler(func(o file.Op) error {
		file, err := os.Open(f.FilePath)
		if err != nil {
			return errors.WithMessage(err, "open file failed")
		}
		defer file.Close()

		data, err := io.ReadAll(file)
		if err != nil {
			return errors.WithMessage(err, "read file failed")
		}

		err = f.Serializer.Unmarshal(data, f.Config)
		return err
	})

	watcher.ErrorHandler(func(err error) {
		log.Println("watcher error:", err)
	})

	go func() {
		defer watcher.Stop()

		err := watcher.Monitor()
		if err != nil {
			log.Fatalln("monitor failed:", err)
		}
	}()

	return nil
}

func defaultFileChangeHandler[T serialize.Serializer, D any](f *FileConfig[T, D]) func(o file.Op) error {
	return func(o file.Op) error {
		file, err := os.Open(f.FilePath)
		if err != nil {
			return errors.WithMessage(err, "open file failed")
		}
		defer file.Close()
	
		data, err := io.ReadAll(file)
		if err != nil {
			return errors.WithMessage(err, "read file failed")
		}
	
		err = f.Serializer.Unmarshal(data, f.Config)
		return err
	}
}

func defaultErrorHandler[T serialize.Serializer, D any](f *FileConfig[T, D]) func(err error) {
	return func(err error) {
		log.Printf("watcher file %s change error: %v\n", f.FilePath, err)
	}
}
