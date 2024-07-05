// generated code - do not edit

//insertion point for imports
import { LineAPI } from './line-api'

import { ParameterAPI } from './parameter-api'

import { RhombusAPI } from './rhombus-api'


export class BackRepoData {
	// insertion point for declarations
	LineAPIs = new Array<LineAPI>()

	ParameterAPIs = new Array<ParameterAPI>()

	RhombusAPIs = new Array<RhombusAPI>()



	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.LineAPIs = data?.LineAPIs || [];

		this.ParameterAPIs = data?.ParameterAPIs || [];

		this.RhombusAPIs = data?.RhombusAPIs || [];

	}

}