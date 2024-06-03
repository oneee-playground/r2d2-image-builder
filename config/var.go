package config

import "github.com/google/go-containerregistry/pkg/authn"

var (
	RegistryAddr string
	RegistryUser string

	RegistryAuth authn.Authenticator
)

var (
	MessageQueueURL string
)
