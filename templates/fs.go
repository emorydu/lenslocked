package templates

import "embed"

//go:embed *.gohtml galleries/*.gohtml
var FS embed.FS
