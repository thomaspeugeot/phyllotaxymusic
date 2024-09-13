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
	case Gongtone:
		res = "gongtone"
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
	case "gongtone":
		*stacksnames = Gongtone
		return
	case "gongsvg":
		*stacksnames = GongsvgStackName
		return
	case "sidebar tree":
		*stacksnames = SidebarTree
		return
	case "gongtree":
		*stacksnames = GongtreeStackName
		return
	case "gongtable":
		*stacksnames = GongtableStackName
		return
	case "gongsim":
		*stacksnames = GongsimStackName
		return
	default:
		return errUnkownEnum
	}
}

func (stacksnames *StacksNames) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Gongtone":
		*stacksnames = Gongtone
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
	case Gongtone:
		res = "Gongtone"
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
	res = append(res, "Gongtone")
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
	res = append(res, "gongtone")
	res = append(res, "gongsvg")
	res = append(res, "sidebar tree")
	res = append(res, "gongtree")
	res = append(res, "gongtable")
	res = append(res, "gongsim")

	return
}

// end of insertion point for enum utility functions

type GongstructEnumStringField interface {
	Codes() []string
	CodeValues() []string
	ToString() string
}

type PointerToGongstructEnumStringField interface {
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
