// Copyright (C) Nguyen Nhat Tung
//
// Reverbzer is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package config

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/rs/zerolog"
	"github.com/tforce-io/tf-golib/opx"
)

// Entrypoint for creating a ZeroLog logger instance.
func InitZerolog(configDir string, useFS bool) (zerolog.Logger, *os.File, error) {
	consoleWriter := &zerolog.FilteredLevelWriter{
		Writer: zerolog.LevelWriterAdapter{
			Writer: zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false, TimeFormat: time.DateTime},
		},
		Level: zerolog.TraceLevel,
	}

	logFile, err := InitLogFile(useFS, configDir)
	if logFile == nil {
		consoleLogger := zerolog.New(consoleWriter).With().Timestamp().Logger()
		return consoleLogger, nil, err
	}

	fileWriter := &zerolog.FilteredLevelWriter{
		Writer: zerolog.LevelWriterAdapter{
			Writer: logFile,
		},
		Level: zerolog.TraceLevel,
	}
	multiWriter := zerolog.MultiLevelWriter(consoleWriter, fileWriter)
	logger := zerolog.New(multiWriter).With().Timestamp().Logger()
	return logger, logFile, nil
}

// Create and return log file handle only if useFS is true.
func InitLogFile(useFS bool, workdingDir string) (*os.File, error) {
	if !useFS {
		return nil, nil
	}
	logDir := path.Join(opx.Ternary(workdingDir == "", ".", workdingDir), "logs")
	if !isExist(logDir) {
		err := createDirectoryRecursive(logDir)
		if err != nil {
			return nil, err
		}
	}
	logFileName := fmt.Sprintf("unifiler-%s.log", time.Now().UTC().Format("20060102"))
	logFilePath := filepath.Join(logDir, logFileName)
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}
