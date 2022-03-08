package config

import "embed"

//go:embed crd/bases/*
var manifests embed.FS

func Manifests() embed.FS {
	return manifests
}
