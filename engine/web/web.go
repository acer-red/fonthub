package web

import (
	"fmt"
	"os"
	"path/filepath"
	sys "system"

	log "github.com/tengfei-xy/go-log"

	"github.com/gin-gonic/gin"
)

func FontRoute(g *gin.Engine) {
	g.POST("/api/font/:filename", PostFont)
	g.DELETE("/api/font/:filename", DeleteFont)
}
func FontsRoute(g *gin.Engine) {
	g.GET("/api/fonts", GetFonts)
}

func GetFonts(c *gin.Context) {
	log.Infof("GetFonts")
	// locale=zh-CN
	type FileInfoTotal struct {
		sys.Font
		FileName    string `json:"file_name"`
		DownloadURL string `json:"download_url"`
	}
	type req struct {
		Items []FileInfoTotal `json:"items"`
	}
	fontpath := c.MustGet("fontpath").(string)
	serverAddress := c.MustGet("server_address").(string)
	fileDir, err := os.ReadDir(fontpath)
	if err != nil {
		notFound(c)
		return
	}

	var files []FileInfoTotal
	for _, f := range fileDir {

		content, err := os.ReadFile(filepath.Join(fontpath, f.Name()))
		if err != nil {
			internalServerError(c)
			return
		}

		font, err := sys.GetInfo(content)
		if err != nil {
			internalServerError(c)
			return
		}

		files = append(files, FileInfoTotal{
			FileName:    f.Name(),
			DownloadURL: fmt.Sprintf("%s/fonts/%s", serverAddress, f.Name()),
			Font:        font,
		})
	}

	okData(c, req{
		Items: files,
	})
}
func PostFont(c *gin.Context) {

}
func DeleteFont(c *gin.Context) {

}
