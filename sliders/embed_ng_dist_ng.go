package sliders

import "embed"

// NgDistNg is the export of angular distribution. This allows
// embedding of the pages in the web server
//
//go:embed ng-github.com-thomaspeugeot-phyllotaxymusic-sliders/dist/ng-github.com-thomaspeugeot-phyllotaxymusic-sliders
var NgDistNg embed.FS
