// generated code - do not edit

import { ExportToMusicxmlAPI } from './exporttomusicxml-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ExportToMusicxml {

	static GONGSTRUCT_NAME = "ExportToMusicxml"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyExportToMusicxmlToExportToMusicxmlAPI(exporttomusicxml: ExportToMusicxml, exporttomusicxmlAPI: ExportToMusicxmlAPI) {

	exporttomusicxmlAPI.CreatedAt = exporttomusicxml.CreatedAt
	exporttomusicxmlAPI.DeletedAt = exporttomusicxml.DeletedAt
	exporttomusicxmlAPI.ID = exporttomusicxml.ID

	// insertion point for basic fields copy operations
	exporttomusicxmlAPI.Name = exporttomusicxml.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyExportToMusicxmlAPIToExportToMusicxml update basic, pointers and slice of pointers fields of exporttomusicxml
// from respectively the basic fields and encoded fields of pointers and slices of pointers of exporttomusicxmlAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyExportToMusicxmlAPIToExportToMusicxml(exporttomusicxmlAPI: ExportToMusicxmlAPI, exporttomusicxml: ExportToMusicxml, frontRepo: FrontRepo) {

	exporttomusicxml.CreatedAt = exporttomusicxmlAPI.CreatedAt
	exporttomusicxml.DeletedAt = exporttomusicxmlAPI.DeletedAt
	exporttomusicxml.ID = exporttomusicxmlAPI.ID

	// insertion point for basic fields copy operations
	exporttomusicxml.Name = exporttomusicxmlAPI.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
