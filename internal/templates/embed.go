package templates

import (
    "embed"
)

//go:embed template/*.hbs
var TemplateFiles embed.FS