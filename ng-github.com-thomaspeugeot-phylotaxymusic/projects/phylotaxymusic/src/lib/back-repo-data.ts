// generated code - do not edit

//insertion point for imports
import { DiagramAPI } from './diagram-api'

import { LineAPI } from './line-api'


export class BackRepoData {
	// insertion point for declarations
	DiagramAPIs = new Array<DiagramAPI>()

	LineAPIs = new Array<LineAPI>()



	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.DiagramAPIs = data?.DiagramAPIs || [];

		this.LineAPIs = data?.LineAPIs || [];

	}

}