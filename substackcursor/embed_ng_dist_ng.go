package substackcursor

import "embed"

// NgDistNg is the export of angular distribution. This allows
// embedding of the pages in the web server
//
//go:embed ng-github.com-thomaspeugeot-phylotaxymusic-substackcursor/dist/ng-github.com-thomaspeugeot-phylotaxymusic-substackcursor
var NgDistNg embed.FS
