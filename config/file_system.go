// Copyright (C) Nguyen Nhat Tung
//
// Reverbzer is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package config

import "os"

func createDirectoryRecursive(dPath string) error {
	return os.MkdirAll(dPath, 0755)
}

func isExist(fPath string) bool {
	_, err := os.Stat(fPath)
	return !os.IsNotExist(err)
}
