// generated code - do not edit

//insertion point for imports
import { AxisAPI } from './axis-api'

import { CircleAPI } from './circle-api'

import { CircleGridAPI } from './circlegrid-api'

import { HorizontalAxisAPI } from './horizontalaxis-api'

import { LineAPI } from './line-api'

import { ParameterAPI } from './parameter-api'

import { RhombusAPI } from './rhombus-api'

import { RhombusGridAPI } from './rhombusgrid-api'

import { VerticalAxisAPI } from './verticalaxis-api'


export class BackRepoData {
	// insertion point for declarations
	AxisAPIs = new Array<AxisAPI>()

	CircleAPIs = new Array<CircleAPI>()

	CircleGridAPIs = new Array<CircleGridAPI>()

	HorizontalAxisAPIs = new Array<HorizontalAxisAPI>()

	LineAPIs = new Array<LineAPI>()

	ParameterAPIs = new Array<ParameterAPI>()

	RhombusAPIs = new Array<RhombusAPI>()

	RhombusGridAPIs = new Array<RhombusGridAPI>()

	VerticalAxisAPIs = new Array<VerticalAxisAPI>()



	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.AxisAPIs = data?.AxisAPIs || [];

		this.CircleAPIs = data?.CircleAPIs || [];

		this.CircleGridAPIs = data?.CircleGridAPIs || [];

		this.HorizontalAxisAPIs = data?.HorizontalAxisAPIs || [];

		this.LineAPIs = data?.LineAPIs || [];

		this.ParameterAPIs = data?.ParameterAPIs || [];

		this.RhombusAPIs = data?.RhombusAPIs || [];

		this.RhombusGridAPIs = data?.RhombusGridAPIs || [];

		this.VerticalAxisAPIs = data?.VerticalAxisAPIs || [];

	}

}