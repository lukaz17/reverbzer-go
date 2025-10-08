// Copyright (C) Nguyen Nhat Tung
//
// Reverbzer is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package engine

import (
	"github.com/lukaz17/reverbzer-go/server"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

// ServerModule handles user requests related checksum file creation and verification.
type ServerModule struct {
	logger zerolog.Logger
}

// Return new ServerModule.
func NewServerModule(c *Controller, cmdName string) *ServerModule {
	return &ServerModule{
		logger: c.CommandLogger("checksum", cmdName),
	}
}

func (m *ServerModule) Main(port uint16) error {
	server.Run(port, m.logger)
	return nil
}

// Decorator to log error occurred when calling handlers.
func (m *ServerModule) logError(err error) {
	if err != nil {
		m.logger.Err(err).Msg("Unexpected error has occurred. Program will exit.")
	}
}

// Define Cobra Command for Checksum module.
func ServerCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "server",
		Short: "Start web server application.",
		Run: func(cmd *cobra.Command, args []string) {
			c := NewController(true)
			defer c.Close()
			flags := ParseChecksumFlags(cmd)
			m := NewServerModule(c, "main")
			m.logError(m.Main(flags.Port))
		},
	}
	rootCmd.Flags().Uint16P("port", "p", 11111, "Port for the web server to listen.")

	return rootCmd
}

// Struct ChecksumFlags contains all flags used by Checksum module.
type ChecksumFlags struct {
	Port uint16
}

// Extract all flags from a Cobra Command.
func ParseChecksumFlags(cmd *cobra.Command) *ChecksumFlags {
	port, _ := cmd.Flags().GetUint16("port")

	return &ChecksumFlags{
		Port: port,
	}
}
