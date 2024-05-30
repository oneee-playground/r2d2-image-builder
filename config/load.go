package config

import (
	"os"

	"github.com/google/go-containerregistry/pkg/authn"
)

func LoadFromEnv() {
	RegistryAddr = os.Getenv("REGISTRY_ADDR")
	RegistryUser = os.Getenv("REGISTRY_USER")
	Repository = os.Getenv("REGISTRY_REPO")

	dockerhubSecret := os.Getenv("DOCKERHUB_SECRET")

	RegistryAuth = authn.FromConfig(authn.AuthConfig{
		Username: RegistryUser,
		Password: dockerhubSecret,
	})

	MessageQueueURL = os.Getenv("MESSAGE_QUEUE_URL")
}
