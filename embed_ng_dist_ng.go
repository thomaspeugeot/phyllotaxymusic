package phyllotaxymusic

import "embed"

// NgDistNg is the export of angular distribution. This allows
// embedding of the pages in the web server
//
//go:embed ng-github.com-thomaspeugeot-phyllotaxymusic/dist/ng-github.com-thomaspeugeot-phyllotaxymusic
var NgDistNg embed.FS
