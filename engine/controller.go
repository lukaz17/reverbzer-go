// Copyright (C) Nguyen Nhat Tung
//
// Reverbzer is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package engine

import (
	"os"

	"github.com/lukaz17/reverbzer-go/config"
	"github.com/rs/zerolog"
)

// Controller is the entrypoint for working with application configurations and
// loggings.
type Controller struct {
	Logger  zerolog.Logger
	logFile *os.File
}

// Entrypoint for creating new instance of Controller.
// useFS will instruct this function to read configurations and create log file.
func NewController(useFS bool) *Controller {
	useFS = false // Temporarily disable file system for logging and configuration
	logger, logFile, err2 := config.InitZerolog("", useFS)
	if err2 != nil {
		logger.Err(err2).Msg("error initializing log file")
	}
	return &Controller{
		Logger: logger,

		logFile: logFile,
	}
}

// Execute additional clean up when terminate the app.
func (c *Controller) Close() {
	if c.logFile != nil {
		c.logFile.Close()
		c.logFile = nil
	}
}

// Get a ZeroLog logger instance for command handler from root instance.
func (c *Controller) CommandLogger(module, command string) zerolog.Logger {
	return c.Logger.With().Str("module", module).Str("command", command).Logger()
}

// Get a ZeroLog logger instance for module from root instance.
func (c *Controller) ModuleLogger(module string) zerolog.Logger {
	return c.Logger.With().Str("module", module).Logger()
}
