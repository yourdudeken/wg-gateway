package templates

import (
	"embed"
)

//go:embed files/*
var Templates embed.FS
