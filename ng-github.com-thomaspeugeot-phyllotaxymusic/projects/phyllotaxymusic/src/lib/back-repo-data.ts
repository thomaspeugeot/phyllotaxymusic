// generated code - do not edit

//insertion point for imports
import { AxisAPI } from './axis-api'

import { AxisGridAPI } from './axisgrid-api'

import { BezierAPI } from './bezier-api'

import { BezierGridAPI } from './beziergrid-api'

import { BezierGridStackAPI } from './beziergridstack-api'

import { CircleAPI } from './circle-api'

import { CircleGridAPI } from './circlegrid-api'

import { ExportToMusicxmlAPI } from './exporttomusicxml-api'

import { FrontCurveAPI } from './frontcurve-api'

import { FrontCurveStackAPI } from './frontcurvestack-api'

import { HorizontalAxisAPI } from './horizontalaxis-api'

import { KeyAPI } from './key-api'

import { ParameterAPI } from './parameter-api'

import { RhombusAPI } from './rhombus-api'

import { RhombusGridAPI } from './rhombusgrid-api'

import { ShapeCategoryAPI } from './shapecategory-api'

import { SpiralBezierAPI } from './spiralbezier-api'

import { SpiralBezierGridAPI } from './spiralbeziergrid-api'

import { SpiralCircleAPI } from './spiralcircle-api'

import { SpiralCircleGridAPI } from './spiralcirclegrid-api'

import { SpiralLineAPI } from './spiralline-api'

import { SpiralLineGridAPI } from './spirallinegrid-api'

import { SpiralOriginAPI } from './spiralorigin-api'

import { SpiralRhombusAPI } from './spiralrhombus-api'

import { SpiralRhombusGridAPI } from './spiralrhombusgrid-api'

import { VerticalAxisAPI } from './verticalaxis-api'


export class BackRepoData {
	// insertion point for declarations
	AxisAPIs = new Array<AxisAPI>()

	AxisGridAPIs = new Array<AxisGridAPI>()

	BezierAPIs = new Array<BezierAPI>()

	BezierGridAPIs = new Array<BezierGridAPI>()

	BezierGridStackAPIs = new Array<BezierGridStackAPI>()

	CircleAPIs = new Array<CircleAPI>()

	CircleGridAPIs = new Array<CircleGridAPI>()

	ExportToMusicxmlAPIs = new Array<ExportToMusicxmlAPI>()

	FrontCurveAPIs = new Array<FrontCurveAPI>()

	FrontCurveStackAPIs = new Array<FrontCurveStackAPI>()

	HorizontalAxisAPIs = new Array<HorizontalAxisAPI>()

	KeyAPIs = new Array<KeyAPI>()

	ParameterAPIs = new Array<ParameterAPI>()

	RhombusAPIs = new Array<RhombusAPI>()

	RhombusGridAPIs = new Array<RhombusGridAPI>()

	ShapeCategoryAPIs = new Array<ShapeCategoryAPI>()

	SpiralBezierAPIs = new Array<SpiralBezierAPI>()

	SpiralBezierGridAPIs = new Array<SpiralBezierGridAPI>()

	SpiralCircleAPIs = new Array<SpiralCircleAPI>()

	SpiralCircleGridAPIs = new Array<SpiralCircleGridAPI>()

	SpiralLineAPIs = new Array<SpiralLineAPI>()

	SpiralLineGridAPIs = new Array<SpiralLineGridAPI>()

	SpiralOriginAPIs = new Array<SpiralOriginAPI>()

	SpiralRhombusAPIs = new Array<SpiralRhombusAPI>()

	SpiralRhombusGridAPIs = new Array<SpiralRhombusGridAPI>()

	VerticalAxisAPIs = new Array<VerticalAxisAPI>()


	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index : number

	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.AxisAPIs = data?.AxisAPIs || [];

		this.AxisGridAPIs = data?.AxisGridAPIs || [];

		this.BezierAPIs = data?.BezierAPIs || [];

		this.BezierGridAPIs = data?.BezierGridAPIs || [];

		this.BezierGridStackAPIs = data?.BezierGridStackAPIs || [];

		this.CircleAPIs = data?.CircleAPIs || [];

		this.CircleGridAPIs = data?.CircleGridAPIs || [];

		this.ExportToMusicxmlAPIs = data?.ExportToMusicxmlAPIs || [];

		this.FrontCurveAPIs = data?.FrontCurveAPIs || [];

		this.FrontCurveStackAPIs = data?.FrontCurveStackAPIs || [];

		this.HorizontalAxisAPIs = data?.HorizontalAxisAPIs || [];

		this.KeyAPIs = data?.KeyAPIs || [];

		this.ParameterAPIs = data?.ParameterAPIs || [];

		this.RhombusAPIs = data?.RhombusAPIs || [];

		this.RhombusGridAPIs = data?.RhombusGridAPIs || [];

		this.ShapeCategoryAPIs = data?.ShapeCategoryAPIs || [];

		this.SpiralBezierAPIs = data?.SpiralBezierAPIs || [];

		this.SpiralBezierGridAPIs = data?.SpiralBezierGridAPIs || [];

		this.SpiralCircleAPIs = data?.SpiralCircleAPIs || [];

		this.SpiralCircleGridAPIs = data?.SpiralCircleGridAPIs || [];

		this.SpiralLineAPIs = data?.SpiralLineAPIs || [];

		this.SpiralLineGridAPIs = data?.SpiralLineGridAPIs || [];

		this.SpiralOriginAPIs = data?.SpiralOriginAPIs || [];

		this.SpiralRhombusAPIs = data?.SpiralRhombusAPIs || [];

		this.SpiralRhombusGridAPIs = data?.SpiralRhombusGridAPIs || [];

		this.VerticalAxisAPIs = data?.VerticalAxisAPIs || [];

		this.GONG__Index = data?.GONG__Index ?? -1;   // Assign Index here
	}

}