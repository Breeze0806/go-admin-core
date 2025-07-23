package logger

import (
	"io"
	"os"

	"github.com/Breeze0806/go-admin-core/debug/writer"
	"github.com/Breeze0806/go-admin-core/logger"
	"github.com/Breeze0806/go-admin-core/sdk/pkg"

	log "github.com/Breeze0806/go-admin-core/logger"
)

// SetupLogger 日志 cap 单位为kb
func SetupLogger(opts ...Option) logger.Logger {
	op := setDefault()
	for _, o := range opts {
		o(&op)
	}
	if !pkg.PathExist(op.path) {
		err := pkg.PathCreate(op.path)
		if err != nil {
			log.Fatalf("create dir error: %s", err.Error())
		}
	}
	var err error
	var output io.Writer
	switch op.stdout {
	case "file":
		output, err = writer.NewFileWriter(
			writer.WithPath(op.path),
			writer.WithCap(op.cap<<10),
		)
		if err != nil {
			log.Fatal("logger setup error: %s", err.Error())
		}
	default:
		output = os.Stdout
	}
	var level logger.Level
	level, err = logger.GetLevel(op.level)
	if err != nil {
		log.Fatalf("get logger level error, %s", err.Error())
	}

	switch op.driver {
	default:
		log.DefaultLogger = logger.NewLogger(logger.WithLevel(level), logger.WithOutput(output))
	}
	return log.DefaultLogger
}
