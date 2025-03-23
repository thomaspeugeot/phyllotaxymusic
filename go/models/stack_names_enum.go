package models

// StacksNames - enumeration of possible 'type' values for an HTML <input> element
// swagger:enum StacksNames
type StacksNames string

// values for TableExtraNameEnum
const (
	// we have three application stacks. All have the same name
	Phylotaxy StacksNames = "phyllotaxymusic"

	GongsvgStackName StacksNames = "gongsvg"
	SidebarTree      StacksNames = "sidebar tree"
	ToneStackName    StacksNames = "gongtone"
	CursorStackName  StacksNames = "cursor"
	LoadStackName    StacksNames = "load"

	ButtonStackName        StacksNames = "button"
	RootSplitStackName     StacksNames = ""
	GongLibSliderStackName StacksNames = "slider"

	GongtreeStackName  StacksNames = "gongtree"
	GongtableStackName StacksNames = "gongtable"
	GongsimStackName   StacksNames = "gongsim"
)
