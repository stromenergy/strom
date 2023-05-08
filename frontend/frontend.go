package frontend

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"strings"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/stromenergy/strom/internal/util"
)

func MountRoutes(engine *gin.Engine) {
	staticFileSystem := NewStaticFileSystem()
	fallbackFileSystem := NewFallbackFileSystem(staticFileSystem)

	engine.Use(static.Serve("/", staticFileSystem))
	engine.Use(static.Serve("/", fallbackFileSystem))
}

//go:embed build/*
var frontendFS embed.FS

type StaticFileSystem struct {
	http.FileSystem
}

var _ static.ServeFileSystem = (*StaticFileSystem)(nil)

func NewStaticFileSystem() *StaticFileSystem {
	sub, err := fs.Sub(frontendFS, "build")
	util.OnErrorPanic(err)

	return &StaticFileSystem{
		FileSystem: http.FS(sub),
	}
}

func (s *StaticFileSystem) Exists(prefix string, path string) bool {
	buildpath := fmt.Sprintf("build%s", path)

	// support for folders
	if strings.HasSuffix(path, "/") {
		_, err := frontendFS.ReadDir(strings.TrimSuffix(buildpath, "/"))
		return err == nil
	}

	// support for files
	file, err := frontendFS.Open(buildpath)

	if file != nil {
		_ = file.Close()
	}

	return err == nil
}

type FallbackFileSystem struct {
	staticFileSystem *StaticFileSystem
}

var _ static.ServeFileSystem = (*FallbackFileSystem)(nil)
var _ http.FileSystem = (*FallbackFileSystem)(nil)

func NewFallbackFileSystem(staticFileSystem *StaticFileSystem) *FallbackFileSystem {
	return &FallbackFileSystem{
		staticFileSystem: staticFileSystem,
	}
}

func (f *FallbackFileSystem) Open(path string) (http.File, error) {
	return f.staticFileSystem.Open("/index.html")
}

func (f *FallbackFileSystem) Exists(prefix string, path string) bool {
	return true
}
