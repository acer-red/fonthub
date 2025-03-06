package web

import (
	"os"
	"path/filepath"
	sys "system"

	log "github.com/tengfei-xy/go-log"

	"github.com/gin-gonic/gin"
)

func FontRoute(g *gin.Engine) {
	g.GET("/api/font/:filename", GetFont)
	g.POST("/api/font/:filename", PostFont)
	g.DELETE("/api/font/:filename", DeleteFont)
}
func FontsRoute(g *gin.Engine) {
	g.GET("/api/fonts", GetFonts)
}

func GetFont(c *gin.Context) {
	filename := c.Param("filename")
	fontpath := c.MustGet("fontpath").(string)

	if filename == "" {
		badRequest(c)
		return
	}
	f, err := os.ReadFile(filepath.Join(fontpath, filename))
	if err != nil {
		notFound(c)
		return
	}
	font, err := sys.GetInfo(f)
	if err != nil {
		internalServerError(c)
		return
	}
	log.Infof("%s", font.FullName)
	c.Data(200, "application/x-font-ttf", f)
}
func GetFonts(c *gin.Context) {
	type req struct {
		Files []string `json:"files"`
	}
	fontpath := c.MustGet("fontpath").(string)
	files, err := os.ReadDir(fontpath)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	var fonts []string
	for _, f := range files {
		fonts = append(fonts, f.Name())
	}
	okData(c, req{
		Files: fonts,
	})
}
func PostFont(c *gin.Context) {

}
func DeleteFont(c *gin.Context) {

}
