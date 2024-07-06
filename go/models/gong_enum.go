// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for StacksNames
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (stacksnames StacksNames) ToString() (res string) {

	// migration of former implementation of enum
	switch stacksnames {
	// insertion code per enum code
	case Phylotaxy:
		res = "phylotaxymusic"
	case GongsvgStackName:
		res = "gongsvg"
	case SidebarTree:
		res = "sidebar tree"
	case GongtreeStackName:
		res = "gongtree"
	case GongtableStackName:
		res = "gongtable"
	case GongsimStackName:
		res = "gongsim"
	}
	return
}

func (stacksnames *StacksNames) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "phylotaxymusic":
		*stacksnames = Phylotaxy
	case "gongsvg":
		*stacksnames = GongsvgStackName
	case "sidebar tree":
		*stacksnames = SidebarTree
	case "gongtree":
		*stacksnames = GongtreeStackName
	case "gongtable":
		*stacksnames = GongtableStackName
	case "gongsim":
		*stacksnames = GongsimStackName
	default:
		return errUnkownEnum
	}
	return
}

func (stacksnames *StacksNames) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Phylotaxy":
		*stacksnames = Phylotaxy
	case "GongsvgStackName":
		*stacksnames = GongsvgStackName
	case "SidebarTree":
		*stacksnames = SidebarTree
	case "GongtreeStackName":
		*stacksnames = GongtreeStackName
	case "GongtableStackName":
		*stacksnames = GongtableStackName
	case "GongsimStackName":
		*stacksnames = GongsimStackName
	default:
		return errUnkownEnum
	}
	return
}

func (stacksnames *StacksNames) ToCodeString() (res string) {

	switch *stacksnames {
	// insertion code per enum code
	case Phylotaxy:
		res = "Phylotaxy"
	case GongsvgStackName:
		res = "GongsvgStackName"
	case SidebarTree:
		res = "SidebarTree"
	case GongtreeStackName:
		res = "GongtreeStackName"
	case GongtableStackName:
		res = "GongtableStackName"
	case GongsimStackName:
		res = "GongsimStackName"
	}
	return
}

func (stacksnames StacksNames) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Phylotaxy")
	res = append(res, "GongsvgStackName")
	res = append(res, "SidebarTree")
	res = append(res, "GongtreeStackName")
	res = append(res, "GongtableStackName")
	res = append(res, "GongsimStackName")

	return
}

func (stacksnames StacksNames) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "phylotaxymusic")
	res = append(res, "gongsvg")
	res = append(res, "sidebar tree")
	res = append(res, "gongtree")
	res = append(res, "gongtable")
	res = append(res, "gongsim")

	return
}

// Utility function for TreeNames
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (treenames TreeNames) ToString() (res string) {

	// migration of former implementation of enum
	switch treenames {
	// insertion code per enum code
	case Sidebar:
		res = "sidebar"
	}
	return
}

func (treenames *TreeNames) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "sidebar":
		*treenames = Sidebar
	default:
		return errUnkownEnum
	}
	return
}

func (treenames *TreeNames) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Sidebar":
		*treenames = Sidebar
	default:
		return errUnkownEnum
	}
	return
}

func (treenames *TreeNames) ToCodeString() (res string) {

	switch *treenames {
	// insertion code per enum code
	case Sidebar:
		res = "Sidebar"
	}
	return
}

func (treenames TreeNames) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Sidebar")

	return
}

func (treenames TreeNames) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "sidebar")

	return
}

// end of insertion point for enum utility functions

type GongstructEnumStringField interface {
	string | StacksNames | TreeNames
	Codes() []string
	CodeValues() []string
}

type PointerToGongstructEnumStringField interface {
	*StacksNames | *TreeNames
	FromCodeString(input string) (err error)
}

type GongstructEnumIntField interface {
	int
	Codes() []string
	CodeValues() []int
}

type PointerToGongstructEnumIntField interface {
	
	FromCodeString(input string) (err error)
}

// Last line of the template
