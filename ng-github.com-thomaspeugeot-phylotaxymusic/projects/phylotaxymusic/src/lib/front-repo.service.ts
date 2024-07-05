import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'

// insertion point sub template for services imports
import { LineAPI } from './line-api'
import { Line, CopyLineAPIToLine } from './line'
import { LineService } from './line.service'

import { ParameterAPI } from './parameter-api'
import { Parameter, CopyParameterAPIToParameter } from './parameter'
import { ParameterService } from './parameter.service'

import { RhombusAPI } from './rhombus-api'
import { Rhombus, CopyRhombusAPIToRhombus } from './rhombus'
import { RhombusService } from './rhombus.service'


import { BackRepoData } from './back-repo-data'

export const StackType = "github.com/thomaspeugeot/phylotaxymusic/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
	array_Lines = new Array<Line>() // array of front instances
	map_ID_Line = new Map<number, Line>() // map of front instances

	array_Parameters = new Array<Parameter>() // array of front instances
	map_ID_Parameter = new Map<number, Parameter>() // map of front instances

	array_Rhombuss = new Array<Rhombus>() // array of front instances
	map_ID_Rhombus = new Map<number, Rhombus>() // map of front instances


	// getFrontArray allows for a get function that is robust to refactoring of the named struct name
	// for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
	// contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
	getFrontArray<Type>(gongStructName: string): Array<Type> {
		switch (gongStructName) {
			// insertion point
			case 'Line':
				return this.array_Lines as unknown as Array<Type>
			case 'Parameter':
				return this.array_Parameters as unknown as Array<Type>
			case 'Rhombus':
				return this.array_Rhombuss as unknown as Array<Type>
			default:
				throw new Error("Type not recognized");
		}
	}

	getFrontMap<Type>(gongStructName: string): Map<number, Type> {
		switch (gongStructName) {
			// insertion point
			case 'Line':
				return this.map_ID_Line as unknown as Map<number, Type>
			case 'Parameter':
				return this.map_ID_Parameter as unknown as Map<number, Type>
			case 'Rhombus':
				return this.map_ID_Rhombus as unknown as Map<number, Type>
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
		private lineService: LineService,
		private parameterService: ParameterService,
		private rhombusService: RhombusService,
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
		Observable<LineAPI[]>,
		Observable<ParameterAPI[]>,
		Observable<RhombusAPI[]>,
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
			this.lineService.getLines(this.GONG__StackPath, this.frontRepo),
			this.parameterService.getParameters(this.GONG__StackPath, this.frontRepo),
			this.rhombusService.getRhombuss(this.GONG__StackPath, this.frontRepo),
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
			this.lineService.getLines(this.GONG__StackPath, this.frontRepo),
			this.parameterService.getParameters(this.GONG__StackPath, this.frontRepo),
			this.rhombusService.getRhombuss(this.GONG__StackPath, this.frontRepo),
		]

		return new Observable<FrontRepo>(
			(observer) => {
				combineLatest(
					this.observableFrontRepo
				).subscribe(
					([
						___of_null, // see above for the explanation about of
						// insertion point sub template for declarations 
						lines_,
						parameters_,
						rhombuss_,
					]) => {
						let _this = this
						// Typing can be messy with many items. Therefore, type casting is necessary here
						// insertion point sub template for type casting 
						var lines: LineAPI[]
						lines = lines_ as LineAPI[]
						var parameters: ParameterAPI[]
						parameters = parameters_ as ParameterAPI[]
						var rhombuss: RhombusAPI[]
						rhombuss = rhombuss_ as RhombusAPI[]

						// 
						// First Step: init map of instances
						// insertion point sub template for init 
						// init the arrays
						this.frontRepo.array_Lines = []
						this.frontRepo.map_ID_Line.clear()

						lines.forEach(
							lineAPI => {
								let line = new Line
								this.frontRepo.array_Lines.push(line)
								this.frontRepo.map_ID_Line.set(lineAPI.ID, line)
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


						// 
						// Second Step: reddeem front objects
						// insertion point sub template for redeem 
						// fill up front objects
						lines.forEach(
							lineAPI => {
								let line = this.frontRepo.map_ID_Line.get(lineAPI.ID)
								CopyLineAPIToLine(lineAPI, line!, this.frontRepo)
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
				let _this = this

				const backRepoData = new BackRepoData(JSON.parse(event.data))

				// 
				// First Step: init map of instances
				// insertion point sub template for init 
				// init the arrays
				// insertion point sub template for init 
				// init the arrays
				this.frontRepo.array_Lines = []
				this.frontRepo.map_ID_Line.clear()

				backRepoData.LineAPIs.forEach(
					lineAPI => {
						let line = new Line
						this.frontRepo.array_Lines.push(line)
						this.frontRepo.map_ID_Line.set(lineAPI.ID, line)
					}
				)

				// init the arrays
				this.frontRepo.array_Parameters = []
				this.frontRepo.map_ID_Parameter.clear()

				backRepoData.ParameterAPIs.forEach(
					parameterAPI => {
						let parameter = new Parameter
						this.frontRepo.array_Parameters.push(parameter)
						this.frontRepo.map_ID_Parameter.set(parameterAPI.ID, parameter)
					}
				)

				// init the arrays
				this.frontRepo.array_Rhombuss = []
				this.frontRepo.map_ID_Rhombus.clear()

				backRepoData.RhombusAPIs.forEach(
					rhombusAPI => {
						let rhombus = new Rhombus
						this.frontRepo.array_Rhombuss.push(rhombus)
						this.frontRepo.map_ID_Rhombus.set(rhombusAPI.ID, rhombus)
					}
				)


				// 
				// Second Step: reddeem front objects
				// insertion point sub template for redeem 
				// fill up front objects
				// insertion point sub template for redeem 
				// fill up front objects
				backRepoData.LineAPIs.forEach(
					lineAPI => {
						let line = this.frontRepo.map_ID_Line.get(lineAPI.ID)
						CopyLineAPIToLine(lineAPI, line!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.ParameterAPIs.forEach(
					parameterAPI => {
						let parameter = this.frontRepo.map_ID_Parameter.get(parameterAPI.ID)
						CopyParameterAPIToParameter(parameterAPI, parameter!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.RhombusAPIs.forEach(
					rhombusAPI => {
						let rhombus = this.frontRepo.map_ID_Rhombus.get(rhombusAPI.ID)
						CopyRhombusAPIToRhombus(rhombusAPI, rhombus!, this.frontRepo)
					}
				)



				observer.next(this.frontRepo)
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
export function getLineUniqueID(id: number): number {
	return 31 * id
}
export function getParameterUniqueID(id: number): number {
	return 37 * id
}
export function getRhombusUniqueID(id: number): number {
	return 41 * id
}
