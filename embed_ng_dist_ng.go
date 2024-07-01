package phylotaxymusic

import "embed"

// NgDistNg is the export of angular distribution. This allows
// embedding of the pages in the web server
//
//go:embed ng-github.com-thomaspeugeot-phylotaxymusic/dist/ng-github.com-thomaspeugeot-phylotaxymusic
var NgDistNg embed.FS
