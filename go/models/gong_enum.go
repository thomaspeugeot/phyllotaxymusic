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
	res = append(res, "gongtree")
	res = append(res, "gongtable")
	res = append(res, "gongsim")

	return
}

// end of insertion point for enum utility functions

type GongstructEnumStringField interface {
	string | StacksNames
	Codes() []string
	CodeValues() []string
}

type PointerToGongstructEnumStringField interface {
	*StacksNames
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
