package config

import "github.com/google/go-containerregistry/pkg/authn"

var (
	RegistryAddr string
	RegistryUser string
	Repository   string

	RegistryAuth authn.Authenticator
)
