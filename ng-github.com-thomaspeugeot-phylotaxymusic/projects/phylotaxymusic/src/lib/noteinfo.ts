// generated code - do not edit

import { NoteInfoAPI } from './noteinfo-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class NoteInfo {

	static GONGSTRUCT_NAME = "NoteInfo"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsSkipped: boolean = false

	// insertion point for pointers and slices of pointers declarations
}

export function CopyNoteInfoToNoteInfoAPI(noteinfo: NoteInfo, noteinfoAPI: NoteInfoAPI) {

	noteinfoAPI.CreatedAt = noteinfo.CreatedAt
	noteinfoAPI.DeletedAt = noteinfo.DeletedAt
	noteinfoAPI.ID = noteinfo.ID

	// insertion point for basic fields copy operations
	noteinfoAPI.Name = noteinfo.Name
	noteinfoAPI.IsSkipped = noteinfo.IsSkipped

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyNoteInfoAPIToNoteInfo update basic, pointers and slice of pointers fields of noteinfo
// from respectively the basic fields and encoded fields of pointers and slices of pointers of noteinfoAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyNoteInfoAPIToNoteInfo(noteinfoAPI: NoteInfoAPI, noteinfo: NoteInfo, frontRepo: FrontRepo) {

	noteinfo.CreatedAt = noteinfoAPI.CreatedAt
	noteinfo.DeletedAt = noteinfoAPI.DeletedAt
	noteinfo.ID = noteinfoAPI.ID

	// insertion point for basic fields copy operations
	noteinfo.Name = noteinfoAPI.Name
	noteinfo.IsSkipped = noteinfoAPI.IsSkipped

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
