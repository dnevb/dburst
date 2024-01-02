package ui

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed all:build/*
var buildDist embed.FS
var Root = http.FS(buildDist)

var config = filesystem.Config{
	Root:         Root,
	PathPrefix:   "build",
	Index:        "index.html",
	NotFoundFile: "build/index.html",
}

var UiFS = filesystem.New(config)
