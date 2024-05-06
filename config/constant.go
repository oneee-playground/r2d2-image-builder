package config

import "os"

var (
	Tmpfs = os.TempDir()
)
