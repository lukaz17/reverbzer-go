// Copyright (C) Nguyen Nhat Tung
//
// Reverbzer is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package engine

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Execute the program.
func Execute() {
	rootCmd := &cobra.Command{
		Use: "reverbzer",
		Long: fmt.Sprintf(
			`Reverbzer Web Server Application v%s.
Copyright (C) %d Lukaz.
Licensed under MIT license. See LICENSE file along with this program for more details.`,
			"0.1.0",
			2025),
		Short:   "Light weight web server application for debugging incoming requests.",
		Version: "0.1.0",
	}

	rootCmd.AddCommand(ServerCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
