package embedstatic

import (
	"embed"
	"io/fs"
)

//go:embed uiwebdist
var uiwebEmbedFS embed.FS

func GetFS(name string) (fs.FS, error) {
	switch name {
	case "uiweb":
		return fs.Sub(uiwebEmbedFS, "uiwebdist")
	default:
		panic("not found directories/files to embed")
	}
}
