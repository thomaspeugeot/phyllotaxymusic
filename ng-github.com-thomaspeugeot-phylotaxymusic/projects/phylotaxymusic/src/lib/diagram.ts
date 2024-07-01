// generated code - do not edit

import { DiagramAPI } from './diagram-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Diagram {

	static GONGSTRUCT_NAME = "Diagram"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyDiagramToDiagramAPI(diagram: Diagram, diagramAPI: DiagramAPI) {

	diagramAPI.CreatedAt = diagram.CreatedAt
	diagramAPI.DeletedAt = diagram.DeletedAt
	diagramAPI.ID = diagram.ID

	// insertion point for basic fields copy operations
	diagramAPI.Name = diagram.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyDiagramAPIToDiagram update basic, pointers and slice of pointers fields of diagram
// from respectively the basic fields and encoded fields of pointers and slices of pointers of diagramAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyDiagramAPIToDiagram(diagramAPI: DiagramAPI, diagram: Diagram, frontRepo: FrontRepo) {

	diagram.CreatedAt = diagramAPI.CreatedAt
	diagram.DeletedAt = diagramAPI.DeletedAt
	diagram.ID = diagramAPI.ID

	// insertion point for basic fields copy operations
	diagram.Name = diagramAPI.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
