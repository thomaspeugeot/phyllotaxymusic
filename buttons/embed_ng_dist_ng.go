package buttons

import "embed"

// NgDistNg is the export of angular distribution. This allows
// embedding of the pages in the web server
//
//go:embed ng-github.com-thomaspeugeot-phyllotaxymusic-buttons/dist/ng-github.com-thomaspeugeot-phyllotaxymusic-buttons
var NgDistNg embed.FS
