package web

import (
	"os"
	"path/filepath"
	sys "system"

	log "github.com/tengfei-xy/go-log"

	"github.com/gin-gonic/gin"
)

func FontRoute(g *gin.Engine) {
	g.GET("/font/:filename", GetFont)
	g.POST("/font/:filename", PostFont)
	g.DELETE("/font/:filename", DeleteFont)
	// g.PUT("/font/:name", PutFont)
}
func FontsRoute(g *gin.Engine) {
	g.GET("/fonts", GetFonts)
	// g.PUT("/font/:name", PutFont)
}

func GetFont(c *gin.Context) {
	filename := c.Param("filename")
	fontpath := c.MustGet("fontpath").(string)

	f, err := os.ReadFile(filepath.Join(fontpath, filename))
	if err != nil {
		notFound(c)
		return
	}
	font, _ := sys.GetInfo(f)
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
