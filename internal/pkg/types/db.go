package types

import "time"

// Database config
type Database struct {
	Host    string        `env:"HOST_DB" default:"127.0.0.1:27017"`
	Name    string        `env:"DATABASE_DB" default:"passtrack"`
	User    string        `env:"USERNAME_DB"`
	Pass    string        `env:"PASSWORD_DB"`
	Timeout time.Duration `env:"TIMEOUT_DB" default:"10s"`
}
