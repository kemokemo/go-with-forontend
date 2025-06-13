package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

//go:embed frontendContents/*
var embedContents embed.FS

func main() {
	os.Exit(run())
}

func run() int {
	router := gin.Default()

	// 第二引数の相対パスは、embedで指定したパスに一致させる
	staticContents, err := fs.Sub(embedContents, "frontendContents")
	if err != nil {
		log.Println("failed to open embedded contents, ", err)
		return 1
	}

	// 静的ファイルとして配信する
	router.StaticFS("/", http.FS(staticContents))

	err = router.Run(":9090")
	if err != nil {
		log.Println("failed to run server, ", err)
		return 1
	}

	return 0
}
