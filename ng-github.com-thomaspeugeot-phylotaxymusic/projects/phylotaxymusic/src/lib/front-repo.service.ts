import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'

// insertion point sub template for services imports
import { AxisAPI } from './axis-api'
import { Axis, CopyAxisAPIToAxis } from './axis'
import { AxisService } from './axis.service'

import { AxisGridAPI } from './axisgrid-api'
import { AxisGrid, CopyAxisGridAPIToAxisGrid } from './axisgrid'
import { AxisGridService } from './axisgrid.service'

import { BezierAPI } from './bezier-api'
import { Bezier, CopyBezierAPIToBezier } from './bezier'
import { BezierService } from './bezier.service'

import { BezierGridAPI } from './beziergrid-api'
import { BezierGrid, CopyBezierGridAPIToBezierGrid } from './beziergrid'
import { BezierGridService } from './beziergrid.service'

import { BezierGridStackAPI } from './beziergridstack-api'
import { BezierGridStack, CopyBezierGridStackAPIToBezierGridStack } from './beziergridstack'
import { BezierGridStackService } from './beziergridstack.service'

import { CircleAPI } from './circle-api'
import { Circle, CopyCircleAPIToCircle } from './circle'
import { CircleService } from './circle.service'

import { CircleGridAPI } from './circlegrid-api'
import { CircleGrid, CopyCircleGridAPIToCircleGrid } from './circlegrid'
import { CircleGridService } from './circlegrid.service'

import { FrontCurveAPI } from './frontcurve-api'
import { FrontCurve, CopyFrontCurveAPIToFrontCurve } from './frontcurve'
import { FrontCurveService } from './frontcurve.service'

import { FrontCurveStackAPI } from './frontcurvestack-api'
import { FrontCurveStack, CopyFrontCurveStackAPIToFrontCurveStack } from './frontcurvestack'
import { FrontCurveStackService } from './frontcurvestack.service'

import { HorizontalAxisAPI } from './horizontalaxis-api'
import { HorizontalAxis, CopyHorizontalAxisAPIToHorizontalAxis } from './horizontalaxis'
import { HorizontalAxisService } from './horizontalaxis.service'

import { KeyAPI } from './key-api'
import { Key, CopyKeyAPIToKey } from './key'
import { KeyService } from './key.service'

import { NoteInfoAPI } from './noteinfo-api'
import { NoteInfo, CopyNoteInfoAPIToNoteInfo } from './noteinfo'
import { NoteInfoService } from './noteinfo.service'

import { ParameterAPI } from './parameter-api'
import { Parameter, CopyParameterAPIToParameter } from './parameter'
import { ParameterService } from './parameter.service'

import { RhombusAPI } from './rhombus-api'
import { Rhombus, CopyRhombusAPIToRhombus } from './rhombus'
import { RhombusService } from './rhombus.service'

import { RhombusGridAPI } from './rhombusgrid-api'
import { RhombusGrid, CopyRhombusGridAPIToRhombusGrid } from './rhombusgrid'
import { RhombusGridService } from './rhombusgrid.service'

import { ShapeCategoryAPI } from './shapecategory-api'
import { ShapeCategory, CopyShapeCategoryAPIToShapeCategory } from './shapecategory'
import { ShapeCategoryService } from './shapecategory.service'

import { SpiralBezierAPI } from './spiralbezier-api'
import { SpiralBezier, CopySpiralBezierAPIToSpiralBezier } from './spiralbezier'
import { SpiralBezierService } from './spiralbezier.service'

import { SpiralBezierGridAPI } from './spiralbeziergrid-api'
import { SpiralBezierGrid, CopySpiralBezierGridAPIToSpiralBezierGrid } from './spiralbeziergrid'
import { SpiralBezierGridService } from './spiralbeziergrid.service'

import { SpiralCircleAPI } from './spiralcircle-api'
import { SpiralCircle, CopySpiralCircleAPIToSpiralCircle } from './spiralcircle'
import { SpiralCircleService } from './spiralcircle.service'

import { SpiralCircleGridAPI } from './spiralcirclegrid-api'
import { SpiralCircleGrid, CopySpiralCircleGridAPIToSpiralCircleGrid } from './spiralcirclegrid'
import { SpiralCircleGridService } from './spiralcirclegrid.service'

import { SpiralLineAPI } from './spiralline-api'
import { SpiralLine, CopySpiralLineAPIToSpiralLine } from './spiralline'
import { SpiralLineService } from './spiralline.service'

import { SpiralLineGridAPI } from './spirallinegrid-api'
import { SpiralLineGrid, CopySpiralLineGridAPIToSpiralLineGrid } from './spirallinegrid'
import { SpiralLineGridService } from './spirallinegrid.service'

import { SpiralOriginAPI } from './spiralorigin-api'
import { SpiralOrigin, CopySpiralOriginAPIToSpiralOrigin } from './spiralorigin'
import { SpiralOriginService } from './spiralorigin.service'

import { SpiralRhombusAPI } from './spiralrhombus-api'
import { SpiralRhombus, CopySpiralRhombusAPIToSpiralRhombus } from './spiralrhombus'
import { SpiralRhombusService } from './spiralrhombus.service'

import { SpiralRhombusGridAPI } from './spiralrhombusgrid-api'
import { SpiralRhombusGrid, CopySpiralRhombusGridAPIToSpiralRhombusGrid } from './spiralrhombusgrid'
import { SpiralRhombusGridService } from './spiralrhombusgrid.service'

import { VerticalAxisAPI } from './verticalaxis-api'
import { VerticalAxis, CopyVerticalAxisAPIToVerticalAxis } from './verticalaxis'
import { VerticalAxisService } from './verticalaxis.service'


import { BackRepoData } from './back-repo-data'

export const StackType = "github.com/thomaspeugeot/phylotaxymusic/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
	array_Axiss = new Array<Axis>() // array of front instances
	map_ID_Axis = new Map<number, Axis>() // map of front instances

	array_AxisGrids = new Array<AxisGrid>() // array of front instances
	map_ID_AxisGrid = new Map<number, AxisGrid>() // map of front instances

	array_Beziers = new Array<Bezier>() // array of front instances
	map_ID_Bezier = new Map<number, Bezier>() // map of front instances

	array_BezierGrids = new Array<BezierGrid>() // array of front instances
	map_ID_BezierGrid = new Map<number, BezierGrid>() // map of front instances

	array_BezierGridStacks = new Array<BezierGridStack>() // array of front instances
	map_ID_BezierGridStack = new Map<number, BezierGridStack>() // map of front instances

	array_Circles = new Array<Circle>() // array of front instances
	map_ID_Circle = new Map<number, Circle>() // map of front instances

	array_CircleGrids = new Array<CircleGrid>() // array of front instances
	map_ID_CircleGrid = new Map<number, CircleGrid>() // map of front instances

	array_FrontCurves = new Array<FrontCurve>() // array of front instances
	map_ID_FrontCurve = new Map<number, FrontCurve>() // map of front instances

	array_FrontCurveStacks = new Array<FrontCurveStack>() // array of front instances
	map_ID_FrontCurveStack = new Map<number, FrontCurveStack>() // map of front instances

	array_HorizontalAxiss = new Array<HorizontalAxis>() // array of front instances
	map_ID_HorizontalAxis = new Map<number, HorizontalAxis>() // map of front instances

	array_Keys = new Array<Key>() // array of front instances
	map_ID_Key = new Map<number, Key>() // map of front instances

	array_NoteInfos = new Array<NoteInfo>() // array of front instances
	map_ID_NoteInfo = new Map<number, NoteInfo>() // map of front instances

	array_Parameters = new Array<Parameter>() // array of front instances
	map_ID_Parameter = new Map<number, Parameter>() // map of front instances

	array_Rhombuss = new Array<Rhombus>() // array of front instances
	map_ID_Rhombus = new Map<number, Rhombus>() // map of front instances

	array_RhombusGrids = new Array<RhombusGrid>() // array of front instances
	map_ID_RhombusGrid = new Map<number, RhombusGrid>() // map of front instances

	array_ShapeCategorys = new Array<ShapeCategory>() // array of front instances
	map_ID_ShapeCategory = new Map<number, ShapeCategory>() // map of front instances

	array_SpiralBeziers = new Array<SpiralBezier>() // array of front instances
	map_ID_SpiralBezier = new Map<number, SpiralBezier>() // map of front instances

	array_SpiralBezierGrids = new Array<SpiralBezierGrid>() // array of front instances
	map_ID_SpiralBezierGrid = new Map<number, SpiralBezierGrid>() // map of front instances

	array_SpiralCircles = new Array<SpiralCircle>() // array of front instances
	map_ID_SpiralCircle = new Map<number, SpiralCircle>() // map of front instances

	array_SpiralCircleGrids = new Array<SpiralCircleGrid>() // array of front instances
	map_ID_SpiralCircleGrid = new Map<number, SpiralCircleGrid>() // map of front instances

	array_SpiralLines = new Array<SpiralLine>() // array of front instances
	map_ID_SpiralLine = new Map<number, SpiralLine>() // map of front instances

	array_SpiralLineGrids = new Array<SpiralLineGrid>() // array of front instances
	map_ID_SpiralLineGrid = new Map<number, SpiralLineGrid>() // map of front instances

	array_SpiralOrigins = new Array<SpiralOrigin>() // array of front instances
	map_ID_SpiralOrigin = new Map<number, SpiralOrigin>() // map of front instances

	array_SpiralRhombuss = new Array<SpiralRhombus>() // array of front instances
	map_ID_SpiralRhombus = new Map<number, SpiralRhombus>() // map of front instances

	array_SpiralRhombusGrids = new Array<SpiralRhombusGrid>() // array of front instances
	map_ID_SpiralRhombusGrid = new Map<number, SpiralRhombusGrid>() // map of front instances

	array_VerticalAxiss = new Array<VerticalAxis>() // array of front instances
	map_ID_VerticalAxis = new Map<number, VerticalAxis>() // map of front instances


	// getFrontArray allows for a get function that is robust to refactoring of the named struct name
	// for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
	// contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
	getFrontArray<Type>(gongStructName: string): Array<Type> {
		switch (gongStructName) {
			// insertion point
			case 'Axis':
				return this.array_Axiss as unknown as Array<Type>
			case 'AxisGrid':
				return this.array_AxisGrids as unknown as Array<Type>
			case 'Bezier':
				return this.array_Beziers as unknown as Array<Type>
			case 'BezierGrid':
				return this.array_BezierGrids as unknown as Array<Type>
			case 'BezierGridStack':
				return this.array_BezierGridStacks as unknown as Array<Type>
			case 'Circle':
				return this.array_Circles as unknown as Array<Type>
			case 'CircleGrid':
				return this.array_CircleGrids as unknown as Array<Type>
			case 'FrontCurve':
				return this.array_FrontCurves as unknown as Array<Type>
			case 'FrontCurveStack':
				return this.array_FrontCurveStacks as unknown as Array<Type>
			case 'HorizontalAxis':
				return this.array_HorizontalAxiss as unknown as Array<Type>
			case 'Key':
				return this.array_Keys as unknown as Array<Type>
			case 'NoteInfo':
				return this.array_NoteInfos as unknown as Array<Type>
			case 'Parameter':
				return this.array_Parameters as unknown as Array<Type>
			case 'Rhombus':
				return this.array_Rhombuss as unknown as Array<Type>
			case 'RhombusGrid':
				return this.array_RhombusGrids as unknown as Array<Type>
			case 'ShapeCategory':
				return this.array_ShapeCategorys as unknown as Array<Type>
			case 'SpiralBezier':
				return this.array_SpiralBeziers as unknown as Array<Type>
			case 'SpiralBezierGrid':
				return this.array_SpiralBezierGrids as unknown as Array<Type>
			case 'SpiralCircle':
				return this.array_SpiralCircles as unknown as Array<Type>
			case 'SpiralCircleGrid':
				return this.array_SpiralCircleGrids as unknown as Array<Type>
			case 'SpiralLine':
				return this.array_SpiralLines as unknown as Array<Type>
			case 'SpiralLineGrid':
				return this.array_SpiralLineGrids as unknown as Array<Type>
			case 'SpiralOrigin':
				return this.array_SpiralOrigins as unknown as Array<Type>
			case 'SpiralRhombus':
				return this.array_SpiralRhombuss as unknown as Array<Type>
			case 'SpiralRhombusGrid':
				return this.array_SpiralRhombusGrids as unknown as Array<Type>
			case 'VerticalAxis':
				return this.array_VerticalAxiss as unknown as Array<Type>
			default:
				throw new Error("Type not recognized");
		}
	}

	getFrontMap<Type>(gongStructName: string): Map<number, Type> {
		switch (gongStructName) {
			// insertion point
			case 'Axis':
				return this.map_ID_Axis as unknown as Map<number, Type>
			case 'AxisGrid':
				return this.map_ID_AxisGrid as unknown as Map<number, Type>
			case 'Bezier':
				return this.map_ID_Bezier as unknown as Map<number, Type>
			case 'BezierGrid':
				return this.map_ID_BezierGrid as unknown as Map<number, Type>
			case 'BezierGridStack':
				return this.map_ID_BezierGridStack as unknown as Map<number, Type>
			case 'Circle':
				return this.map_ID_Circle as unknown as Map<number, Type>
			case 'CircleGrid':
				return this.map_ID_CircleGrid as unknown as Map<number, Type>
			case 'FrontCurve':
				return this.map_ID_FrontCurve as unknown as Map<number, Type>
			case 'FrontCurveStack':
				return this.map_ID_FrontCurveStack as unknown as Map<number, Type>
			case 'HorizontalAxis':
				return this.map_ID_HorizontalAxis as unknown as Map<number, Type>
			case 'Key':
				return this.map_ID_Key as unknown as Map<number, Type>
			case 'NoteInfo':
				return this.map_ID_NoteInfo as unknown as Map<number, Type>
			case 'Parameter':
				return this.map_ID_Parameter as unknown as Map<number, Type>
			case 'Rhombus':
				return this.map_ID_Rhombus as unknown as Map<number, Type>
			case 'RhombusGrid':
				return this.map_ID_RhombusGrid as unknown as Map<number, Type>
			case 'ShapeCategory':
				return this.map_ID_ShapeCategory as unknown as Map<number, Type>
			case 'SpiralBezier':
				return this.map_ID_SpiralBezier as unknown as Map<number, Type>
			case 'SpiralBezierGrid':
				return this.map_ID_SpiralBezierGrid as unknown as Map<number, Type>
			case 'SpiralCircle':
				return this.map_ID_SpiralCircle as unknown as Map<number, Type>
			case 'SpiralCircleGrid':
				return this.map_ID_SpiralCircleGrid as unknown as Map<number, Type>
			case 'SpiralLine':
				return this.map_ID_SpiralLine as unknown as Map<number, Type>
			case 'SpiralLineGrid':
				return this.map_ID_SpiralLineGrid as unknown as Map<number, Type>
			case 'SpiralOrigin':
				return this.map_ID_SpiralOrigin as unknown as Map<number, Type>
			case 'SpiralRhombus':
				return this.map_ID_SpiralRhombus as unknown as Map<number, Type>
			case 'SpiralRhombusGrid':
				return this.map_ID_SpiralRhombusGrid as unknown as Map<number, Type>
			case 'VerticalAxis':
				return this.map_ID_VerticalAxis as unknown as Map<number, Type>
			default:
				throw new Error("Type not recognized");
		}
	}
}

// the table component is called in different ways
//
// DISPLAY or ASSOCIATION MODE
//
// in ASSOCIATION MODE, it is invoked within a diaglo and a Dialog Data item is used to
// configure the component
// DialogData define the interface for information that is forwarded from the calling instance to 
// the select table
export class DialogData {
	ID: number = 0 // ID of the calling instance

	// the reverse pointer is the name of the generated field on the destination
	// struct of the ONE-MANY association
	ReversePointer: string = "" // field of {{Structname}} that serve as reverse pointer
	OrderingMode: boolean = false // if true, this is for ordering items

	// there are different selection mode : ONE_MANY or MANY_MANY
	SelectionMode: SelectionMode = SelectionMode.ONE_MANY_ASSOCIATION_MODE

	// used if SelectionMode is MANY_MANY_ASSOCIATION_MODE
	//
	// In Gong, a MANY-MANY association is implemented as a ONE-ZERO/ONE followed by a ONE_MANY association
	// 
	// in the MANY_MANY_ASSOCIATION_MODE case, we need also the Struct and the FieldName that are
	// at the end of the ONE-MANY association
	SourceStruct: string = ""	// The "Aclass"
	SourceField: string = "" // the "AnarrayofbUse"
	IntermediateStruct: string = "" // the "AclassBclassUse" 
	IntermediateStructField: string = "" // the "Bclass" as field
	NextAssociationStruct: string = "" // the "Bclass"

	GONG__StackPath: string = ""
}

export enum SelectionMode {
	ONE_MANY_ASSOCIATION_MODE = "ONE_MANY_ASSOCIATION_MODE",
	MANY_MANY_ASSOCIATION_MODE = "MANY_MANY_ASSOCIATION_MODE",
}

//
// observable that fetch all elements of the stack and store them in the FrontRepo
//
@Injectable({
	providedIn: 'root'
})
export class FrontRepoService {

	GONG__StackPath: string = ""
	private socket: WebSocket | undefined

	httpOptions = {
		headers: new HttpHeaders({ 'Content-Type': 'application/json' })
	};

	//
	// Store of all instances of the stack
	//
	frontRepo = new (FrontRepo)

	constructor(
		private http: HttpClient, // insertion point sub template 
		private axisService: AxisService,
		private axisgridService: AxisGridService,
		private bezierService: BezierService,
		private beziergridService: BezierGridService,
		private beziergridstackService: BezierGridStackService,
		private circleService: CircleService,
		private circlegridService: CircleGridService,
		private frontcurveService: FrontCurveService,
		private frontcurvestackService: FrontCurveStackService,
		private horizontalaxisService: HorizontalAxisService,
		private keyService: KeyService,
		private noteinfoService: NoteInfoService,
		private parameterService: ParameterService,
		private rhombusService: RhombusService,
		private rhombusgridService: RhombusGridService,
		private shapecategoryService: ShapeCategoryService,
		private spiralbezierService: SpiralBezierService,
		private spiralbeziergridService: SpiralBezierGridService,
		private spiralcircleService: SpiralCircleService,
		private spiralcirclegridService: SpiralCircleGridService,
		private spirallineService: SpiralLineService,
		private spirallinegridService: SpiralLineGridService,
		private spiraloriginService: SpiralOriginService,
		private spiralrhombusService: SpiralRhombusService,
		private spiralrhombusgridService: SpiralRhombusGridService,
		private verticalaxisService: VerticalAxisService,
	) { }

	// postService provides a post function for each struct name
	postService(structName: string, instanceToBePosted: any) {
		let service = this[structName.toLowerCase() + "Service" + "Service" as keyof FrontRepoService]
		let servicePostFunction = service[("post" + structName) as keyof typeof service] as (instance: typeof instanceToBePosted) => Observable<typeof instanceToBePosted>

		servicePostFunction(instanceToBePosted).subscribe(
			instance => {
				let behaviorSubject = instanceToBePosted[(structName + "ServiceChanged") as keyof typeof instanceToBePosted] as unknown as BehaviorSubject<string>
				behaviorSubject.next("post")
			}
		);
	}

	// deleteService provides a delete function for each struct name
	deleteService(structName: string, instanceToBeDeleted: any) {
		let service = this[structName.toLowerCase() + "Service" as keyof FrontRepoService]
		let serviceDeleteFunction = service["delete" + structName as keyof typeof service] as (instance: typeof instanceToBeDeleted) => Observable<typeof instanceToBeDeleted>

		serviceDeleteFunction(instanceToBeDeleted).subscribe(
			instance => {
				let behaviorSubject = instanceToBeDeleted[(structName + "ServiceChanged") as keyof typeof instanceToBeDeleted] as unknown as BehaviorSubject<string>
				behaviorSubject.next("delete")
			}
		);
	}

	// typing of observable can be messy in typescript. Therefore, one force the type
	observableFrontRepo: [
		Observable<null>, // see below for the of(null) observable
		// insertion point sub template 
		Observable<AxisAPI[]>,
		Observable<AxisGridAPI[]>,
		Observable<BezierAPI[]>,
		Observable<BezierGridAPI[]>,
		Observable<BezierGridStackAPI[]>,
		Observable<CircleAPI[]>,
		Observable<CircleGridAPI[]>,
		Observable<FrontCurveAPI[]>,
		Observable<FrontCurveStackAPI[]>,
		Observable<HorizontalAxisAPI[]>,
		Observable<KeyAPI[]>,
		Observable<NoteInfoAPI[]>,
		Observable<ParameterAPI[]>,
		Observable<RhombusAPI[]>,
		Observable<RhombusGridAPI[]>,
		Observable<ShapeCategoryAPI[]>,
		Observable<SpiralBezierAPI[]>,
		Observable<SpiralBezierGridAPI[]>,
		Observable<SpiralCircleAPI[]>,
		Observable<SpiralCircleGridAPI[]>,
		Observable<SpiralLineAPI[]>,
		Observable<SpiralLineGridAPI[]>,
		Observable<SpiralOriginAPI[]>,
		Observable<SpiralRhombusAPI[]>,
		Observable<SpiralRhombusGridAPI[]>,
		Observable<VerticalAxisAPI[]>,
	] = [
			// Using "combineLatest" with a placeholder observable.
			//
			// This allows the typescript compiler to pass when no GongStruct is present in the front API
			//
			// The "of(null)" is a "meaningless" observable that emits a single value (null) and completes.
			// This is used as a workaround to satisfy TypeScript requirements and the "combineLatest" 
			// expectation for a non-empty array of observables.
			of(null), // 
			// insertion point sub template
			this.axisService.getAxiss(this.GONG__StackPath, this.frontRepo),
			this.axisgridService.getAxisGrids(this.GONG__StackPath, this.frontRepo),
			this.bezierService.getBeziers(this.GONG__StackPath, this.frontRepo),
			this.beziergridService.getBezierGrids(this.GONG__StackPath, this.frontRepo),
			this.beziergridstackService.getBezierGridStacks(this.GONG__StackPath, this.frontRepo),
			this.circleService.getCircles(this.GONG__StackPath, this.frontRepo),
			this.circlegridService.getCircleGrids(this.GONG__StackPath, this.frontRepo),
			this.frontcurveService.getFrontCurves(this.GONG__StackPath, this.frontRepo),
			this.frontcurvestackService.getFrontCurveStacks(this.GONG__StackPath, this.frontRepo),
			this.horizontalaxisService.getHorizontalAxiss(this.GONG__StackPath, this.frontRepo),
			this.keyService.getKeys(this.GONG__StackPath, this.frontRepo),
			this.noteinfoService.getNoteInfos(this.GONG__StackPath, this.frontRepo),
			this.parameterService.getParameters(this.GONG__StackPath, this.frontRepo),
			this.rhombusService.getRhombuss(this.GONG__StackPath, this.frontRepo),
			this.rhombusgridService.getRhombusGrids(this.GONG__StackPath, this.frontRepo),
			this.shapecategoryService.getShapeCategorys(this.GONG__StackPath, this.frontRepo),
			this.spiralbezierService.getSpiralBeziers(this.GONG__StackPath, this.frontRepo),
			this.spiralbeziergridService.getSpiralBezierGrids(this.GONG__StackPath, this.frontRepo),
			this.spiralcircleService.getSpiralCircles(this.GONG__StackPath, this.frontRepo),
			this.spiralcirclegridService.getSpiralCircleGrids(this.GONG__StackPath, this.frontRepo),
			this.spirallineService.getSpiralLines(this.GONG__StackPath, this.frontRepo),
			this.spirallinegridService.getSpiralLineGrids(this.GONG__StackPath, this.frontRepo),
			this.spiraloriginService.getSpiralOrigins(this.GONG__StackPath, this.frontRepo),
			this.spiralrhombusService.getSpiralRhombuss(this.GONG__StackPath, this.frontRepo),
			this.spiralrhombusgridService.getSpiralRhombusGrids(this.GONG__StackPath, this.frontRepo),
			this.verticalaxisService.getVerticalAxiss(this.GONG__StackPath, this.frontRepo),
		];

	//
	// pull performs a GET on all struct of the stack and redeem association pointers 
	//
	// This is an observable. Therefore, the control flow forks with
	// - pull() return immediatly the observable
	// - the observable observer, if it subscribe, is called when all GET calls are performs
	pull(GONG__StackPath: string = ""): Observable<FrontRepo> {

		this.GONG__StackPath = GONG__StackPath

		this.observableFrontRepo = [
			of(null), // see above for justification
			// insertion point sub template
			this.axisService.getAxiss(this.GONG__StackPath, this.frontRepo),
			this.axisgridService.getAxisGrids(this.GONG__StackPath, this.frontRepo),
			this.bezierService.getBeziers(this.GONG__StackPath, this.frontRepo),
			this.beziergridService.getBezierGrids(this.GONG__StackPath, this.frontRepo),
			this.beziergridstackService.getBezierGridStacks(this.GONG__StackPath, this.frontRepo),
			this.circleService.getCircles(this.GONG__StackPath, this.frontRepo),
			this.circlegridService.getCircleGrids(this.GONG__StackPath, this.frontRepo),
			this.frontcurveService.getFrontCurves(this.GONG__StackPath, this.frontRepo),
			this.frontcurvestackService.getFrontCurveStacks(this.GONG__StackPath, this.frontRepo),
			this.horizontalaxisService.getHorizontalAxiss(this.GONG__StackPath, this.frontRepo),
			this.keyService.getKeys(this.GONG__StackPath, this.frontRepo),
			this.noteinfoService.getNoteInfos(this.GONG__StackPath, this.frontRepo),
			this.parameterService.getParameters(this.GONG__StackPath, this.frontRepo),
			this.rhombusService.getRhombuss(this.GONG__StackPath, this.frontRepo),
			this.rhombusgridService.getRhombusGrids(this.GONG__StackPath, this.frontRepo),
			this.shapecategoryService.getShapeCategorys(this.GONG__StackPath, this.frontRepo),
			this.spiralbezierService.getSpiralBeziers(this.GONG__StackPath, this.frontRepo),
			this.spiralbeziergridService.getSpiralBezierGrids(this.GONG__StackPath, this.frontRepo),
			this.spiralcircleService.getSpiralCircles(this.GONG__StackPath, this.frontRepo),
			this.spiralcirclegridService.getSpiralCircleGrids(this.GONG__StackPath, this.frontRepo),
			this.spirallineService.getSpiralLines(this.GONG__StackPath, this.frontRepo),
			this.spirallinegridService.getSpiralLineGrids(this.GONG__StackPath, this.frontRepo),
			this.spiraloriginService.getSpiralOrigins(this.GONG__StackPath, this.frontRepo),
			this.spiralrhombusService.getSpiralRhombuss(this.GONG__StackPath, this.frontRepo),
			this.spiralrhombusgridService.getSpiralRhombusGrids(this.GONG__StackPath, this.frontRepo),
			this.verticalaxisService.getVerticalAxiss(this.GONG__StackPath, this.frontRepo),
		]

		return new Observable<FrontRepo>(
			(observer) => {
				combineLatest(
					this.observableFrontRepo
				).subscribe(
					([
						___of_null, // see above for the explanation about of
						// insertion point sub template for declarations 
						axiss_,
						axisgrids_,
						beziers_,
						beziergrids_,
						beziergridstacks_,
						circles_,
						circlegrids_,
						frontcurves_,
						frontcurvestacks_,
						horizontalaxiss_,
						keys_,
						noteinfos_,
						parameters_,
						rhombuss_,
						rhombusgrids_,
						shapecategorys_,
						spiralbeziers_,
						spiralbeziergrids_,
						spiralcircles_,
						spiralcirclegrids_,
						spirallines_,
						spirallinegrids_,
						spiralorigins_,
						spiralrhombuss_,
						spiralrhombusgrids_,
						verticalaxiss_,
					]) => {
						let _this = this
						// Typing can be messy with many items. Therefore, type casting is necessary here
						// insertion point sub template for type casting 
						var axiss: AxisAPI[]
						axiss = axiss_ as AxisAPI[]
						var axisgrids: AxisGridAPI[]
						axisgrids = axisgrids_ as AxisGridAPI[]
						var beziers: BezierAPI[]
						beziers = beziers_ as BezierAPI[]
						var beziergrids: BezierGridAPI[]
						beziergrids = beziergrids_ as BezierGridAPI[]
						var beziergridstacks: BezierGridStackAPI[]
						beziergridstacks = beziergridstacks_ as BezierGridStackAPI[]
						var circles: CircleAPI[]
						circles = circles_ as CircleAPI[]
						var circlegrids: CircleGridAPI[]
						circlegrids = circlegrids_ as CircleGridAPI[]
						var frontcurves: FrontCurveAPI[]
						frontcurves = frontcurves_ as FrontCurveAPI[]
						var frontcurvestacks: FrontCurveStackAPI[]
						frontcurvestacks = frontcurvestacks_ as FrontCurveStackAPI[]
						var horizontalaxiss: HorizontalAxisAPI[]
						horizontalaxiss = horizontalaxiss_ as HorizontalAxisAPI[]
						var keys: KeyAPI[]
						keys = keys_ as KeyAPI[]
						var noteinfos: NoteInfoAPI[]
						noteinfos = noteinfos_ as NoteInfoAPI[]
						var parameters: ParameterAPI[]
						parameters = parameters_ as ParameterAPI[]
						var rhombuss: RhombusAPI[]
						rhombuss = rhombuss_ as RhombusAPI[]
						var rhombusgrids: RhombusGridAPI[]
						rhombusgrids = rhombusgrids_ as RhombusGridAPI[]
						var shapecategorys: ShapeCategoryAPI[]
						shapecategorys = shapecategorys_ as ShapeCategoryAPI[]
						var spiralbeziers: SpiralBezierAPI[]
						spiralbeziers = spiralbeziers_ as SpiralBezierAPI[]
						var spiralbeziergrids: SpiralBezierGridAPI[]
						spiralbeziergrids = spiralbeziergrids_ as SpiralBezierGridAPI[]
						var spiralcircles: SpiralCircleAPI[]
						spiralcircles = spiralcircles_ as SpiralCircleAPI[]
						var spiralcirclegrids: SpiralCircleGridAPI[]
						spiralcirclegrids = spiralcirclegrids_ as SpiralCircleGridAPI[]
						var spirallines: SpiralLineAPI[]
						spirallines = spirallines_ as SpiralLineAPI[]
						var spirallinegrids: SpiralLineGridAPI[]
						spirallinegrids = spirallinegrids_ as SpiralLineGridAPI[]
						var spiralorigins: SpiralOriginAPI[]
						spiralorigins = spiralorigins_ as SpiralOriginAPI[]
						var spiralrhombuss: SpiralRhombusAPI[]
						spiralrhombuss = spiralrhombuss_ as SpiralRhombusAPI[]
						var spiralrhombusgrids: SpiralRhombusGridAPI[]
						spiralrhombusgrids = spiralrhombusgrids_ as SpiralRhombusGridAPI[]
						var verticalaxiss: VerticalAxisAPI[]
						verticalaxiss = verticalaxiss_ as VerticalAxisAPI[]

						// 
						// First Step: init map of instances
						// insertion point sub template for init 
						// init the arrays
						this.frontRepo.array_Axiss = []
						this.frontRepo.map_ID_Axis.clear()

						axiss.forEach(
							axisAPI => {
								let axis = new Axis
								this.frontRepo.array_Axiss.push(axis)
								this.frontRepo.map_ID_Axis.set(axisAPI.ID, axis)
							}
						)

						// init the arrays
						this.frontRepo.array_AxisGrids = []
						this.frontRepo.map_ID_AxisGrid.clear()

						axisgrids.forEach(
							axisgridAPI => {
								let axisgrid = new AxisGrid
								this.frontRepo.array_AxisGrids.push(axisgrid)
								this.frontRepo.map_ID_AxisGrid.set(axisgridAPI.ID, axisgrid)
							}
						)

						// init the arrays
						this.frontRepo.array_Beziers = []
						this.frontRepo.map_ID_Bezier.clear()

						beziers.forEach(
							bezierAPI => {
								let bezier = new Bezier
								this.frontRepo.array_Beziers.push(bezier)
								this.frontRepo.map_ID_Bezier.set(bezierAPI.ID, bezier)
							}
						)

						// init the arrays
						this.frontRepo.array_BezierGrids = []
						this.frontRepo.map_ID_BezierGrid.clear()

						beziergrids.forEach(
							beziergridAPI => {
								let beziergrid = new BezierGrid
								this.frontRepo.array_BezierGrids.push(beziergrid)
								this.frontRepo.map_ID_BezierGrid.set(beziergridAPI.ID, beziergrid)
							}
						)

						// init the arrays
						this.frontRepo.array_BezierGridStacks = []
						this.frontRepo.map_ID_BezierGridStack.clear()

						beziergridstacks.forEach(
							beziergridstackAPI => {
								let beziergridstack = new BezierGridStack
								this.frontRepo.array_BezierGridStacks.push(beziergridstack)
								this.frontRepo.map_ID_BezierGridStack.set(beziergridstackAPI.ID, beziergridstack)
							}
						)

						// init the arrays
						this.frontRepo.array_Circles = []
						this.frontRepo.map_ID_Circle.clear()

						circles.forEach(
							circleAPI => {
								let circle = new Circle
								this.frontRepo.array_Circles.push(circle)
								this.frontRepo.map_ID_Circle.set(circleAPI.ID, circle)
							}
						)

						// init the arrays
						this.frontRepo.array_CircleGrids = []
						this.frontRepo.map_ID_CircleGrid.clear()

						circlegrids.forEach(
							circlegridAPI => {
								let circlegrid = new CircleGrid
								this.frontRepo.array_CircleGrids.push(circlegrid)
								this.frontRepo.map_ID_CircleGrid.set(circlegridAPI.ID, circlegrid)
							}
						)

						// init the arrays
						this.frontRepo.array_FrontCurves = []
						this.frontRepo.map_ID_FrontCurve.clear()

						frontcurves.forEach(
							frontcurveAPI => {
								let frontcurve = new FrontCurve
								this.frontRepo.array_FrontCurves.push(frontcurve)
								this.frontRepo.map_ID_FrontCurve.set(frontcurveAPI.ID, frontcurve)
							}
						)

						// init the arrays
						this.frontRepo.array_FrontCurveStacks = []
						this.frontRepo.map_ID_FrontCurveStack.clear()

						frontcurvestacks.forEach(
							frontcurvestackAPI => {
								let frontcurvestack = new FrontCurveStack
								this.frontRepo.array_FrontCurveStacks.push(frontcurvestack)
								this.frontRepo.map_ID_FrontCurveStack.set(frontcurvestackAPI.ID, frontcurvestack)
							}
						)

						// init the arrays
						this.frontRepo.array_HorizontalAxiss = []
						this.frontRepo.map_ID_HorizontalAxis.clear()

						horizontalaxiss.forEach(
							horizontalaxisAPI => {
								let horizontalaxis = new HorizontalAxis
								this.frontRepo.array_HorizontalAxiss.push(horizontalaxis)
								this.frontRepo.map_ID_HorizontalAxis.set(horizontalaxisAPI.ID, horizontalaxis)
							}
						)

						// init the arrays
						this.frontRepo.array_Keys = []
						this.frontRepo.map_ID_Key.clear()

						keys.forEach(
							keyAPI => {
								let key = new Key
								this.frontRepo.array_Keys.push(key)
								this.frontRepo.map_ID_Key.set(keyAPI.ID, key)
							}
						)

						// init the arrays
						this.frontRepo.array_NoteInfos = []
						this.frontRepo.map_ID_NoteInfo.clear()

						noteinfos.forEach(
							noteinfoAPI => {
								let noteinfo = new NoteInfo
								this.frontRepo.array_NoteInfos.push(noteinfo)
								this.frontRepo.map_ID_NoteInfo.set(noteinfoAPI.ID, noteinfo)
							}
						)

						// init the arrays
						this.frontRepo.array_Parameters = []
						this.frontRepo.map_ID_Parameter.clear()

						parameters.forEach(
							parameterAPI => {
								let parameter = new Parameter
								this.frontRepo.array_Parameters.push(parameter)
								this.frontRepo.map_ID_Parameter.set(parameterAPI.ID, parameter)
							}
						)

						// init the arrays
						this.frontRepo.array_Rhombuss = []
						this.frontRepo.map_ID_Rhombus.clear()

						rhombuss.forEach(
							rhombusAPI => {
								let rhombus = new Rhombus
								this.frontRepo.array_Rhombuss.push(rhombus)
								this.frontRepo.map_ID_Rhombus.set(rhombusAPI.ID, rhombus)
							}
						)

						// init the arrays
						this.frontRepo.array_RhombusGrids = []
						this.frontRepo.map_ID_RhombusGrid.clear()

						rhombusgrids.forEach(
							rhombusgridAPI => {
								let rhombusgrid = new RhombusGrid
								this.frontRepo.array_RhombusGrids.push(rhombusgrid)
								this.frontRepo.map_ID_RhombusGrid.set(rhombusgridAPI.ID, rhombusgrid)
							}
						)

						// init the arrays
						this.frontRepo.array_ShapeCategorys = []
						this.frontRepo.map_ID_ShapeCategory.clear()

						shapecategorys.forEach(
							shapecategoryAPI => {
								let shapecategory = new ShapeCategory
								this.frontRepo.array_ShapeCategorys.push(shapecategory)
								this.frontRepo.map_ID_ShapeCategory.set(shapecategoryAPI.ID, shapecategory)
							}
						)

						// init the arrays
						this.frontRepo.array_SpiralBeziers = []
						this.frontRepo.map_ID_SpiralBezier.clear()

						spiralbeziers.forEach(
							spiralbezierAPI => {
								let spiralbezier = new SpiralBezier
								this.frontRepo.array_SpiralBeziers.push(spiralbezier)
								this.frontRepo.map_ID_SpiralBezier.set(spiralbezierAPI.ID, spiralbezier)
							}
						)

						// init the arrays
						this.frontRepo.array_SpiralBezierGrids = []
						this.frontRepo.map_ID_SpiralBezierGrid.clear()

						spiralbeziergrids.forEach(
							spiralbeziergridAPI => {
								let spiralbeziergrid = new SpiralBezierGrid
								this.frontRepo.array_SpiralBezierGrids.push(spiralbeziergrid)
								this.frontRepo.map_ID_SpiralBezierGrid.set(spiralbeziergridAPI.ID, spiralbeziergrid)
							}
						)

						// init the arrays
						this.frontRepo.array_SpiralCircles = []
						this.frontRepo.map_ID_SpiralCircle.clear()

						spiralcircles.forEach(
							spiralcircleAPI => {
								let spiralcircle = new SpiralCircle
								this.frontRepo.array_SpiralCircles.push(spiralcircle)
								this.frontRepo.map_ID_SpiralCircle.set(spiralcircleAPI.ID, spiralcircle)
							}
						)

						// init the arrays
						this.frontRepo.array_SpiralCircleGrids = []
						this.frontRepo.map_ID_SpiralCircleGrid.clear()

						spiralcirclegrids.forEach(
							spiralcirclegridAPI => {
								let spiralcirclegrid = new SpiralCircleGrid
								this.frontRepo.array_SpiralCircleGrids.push(spiralcirclegrid)
								this.frontRepo.map_ID_SpiralCircleGrid.set(spiralcirclegridAPI.ID, spiralcirclegrid)
							}
						)

						// init the arrays
						this.frontRepo.array_SpiralLines = []
						this.frontRepo.map_ID_SpiralLine.clear()

						spirallines.forEach(
							spirallineAPI => {
								let spiralline = new SpiralLine
								this.frontRepo.array_SpiralLines.push(spiralline)
								this.frontRepo.map_ID_SpiralLine.set(spirallineAPI.ID, spiralline)
							}
						)

						// init the arrays
						this.frontRepo.array_SpiralLineGrids = []
						this.frontRepo.map_ID_SpiralLineGrid.clear()

						spirallinegrids.forEach(
							spirallinegridAPI => {
								let spirallinegrid = new SpiralLineGrid
								this.frontRepo.array_SpiralLineGrids.push(spirallinegrid)
								this.frontRepo.map_ID_SpiralLineGrid.set(spirallinegridAPI.ID, spirallinegrid)
							}
						)

						// init the arrays
						this.frontRepo.array_SpiralOrigins = []
						this.frontRepo.map_ID_SpiralOrigin.clear()

						spiralorigins.forEach(
							spiraloriginAPI => {
								let spiralorigin = new SpiralOrigin
								this.frontRepo.array_SpiralOrigins.push(spiralorigin)
								this.frontRepo.map_ID_SpiralOrigin.set(spiraloriginAPI.ID, spiralorigin)
							}
						)

						// init the arrays
						this.frontRepo.array_SpiralRhombuss = []
						this.frontRepo.map_ID_SpiralRhombus.clear()

						spiralrhombuss.forEach(
							spiralrhombusAPI => {
								let spiralrhombus = new SpiralRhombus
								this.frontRepo.array_SpiralRhombuss.push(spiralrhombus)
								this.frontRepo.map_ID_SpiralRhombus.set(spiralrhombusAPI.ID, spiralrhombus)
							}
						)

						// init the arrays
						this.frontRepo.array_SpiralRhombusGrids = []
						this.frontRepo.map_ID_SpiralRhombusGrid.clear()

						spiralrhombusgrids.forEach(
							spiralrhombusgridAPI => {
								let spiralrhombusgrid = new SpiralRhombusGrid
								this.frontRepo.array_SpiralRhombusGrids.push(spiralrhombusgrid)
								this.frontRepo.map_ID_SpiralRhombusGrid.set(spiralrhombusgridAPI.ID, spiralrhombusgrid)
							}
						)

						// init the arrays
						this.frontRepo.array_VerticalAxiss = []
						this.frontRepo.map_ID_VerticalAxis.clear()

						verticalaxiss.forEach(
							verticalaxisAPI => {
								let verticalaxis = new VerticalAxis
								this.frontRepo.array_VerticalAxiss.push(verticalaxis)
								this.frontRepo.map_ID_VerticalAxis.set(verticalaxisAPI.ID, verticalaxis)
							}
						)


						// 
						// Second Step: reddeem front objects
						// insertion point sub template for redeem 
						// fill up front objects
						axiss.forEach(
							axisAPI => {
								let axis = this.frontRepo.map_ID_Axis.get(axisAPI.ID)
								CopyAxisAPIToAxis(axisAPI, axis!, this.frontRepo)
							}
						)

						// fill up front objects
						axisgrids.forEach(
							axisgridAPI => {
								let axisgrid = this.frontRepo.map_ID_AxisGrid.get(axisgridAPI.ID)
								CopyAxisGridAPIToAxisGrid(axisgridAPI, axisgrid!, this.frontRepo)
							}
						)

						// fill up front objects
						beziers.forEach(
							bezierAPI => {
								let bezier = this.frontRepo.map_ID_Bezier.get(bezierAPI.ID)
								CopyBezierAPIToBezier(bezierAPI, bezier!, this.frontRepo)
							}
						)

						// fill up front objects
						beziergrids.forEach(
							beziergridAPI => {
								let beziergrid = this.frontRepo.map_ID_BezierGrid.get(beziergridAPI.ID)
								CopyBezierGridAPIToBezierGrid(beziergridAPI, beziergrid!, this.frontRepo)
							}
						)

						// fill up front objects
						beziergridstacks.forEach(
							beziergridstackAPI => {
								let beziergridstack = this.frontRepo.map_ID_BezierGridStack.get(beziergridstackAPI.ID)
								CopyBezierGridStackAPIToBezierGridStack(beziergridstackAPI, beziergridstack!, this.frontRepo)
							}
						)

						// fill up front objects
						circles.forEach(
							circleAPI => {
								let circle = this.frontRepo.map_ID_Circle.get(circleAPI.ID)
								CopyCircleAPIToCircle(circleAPI, circle!, this.frontRepo)
							}
						)

						// fill up front objects
						circlegrids.forEach(
							circlegridAPI => {
								let circlegrid = this.frontRepo.map_ID_CircleGrid.get(circlegridAPI.ID)
								CopyCircleGridAPIToCircleGrid(circlegridAPI, circlegrid!, this.frontRepo)
							}
						)

						// fill up front objects
						frontcurves.forEach(
							frontcurveAPI => {
								let frontcurve = this.frontRepo.map_ID_FrontCurve.get(frontcurveAPI.ID)
								CopyFrontCurveAPIToFrontCurve(frontcurveAPI, frontcurve!, this.frontRepo)
							}
						)

						// fill up front objects
						frontcurvestacks.forEach(
							frontcurvestackAPI => {
								let frontcurvestack = this.frontRepo.map_ID_FrontCurveStack.get(frontcurvestackAPI.ID)
								CopyFrontCurveStackAPIToFrontCurveStack(frontcurvestackAPI, frontcurvestack!, this.frontRepo)
							}
						)

						// fill up front objects
						horizontalaxiss.forEach(
							horizontalaxisAPI => {
								let horizontalaxis = this.frontRepo.map_ID_HorizontalAxis.get(horizontalaxisAPI.ID)
								CopyHorizontalAxisAPIToHorizontalAxis(horizontalaxisAPI, horizontalaxis!, this.frontRepo)
							}
						)

						// fill up front objects
						keys.forEach(
							keyAPI => {
								let key = this.frontRepo.map_ID_Key.get(keyAPI.ID)
								CopyKeyAPIToKey(keyAPI, key!, this.frontRepo)
							}
						)

						// fill up front objects
						noteinfos.forEach(
							noteinfoAPI => {
								let noteinfo = this.frontRepo.map_ID_NoteInfo.get(noteinfoAPI.ID)
								CopyNoteInfoAPIToNoteInfo(noteinfoAPI, noteinfo!, this.frontRepo)
							}
						)

						// fill up front objects
						parameters.forEach(
							parameterAPI => {
								let parameter = this.frontRepo.map_ID_Parameter.get(parameterAPI.ID)
								CopyParameterAPIToParameter(parameterAPI, parameter!, this.frontRepo)
							}
						)

						// fill up front objects
						rhombuss.forEach(
							rhombusAPI => {
								let rhombus = this.frontRepo.map_ID_Rhombus.get(rhombusAPI.ID)
								CopyRhombusAPIToRhombus(rhombusAPI, rhombus!, this.frontRepo)
							}
						)

						// fill up front objects
						rhombusgrids.forEach(
							rhombusgridAPI => {
								let rhombusgrid = this.frontRepo.map_ID_RhombusGrid.get(rhombusgridAPI.ID)
								CopyRhombusGridAPIToRhombusGrid(rhombusgridAPI, rhombusgrid!, this.frontRepo)
							}
						)

						// fill up front objects
						shapecategorys.forEach(
							shapecategoryAPI => {
								let shapecategory = this.frontRepo.map_ID_ShapeCategory.get(shapecategoryAPI.ID)
								CopyShapeCategoryAPIToShapeCategory(shapecategoryAPI, shapecategory!, this.frontRepo)
							}
						)

						// fill up front objects
						spiralbeziers.forEach(
							spiralbezierAPI => {
								let spiralbezier = this.frontRepo.map_ID_SpiralBezier.get(spiralbezierAPI.ID)
								CopySpiralBezierAPIToSpiralBezier(spiralbezierAPI, spiralbezier!, this.frontRepo)
							}
						)

						// fill up front objects
						spiralbeziergrids.forEach(
							spiralbeziergridAPI => {
								let spiralbeziergrid = this.frontRepo.map_ID_SpiralBezierGrid.get(spiralbeziergridAPI.ID)
								CopySpiralBezierGridAPIToSpiralBezierGrid(spiralbeziergridAPI, spiralbeziergrid!, this.frontRepo)
							}
						)

						// fill up front objects
						spiralcircles.forEach(
							spiralcircleAPI => {
								let spiralcircle = this.frontRepo.map_ID_SpiralCircle.get(spiralcircleAPI.ID)
								CopySpiralCircleAPIToSpiralCircle(spiralcircleAPI, spiralcircle!, this.frontRepo)
							}
						)

						// fill up front objects
						spiralcirclegrids.forEach(
							spiralcirclegridAPI => {
								let spiralcirclegrid = this.frontRepo.map_ID_SpiralCircleGrid.get(spiralcirclegridAPI.ID)
								CopySpiralCircleGridAPIToSpiralCircleGrid(spiralcirclegridAPI, spiralcirclegrid!, this.frontRepo)
							}
						)

						// fill up front objects
						spirallines.forEach(
							spirallineAPI => {
								let spiralline = this.frontRepo.map_ID_SpiralLine.get(spirallineAPI.ID)
								CopySpiralLineAPIToSpiralLine(spirallineAPI, spiralline!, this.frontRepo)
							}
						)

						// fill up front objects
						spirallinegrids.forEach(
							spirallinegridAPI => {
								let spirallinegrid = this.frontRepo.map_ID_SpiralLineGrid.get(spirallinegridAPI.ID)
								CopySpiralLineGridAPIToSpiralLineGrid(spirallinegridAPI, spirallinegrid!, this.frontRepo)
							}
						)

						// fill up front objects
						spiralorigins.forEach(
							spiraloriginAPI => {
								let spiralorigin = this.frontRepo.map_ID_SpiralOrigin.get(spiraloriginAPI.ID)
								CopySpiralOriginAPIToSpiralOrigin(spiraloriginAPI, spiralorigin!, this.frontRepo)
							}
						)

						// fill up front objects
						spiralrhombuss.forEach(
							spiralrhombusAPI => {
								let spiralrhombus = this.frontRepo.map_ID_SpiralRhombus.get(spiralrhombusAPI.ID)
								CopySpiralRhombusAPIToSpiralRhombus(spiralrhombusAPI, spiralrhombus!, this.frontRepo)
							}
						)

						// fill up front objects
						spiralrhombusgrids.forEach(
							spiralrhombusgridAPI => {
								let spiralrhombusgrid = this.frontRepo.map_ID_SpiralRhombusGrid.get(spiralrhombusgridAPI.ID)
								CopySpiralRhombusGridAPIToSpiralRhombusGrid(spiralrhombusgridAPI, spiralrhombusgrid!, this.frontRepo)
							}
						)

						// fill up front objects
						verticalaxiss.forEach(
							verticalaxisAPI => {
								let verticalaxis = this.frontRepo.map_ID_VerticalAxis.get(verticalaxisAPI.ID)
								CopyVerticalAxisAPIToVerticalAxis(verticalaxisAPI, verticalaxis!, this.frontRepo)
							}
						)


						// hand over control flow to observer
						observer.next(this.frontRepo)
					}
				)
			}
		)
	}

	public connectToWebSocket(GONG__StackPath: string): Observable<FrontRepo> {

		this.GONG__StackPath = GONG__StackPath


		let params = new HttpParams().set("GONG__StackPath", this.GONG__StackPath)
		let basePath = 'ws://localhost:8080/api/github.com/thomaspeugeot/phylotaxymusic/go/v1/ws/stage'
		let paramString = params.toString()
		let url = `${basePath}?${paramString}`
		this.socket = new WebSocket(url)

		return new Observable(observer => {
			this.socket!.onmessage = event => {


				const backRepoData = new BackRepoData(JSON.parse(event.data))

				let frontRepo = new (FrontRepo)

				// 
				// First Step: init map of instances
				// insertion point sub template for init 
				// init the arrays
				// insertion point sub template for init 
				// init the arrays
				frontRepo.array_Axiss = []
				frontRepo.map_ID_Axis.clear()

				backRepoData.AxisAPIs.forEach(
					axisAPI => {
						let axis = new Axis
						frontRepo.array_Axiss.push(axis)
						frontRepo.map_ID_Axis.set(axisAPI.ID, axis)
					}
				)

				// init the arrays
				frontRepo.array_AxisGrids = []
				frontRepo.map_ID_AxisGrid.clear()

				backRepoData.AxisGridAPIs.forEach(
					axisgridAPI => {
						let axisgrid = new AxisGrid
						frontRepo.array_AxisGrids.push(axisgrid)
						frontRepo.map_ID_AxisGrid.set(axisgridAPI.ID, axisgrid)
					}
				)

				// init the arrays
				frontRepo.array_Beziers = []
				frontRepo.map_ID_Bezier.clear()

				backRepoData.BezierAPIs.forEach(
					bezierAPI => {
						let bezier = new Bezier
						frontRepo.array_Beziers.push(bezier)
						frontRepo.map_ID_Bezier.set(bezierAPI.ID, bezier)
					}
				)

				// init the arrays
				frontRepo.array_BezierGrids = []
				frontRepo.map_ID_BezierGrid.clear()

				backRepoData.BezierGridAPIs.forEach(
					beziergridAPI => {
						let beziergrid = new BezierGrid
						frontRepo.array_BezierGrids.push(beziergrid)
						frontRepo.map_ID_BezierGrid.set(beziergridAPI.ID, beziergrid)
					}
				)

				// init the arrays
				frontRepo.array_BezierGridStacks = []
				frontRepo.map_ID_BezierGridStack.clear()

				backRepoData.BezierGridStackAPIs.forEach(
					beziergridstackAPI => {
						let beziergridstack = new BezierGridStack
						frontRepo.array_BezierGridStacks.push(beziergridstack)
						frontRepo.map_ID_BezierGridStack.set(beziergridstackAPI.ID, beziergridstack)
					}
				)

				// init the arrays
				frontRepo.array_Circles = []
				frontRepo.map_ID_Circle.clear()

				backRepoData.CircleAPIs.forEach(
					circleAPI => {
						let circle = new Circle
						frontRepo.array_Circles.push(circle)
						frontRepo.map_ID_Circle.set(circleAPI.ID, circle)
					}
				)

				// init the arrays
				frontRepo.array_CircleGrids = []
				frontRepo.map_ID_CircleGrid.clear()

				backRepoData.CircleGridAPIs.forEach(
					circlegridAPI => {
						let circlegrid = new CircleGrid
						frontRepo.array_CircleGrids.push(circlegrid)
						frontRepo.map_ID_CircleGrid.set(circlegridAPI.ID, circlegrid)
					}
				)

				// init the arrays
				frontRepo.array_FrontCurves = []
				frontRepo.map_ID_FrontCurve.clear()

				backRepoData.FrontCurveAPIs.forEach(
					frontcurveAPI => {
						let frontcurve = new FrontCurve
						frontRepo.array_FrontCurves.push(frontcurve)
						frontRepo.map_ID_FrontCurve.set(frontcurveAPI.ID, frontcurve)
					}
				)

				// init the arrays
				frontRepo.array_FrontCurveStacks = []
				frontRepo.map_ID_FrontCurveStack.clear()

				backRepoData.FrontCurveStackAPIs.forEach(
					frontcurvestackAPI => {
						let frontcurvestack = new FrontCurveStack
						frontRepo.array_FrontCurveStacks.push(frontcurvestack)
						frontRepo.map_ID_FrontCurveStack.set(frontcurvestackAPI.ID, frontcurvestack)
					}
				)

				// init the arrays
				frontRepo.array_HorizontalAxiss = []
				frontRepo.map_ID_HorizontalAxis.clear()

				backRepoData.HorizontalAxisAPIs.forEach(
					horizontalaxisAPI => {
						let horizontalaxis = new HorizontalAxis
						frontRepo.array_HorizontalAxiss.push(horizontalaxis)
						frontRepo.map_ID_HorizontalAxis.set(horizontalaxisAPI.ID, horizontalaxis)
					}
				)

				// init the arrays
				frontRepo.array_Keys = []
				frontRepo.map_ID_Key.clear()

				backRepoData.KeyAPIs.forEach(
					keyAPI => {
						let key = new Key
						frontRepo.array_Keys.push(key)
						frontRepo.map_ID_Key.set(keyAPI.ID, key)
					}
				)

				// init the arrays
				frontRepo.array_NoteInfos = []
				frontRepo.map_ID_NoteInfo.clear()

				backRepoData.NoteInfoAPIs.forEach(
					noteinfoAPI => {
						let noteinfo = new NoteInfo
						frontRepo.array_NoteInfos.push(noteinfo)
						frontRepo.map_ID_NoteInfo.set(noteinfoAPI.ID, noteinfo)
					}
				)

				// init the arrays
				frontRepo.array_Parameters = []
				frontRepo.map_ID_Parameter.clear()

				backRepoData.ParameterAPIs.forEach(
					parameterAPI => {
						let parameter = new Parameter
						frontRepo.array_Parameters.push(parameter)
						frontRepo.map_ID_Parameter.set(parameterAPI.ID, parameter)
					}
				)

				// init the arrays
				frontRepo.array_Rhombuss = []
				frontRepo.map_ID_Rhombus.clear()

				backRepoData.RhombusAPIs.forEach(
					rhombusAPI => {
						let rhombus = new Rhombus
						frontRepo.array_Rhombuss.push(rhombus)
						frontRepo.map_ID_Rhombus.set(rhombusAPI.ID, rhombus)
					}
				)

				// init the arrays
				frontRepo.array_RhombusGrids = []
				frontRepo.map_ID_RhombusGrid.clear()

				backRepoData.RhombusGridAPIs.forEach(
					rhombusgridAPI => {
						let rhombusgrid = new RhombusGrid
						frontRepo.array_RhombusGrids.push(rhombusgrid)
						frontRepo.map_ID_RhombusGrid.set(rhombusgridAPI.ID, rhombusgrid)
					}
				)

				// init the arrays
				frontRepo.array_ShapeCategorys = []
				frontRepo.map_ID_ShapeCategory.clear()

				backRepoData.ShapeCategoryAPIs.forEach(
					shapecategoryAPI => {
						let shapecategory = new ShapeCategory
						frontRepo.array_ShapeCategorys.push(shapecategory)
						frontRepo.map_ID_ShapeCategory.set(shapecategoryAPI.ID, shapecategory)
					}
				)

				// init the arrays
				frontRepo.array_SpiralBeziers = []
				frontRepo.map_ID_SpiralBezier.clear()

				backRepoData.SpiralBezierAPIs.forEach(
					spiralbezierAPI => {
						let spiralbezier = new SpiralBezier
						frontRepo.array_SpiralBeziers.push(spiralbezier)
						frontRepo.map_ID_SpiralBezier.set(spiralbezierAPI.ID, spiralbezier)
					}
				)

				// init the arrays
				frontRepo.array_SpiralBezierGrids = []
				frontRepo.map_ID_SpiralBezierGrid.clear()

				backRepoData.SpiralBezierGridAPIs.forEach(
					spiralbeziergridAPI => {
						let spiralbeziergrid = new SpiralBezierGrid
						frontRepo.array_SpiralBezierGrids.push(spiralbeziergrid)
						frontRepo.map_ID_SpiralBezierGrid.set(spiralbeziergridAPI.ID, spiralbeziergrid)
					}
				)

				// init the arrays
				frontRepo.array_SpiralCircles = []
				frontRepo.map_ID_SpiralCircle.clear()

				backRepoData.SpiralCircleAPIs.forEach(
					spiralcircleAPI => {
						let spiralcircle = new SpiralCircle
						frontRepo.array_SpiralCircles.push(spiralcircle)
						frontRepo.map_ID_SpiralCircle.set(spiralcircleAPI.ID, spiralcircle)
					}
				)

				// init the arrays
				frontRepo.array_SpiralCircleGrids = []
				frontRepo.map_ID_SpiralCircleGrid.clear()

				backRepoData.SpiralCircleGridAPIs.forEach(
					spiralcirclegridAPI => {
						let spiralcirclegrid = new SpiralCircleGrid
						frontRepo.array_SpiralCircleGrids.push(spiralcirclegrid)
						frontRepo.map_ID_SpiralCircleGrid.set(spiralcirclegridAPI.ID, spiralcirclegrid)
					}
				)

				// init the arrays
				frontRepo.array_SpiralLines = []
				frontRepo.map_ID_SpiralLine.clear()

				backRepoData.SpiralLineAPIs.forEach(
					spirallineAPI => {
						let spiralline = new SpiralLine
						frontRepo.array_SpiralLines.push(spiralline)
						frontRepo.map_ID_SpiralLine.set(spirallineAPI.ID, spiralline)
					}
				)

				// init the arrays
				frontRepo.array_SpiralLineGrids = []
				frontRepo.map_ID_SpiralLineGrid.clear()

				backRepoData.SpiralLineGridAPIs.forEach(
					spirallinegridAPI => {
						let spirallinegrid = new SpiralLineGrid
						frontRepo.array_SpiralLineGrids.push(spirallinegrid)
						frontRepo.map_ID_SpiralLineGrid.set(spirallinegridAPI.ID, spirallinegrid)
					}
				)

				// init the arrays
				frontRepo.array_SpiralOrigins = []
				frontRepo.map_ID_SpiralOrigin.clear()

				backRepoData.SpiralOriginAPIs.forEach(
					spiraloriginAPI => {
						let spiralorigin = new SpiralOrigin
						frontRepo.array_SpiralOrigins.push(spiralorigin)
						frontRepo.map_ID_SpiralOrigin.set(spiraloriginAPI.ID, spiralorigin)
					}
				)

				// init the arrays
				frontRepo.array_SpiralRhombuss = []
				frontRepo.map_ID_SpiralRhombus.clear()

				backRepoData.SpiralRhombusAPIs.forEach(
					spiralrhombusAPI => {
						let spiralrhombus = new SpiralRhombus
						frontRepo.array_SpiralRhombuss.push(spiralrhombus)
						frontRepo.map_ID_SpiralRhombus.set(spiralrhombusAPI.ID, spiralrhombus)
					}
				)

				// init the arrays
				frontRepo.array_SpiralRhombusGrids = []
				frontRepo.map_ID_SpiralRhombusGrid.clear()

				backRepoData.SpiralRhombusGridAPIs.forEach(
					spiralrhombusgridAPI => {
						let spiralrhombusgrid = new SpiralRhombusGrid
						frontRepo.array_SpiralRhombusGrids.push(spiralrhombusgrid)
						frontRepo.map_ID_SpiralRhombusGrid.set(spiralrhombusgridAPI.ID, spiralrhombusgrid)
					}
				)

				// init the arrays
				frontRepo.array_VerticalAxiss = []
				frontRepo.map_ID_VerticalAxis.clear()

				backRepoData.VerticalAxisAPIs.forEach(
					verticalaxisAPI => {
						let verticalaxis = new VerticalAxis
						frontRepo.array_VerticalAxiss.push(verticalaxis)
						frontRepo.map_ID_VerticalAxis.set(verticalaxisAPI.ID, verticalaxis)
					}
				)


				// 
				// Second Step: reddeem front objects
				// insertion point sub template for redeem 
				// fill up front objects
				// insertion point sub template for redeem 
				// fill up front objects
				backRepoData.AxisAPIs.forEach(
					axisAPI => {
						let axis = frontRepo.map_ID_Axis.get(axisAPI.ID)
						CopyAxisAPIToAxis(axisAPI, axis!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.AxisGridAPIs.forEach(
					axisgridAPI => {
						let axisgrid = frontRepo.map_ID_AxisGrid.get(axisgridAPI.ID)
						CopyAxisGridAPIToAxisGrid(axisgridAPI, axisgrid!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.BezierAPIs.forEach(
					bezierAPI => {
						let bezier = frontRepo.map_ID_Bezier.get(bezierAPI.ID)
						CopyBezierAPIToBezier(bezierAPI, bezier!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.BezierGridAPIs.forEach(
					beziergridAPI => {
						let beziergrid = frontRepo.map_ID_BezierGrid.get(beziergridAPI.ID)
						CopyBezierGridAPIToBezierGrid(beziergridAPI, beziergrid!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.BezierGridStackAPIs.forEach(
					beziergridstackAPI => {
						let beziergridstack = frontRepo.map_ID_BezierGridStack.get(beziergridstackAPI.ID)
						CopyBezierGridStackAPIToBezierGridStack(beziergridstackAPI, beziergridstack!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.CircleAPIs.forEach(
					circleAPI => {
						let circle = frontRepo.map_ID_Circle.get(circleAPI.ID)
						CopyCircleAPIToCircle(circleAPI, circle!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.CircleGridAPIs.forEach(
					circlegridAPI => {
						let circlegrid = frontRepo.map_ID_CircleGrid.get(circlegridAPI.ID)
						CopyCircleGridAPIToCircleGrid(circlegridAPI, circlegrid!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.FrontCurveAPIs.forEach(
					frontcurveAPI => {
						let frontcurve = frontRepo.map_ID_FrontCurve.get(frontcurveAPI.ID)
						CopyFrontCurveAPIToFrontCurve(frontcurveAPI, frontcurve!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.FrontCurveStackAPIs.forEach(
					frontcurvestackAPI => {
						let frontcurvestack = frontRepo.map_ID_FrontCurveStack.get(frontcurvestackAPI.ID)
						CopyFrontCurveStackAPIToFrontCurveStack(frontcurvestackAPI, frontcurvestack!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.HorizontalAxisAPIs.forEach(
					horizontalaxisAPI => {
						let horizontalaxis = frontRepo.map_ID_HorizontalAxis.get(horizontalaxisAPI.ID)
						CopyHorizontalAxisAPIToHorizontalAxis(horizontalaxisAPI, horizontalaxis!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.KeyAPIs.forEach(
					keyAPI => {
						let key = frontRepo.map_ID_Key.get(keyAPI.ID)
						CopyKeyAPIToKey(keyAPI, key!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.NoteInfoAPIs.forEach(
					noteinfoAPI => {
						let noteinfo = frontRepo.map_ID_NoteInfo.get(noteinfoAPI.ID)
						CopyNoteInfoAPIToNoteInfo(noteinfoAPI, noteinfo!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.ParameterAPIs.forEach(
					parameterAPI => {
						let parameter = frontRepo.map_ID_Parameter.get(parameterAPI.ID)
						CopyParameterAPIToParameter(parameterAPI, parameter!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.RhombusAPIs.forEach(
					rhombusAPI => {
						let rhombus = frontRepo.map_ID_Rhombus.get(rhombusAPI.ID)
						CopyRhombusAPIToRhombus(rhombusAPI, rhombus!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.RhombusGridAPIs.forEach(
					rhombusgridAPI => {
						let rhombusgrid = frontRepo.map_ID_RhombusGrid.get(rhombusgridAPI.ID)
						CopyRhombusGridAPIToRhombusGrid(rhombusgridAPI, rhombusgrid!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.ShapeCategoryAPIs.forEach(
					shapecategoryAPI => {
						let shapecategory = frontRepo.map_ID_ShapeCategory.get(shapecategoryAPI.ID)
						CopyShapeCategoryAPIToShapeCategory(shapecategoryAPI, shapecategory!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.SpiralBezierAPIs.forEach(
					spiralbezierAPI => {
						let spiralbezier = frontRepo.map_ID_SpiralBezier.get(spiralbezierAPI.ID)
						CopySpiralBezierAPIToSpiralBezier(spiralbezierAPI, spiralbezier!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.SpiralBezierGridAPIs.forEach(
					spiralbeziergridAPI => {
						let spiralbeziergrid = frontRepo.map_ID_SpiralBezierGrid.get(spiralbeziergridAPI.ID)
						CopySpiralBezierGridAPIToSpiralBezierGrid(spiralbeziergridAPI, spiralbeziergrid!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.SpiralCircleAPIs.forEach(
					spiralcircleAPI => {
						let spiralcircle = frontRepo.map_ID_SpiralCircle.get(spiralcircleAPI.ID)
						CopySpiralCircleAPIToSpiralCircle(spiralcircleAPI, spiralcircle!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.SpiralCircleGridAPIs.forEach(
					spiralcirclegridAPI => {
						let spiralcirclegrid = frontRepo.map_ID_SpiralCircleGrid.get(spiralcirclegridAPI.ID)
						CopySpiralCircleGridAPIToSpiralCircleGrid(spiralcirclegridAPI, spiralcirclegrid!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.SpiralLineAPIs.forEach(
					spirallineAPI => {
						let spiralline = frontRepo.map_ID_SpiralLine.get(spirallineAPI.ID)
						CopySpiralLineAPIToSpiralLine(spirallineAPI, spiralline!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.SpiralLineGridAPIs.forEach(
					spirallinegridAPI => {
						let spirallinegrid = frontRepo.map_ID_SpiralLineGrid.get(spirallinegridAPI.ID)
						CopySpiralLineGridAPIToSpiralLineGrid(spirallinegridAPI, spirallinegrid!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.SpiralOriginAPIs.forEach(
					spiraloriginAPI => {
						let spiralorigin = frontRepo.map_ID_SpiralOrigin.get(spiraloriginAPI.ID)
						CopySpiralOriginAPIToSpiralOrigin(spiraloriginAPI, spiralorigin!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.SpiralRhombusAPIs.forEach(
					spiralrhombusAPI => {
						let spiralrhombus = frontRepo.map_ID_SpiralRhombus.get(spiralrhombusAPI.ID)
						CopySpiralRhombusAPIToSpiralRhombus(spiralrhombusAPI, spiralrhombus!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.SpiralRhombusGridAPIs.forEach(
					spiralrhombusgridAPI => {
						let spiralrhombusgrid = frontRepo.map_ID_SpiralRhombusGrid.get(spiralrhombusgridAPI.ID)
						CopySpiralRhombusGridAPIToSpiralRhombusGrid(spiralrhombusgridAPI, spiralrhombusgrid!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.VerticalAxisAPIs.forEach(
					verticalaxisAPI => {
						let verticalaxis = frontRepo.map_ID_VerticalAxis.get(verticalaxisAPI.ID)
						CopyVerticalAxisAPIToVerticalAxis(verticalaxisAPI, verticalaxis!, frontRepo)
					}
				)



				observer.next(frontRepo)
			}
			this.socket!.onerror = event => {
				observer.error(event)
			}
			this.socket!.onclose = event => {
				observer.complete()
			}

			return () => {
				this.socket!.close()
			}
		})
	}
}

// insertion point for get unique ID per struct 
export function getAxisUniqueID(id: number): number {
	return 31 * id
}
export function getAxisGridUniqueID(id: number): number {
	return 37 * id
}
export function getBezierUniqueID(id: number): number {
	return 41 * id
}
export function getBezierGridUniqueID(id: number): number {
	return 43 * id
}
export function getBezierGridStackUniqueID(id: number): number {
	return 47 * id
}
export function getCircleUniqueID(id: number): number {
	return 53 * id
}
export function getCircleGridUniqueID(id: number): number {
	return 59 * id
}
export function getFrontCurveUniqueID(id: number): number {
	return 61 * id
}
export function getFrontCurveStackUniqueID(id: number): number {
	return 67 * id
}
export function getHorizontalAxisUniqueID(id: number): number {
	return 71 * id
}
export function getKeyUniqueID(id: number): number {
	return 73 * id
}
export function getNoteInfoUniqueID(id: number): number {
	return 79 * id
}
export function getParameterUniqueID(id: number): number {
	return 83 * id
}
export function getRhombusUniqueID(id: number): number {
	return 89 * id
}
export function getRhombusGridUniqueID(id: number): number {
	return 97 * id
}
export function getShapeCategoryUniqueID(id: number): number {
	return 101 * id
}
export function getSpiralBezierUniqueID(id: number): number {
	return 103 * id
}
export function getSpiralBezierGridUniqueID(id: number): number {
	return 107 * id
}
export function getSpiralCircleUniqueID(id: number): number {
	return 109 * id
}
export function getSpiralCircleGridUniqueID(id: number): number {
	return 113 * id
}
export function getSpiralLineUniqueID(id: number): number {
	return 127 * id
}
export function getSpiralLineGridUniqueID(id: number): number {
	return 131 * id
}
export function getSpiralOriginUniqueID(id: number): number {
	return 137 * id
}
export function getSpiralRhombusUniqueID(id: number): number {
	return 139 * id
}
export function getSpiralRhombusGridUniqueID(id: number): number {
	return 149 * id
}
export function getVerticalAxisUniqueID(id: number): number {
	return 151 * id
}
