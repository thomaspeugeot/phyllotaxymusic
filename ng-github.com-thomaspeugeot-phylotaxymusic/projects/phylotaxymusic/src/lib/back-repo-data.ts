// generated code - do not edit

//insertion point for imports
import { CircleAPI } from './circle-api'

import { HorizontalAxisAPI } from './horizontalaxis-api'

import { InitialAxisAPI } from './initialaxis-api'

import { LineAPI } from './line-api'

import { ParameterAPI } from './parameter-api'

import { RhombusAPI } from './rhombus-api'

import { RhombusGridAPI } from './rhombusgrid-api'

import { VerticalAxisAPI } from './verticalaxis-api'


export class BackRepoData {
	// insertion point for declarations
	CircleAPIs = new Array<CircleAPI>()

	HorizontalAxisAPIs = new Array<HorizontalAxisAPI>()

	InitialAxisAPIs = new Array<InitialAxisAPI>()

	LineAPIs = new Array<LineAPI>()

	ParameterAPIs = new Array<ParameterAPI>()

	RhombusAPIs = new Array<RhombusAPI>()

	RhombusGridAPIs = new Array<RhombusGridAPI>()

	VerticalAxisAPIs = new Array<VerticalAxisAPI>()



	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.CircleAPIs = data?.CircleAPIs || [];

		this.HorizontalAxisAPIs = data?.HorizontalAxisAPIs || [];

		this.InitialAxisAPIs = data?.InitialAxisAPIs || [];

		this.LineAPIs = data?.LineAPIs || [];

		this.ParameterAPIs = data?.ParameterAPIs || [];

		this.RhombusAPIs = data?.RhombusAPIs || [];

		this.RhombusGridAPIs = data?.RhombusGridAPIs || [];

		this.VerticalAxisAPIs = data?.VerticalAxisAPIs || [];

	}

}