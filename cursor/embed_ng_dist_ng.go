package cursor

import "embed"

// NgDistNg is the export of angular distribution. This allows
// embedding of the pages in the web server
//
//go:embed ng-github.com-thomaspeugeot-phyllotaxymusic-cursor/dist/ng-github.com-thomaspeugeot-phyllotaxymusic-cursor
var NgDistNg embed.FS
