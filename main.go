package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"path/filepath"
	"strings"
)

var mainTpl = `
<html>
	<body>
		<h1>Список файлов</h1>
		<ul>
			%s
		</ul>
	</body>
</html>
`
var mainListItemTpl = `
<li>
	<a href="view/%v">%v</a>
</li>
`

var viewHandlerTpl = `
<html>
	<body>
		<h1>%s</h1>
        <p><a href="/">Назад</a></p>
		<img src='/images/%s' height="400" width="400" />
        <ul>
			%s
		</ul>
	</body>
</html>
`

func main() {
	e := echo.New()

	e.Static("/images", "images")

	e.GET("/", mainHandler)
	e.GET("/view/:name", viewHandler)

	e.Logger.Fatal(e.Start(":8000"))
}

func getImageListItems() string {
	files, err := filepath.Glob("images/*.jpg")
	if err != nil {
		return err.Error()
	}

	var items string
	for _, file := range files {
		currentFile := strings.TrimPrefix(file, "images/")
		items += fmt.Sprintf(mainListItemTpl, currentFile, currentFile)
	}

	return items
}

func mainHandler(c echo.Context) error {
	return c.HTML(200, fmt.Sprintf(mainTpl, getImageListItems()))
}

func viewHandler(c echo.Context) error {
	name := c.Param("name")
	return c.HTML(200, fmt.Sprintf(viewHandlerTpl, name, name, getImageListItems()))
}
