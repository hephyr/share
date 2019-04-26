package templates

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "index.html",
		FileModTime: time.Unix(1556252719, 0),

		Content: string("{{define \"index\"}}\n<html>\n<head>\n    <title>File Share</title>\n    {{template \"header\"}}\n</head>\n    <body>\n        <div class=\"ui container\">\n            <div class=\"ui list\">\n                <h1>File Share</h1>\n                {{range .}}\n                <div class=\"item\">\n                    {{if .IsDir}}\n                        <i class=\"folder icon\"></i>\n                        <div class=\"content\">\n                            <div class=\"header\"><a href=\"{{.Name}}/\">{{.Name}}/</a></div>\n                        </div>\n                    {{else}}\n                        <i class=\"file icon\"></i>\n                        <div class=\"content\">\n                            <div class=\"header\"><a download href=\"{{.Name}}\">{{.Name}}</a></div>\n                        </div>\n                    {{end}}\n                </div>\n                {{end}}\n            </div>\n        </div>\n    </body>\n    {{template \"footer\"}}\n</html>\n\n{{end}}"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1556252719, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "index.html"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`layouts`, &embedded.EmbeddedBox{
		Name: `layouts`,
		Time: time.Unix(1556252719, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"index.html": file2,
		},
	})
}

func init() {

	// define files
	file4 := &embedded.EmbeddedFile{
		Filename:    "footer.html",
		FileModTime: time.Unix(1556252719, 0),

		Content: string("{{define \"footer\"}}\n<!-- <script src=\"/static/js/semantic.min.js\"></script> -->\n{{end}}"),
	}
	file5 := &embedded.EmbeddedFile{
		Filename:    "header.html",
		FileModTime: time.Unix(1556252719, 0),

		Content: string("{{define \"header\"}}\n<meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\">\n<link rel=\"stylesheet\" href=\"/static/css/semantic.min.css\">\n<link rel=\"stylesheet\" href=\"/static/css/fontawesome.css\">\n{{end}}"),
	}

	// define dirs
	dir3 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1556252719, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file4, // "footer.html"
			file5, // "header.html"

		},
	}

	// link ChildDirs
	dir3.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`widgets`, &embedded.EmbeddedBox{
		Name: `widgets`,
		Time: time.Unix(1556252719, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir3,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"footer.html": file4,
			"header.html": file5,
		},
	})
}
