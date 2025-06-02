// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for GenerationMode
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (generationmode GenerationMode) ToString() (res string) {

	// migration of former implementation of enum
	switch generationmode {
	// insertion code per enum code
	case DOWNLOAD:
		res = "DOWNLOAD"
	case STATIC_WEB_SITE:
		res = "STATIC_WEB_SITE"
	}
	return
}

func (generationmode *GenerationMode) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "DOWNLOAD":
		*generationmode = DOWNLOAD
		return
	case "STATIC_WEB_SITE":
		*generationmode = STATIC_WEB_SITE
		return
	default:
		return errUnkownEnum
	}
}

func (generationmode *GenerationMode) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "DOWNLOAD":
		*generationmode = DOWNLOAD
	case "STATIC_WEB_SITE":
		*generationmode = STATIC_WEB_SITE
	default:
		err = errUnkownEnum
	}
	return
}

func (generationmode *GenerationMode) ToCodeString() (res string) {

	switch *generationmode {
	// insertion code per enum code
	case DOWNLOAD:
		res = "DOWNLOAD"
	case STATIC_WEB_SITE:
		res = "STATIC_WEB_SITE"
	}
	return
}

func (generationmode GenerationMode) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "DOWNLOAD")
	res = append(res, "STATIC_WEB_SITE")

	return
}

func (generationmode GenerationMode) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "DOWNLOAD")
	res = append(res, "STATIC_WEB_SITE")

	return
}

// Utility function for StacksNames
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (stacksnames StacksNames) ToString() (res string) {

	// migration of former implementation of enum
	switch stacksnames {
	// insertion code per enum code
	case Phylotaxy:
		res = "phyllotaxymusic"
	}
	return
}

func (stacksnames *StacksNames) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "phyllotaxymusic":
		*stacksnames = Phylotaxy
		return
	default:
		return errUnkownEnum
	}
}

func (stacksnames *StacksNames) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Phylotaxy":
		*stacksnames = Phylotaxy
	default:
		err = errUnkownEnum
	}
	return
}

func (stacksnames *StacksNames) ToCodeString() (res string) {

	switch *stacksnames {
	// insertion code per enum code
	case Phylotaxy:
		res = "Phylotaxy"
	}
	return
}

func (stacksnames StacksNames) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Phylotaxy")

	return
}

func (stacksnames StacksNames) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "phyllotaxymusic")

	return
}

// Utility function for Target
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (target Target) ToString() (res string) {

	// migration of former implementation of enum
	switch target {
	// insertion code per enum code
	case FILE:
		res = "FILE"
	case WEB:
		res = "WEB"
	}
	return
}

func (target *Target) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "FILE":
		*target = FILE
		return
	case "WEB":
		*target = WEB
		return
	default:
		return errUnkownEnum
	}
}

func (target *Target) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "FILE":
		*target = FILE
	case "WEB":
		*target = WEB
	default:
		err = errUnkownEnum
	}
	return
}

func (target *Target) ToCodeString() (res string) {

	switch *target {
	// insertion code per enum code
	case FILE:
		res = "FILE"
	case WEB:
		res = "WEB"
	}
	return
}

func (target Target) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "FILE")
	res = append(res, "WEB")

	return
}

func (target Target) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "FILE")
	res = append(res, "WEB")

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
		return
	default:
		return errUnkownEnum
	}
}

func (treenames *TreeNames) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Sidebar":
		*treenames = Sidebar
	default:
		err = errUnkownEnum
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
