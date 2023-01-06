package view

import (
	"net/http"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func LoadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	templatePath := templatesDir + "/templates/*.html"

	templates, err := filepath.Glob(templatePath)
	if err != nil {
		panic(err.Error())
	}

	partials, err := filepath.Glob(templatesDir + "/partials/*.html")

	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, include := range partials {
		layoutCopy := make([]string, len(templates))
		copy(layoutCopy, templates)
		files := append(layoutCopy, include)

		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", nil)
}
