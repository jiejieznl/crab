package log

import (
	"time"

	"github.com/huashusu/rotate"
)

func createWriter(cfg *Config, layout string) (*rotate.Rotate, error) {
	file, err := rotate.New(cfg.Directory, layout, "log",
		rotate.WithMaxAge(time.Duration(cfg.MaxAge)*rotate.Day),
		rotate.WithMaxSize(rotate.MB*int64(cfg.MaxSize)),
	)
	if err != nil {
		return nil, err
	}
	return file, err
}
