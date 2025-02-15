// generated code - do not edit

//insertion point for imports
import { CheckboxAPI } from './checkbox-api'

import { GroupAPI } from './group-api'

import { SliderAPI } from './slider-api'


export class BackRepoData {
	// insertion point for declarations
	CheckboxAPIs = new Array<CheckboxAPI>()

	GroupAPIs = new Array<GroupAPI>()

	SliderAPIs = new Array<SliderAPI>()



	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.CheckboxAPIs = data?.CheckboxAPIs || [];

		this.GroupAPIs = data?.GroupAPIs || [];

		this.SliderAPIs = data?.SliderAPIs || [];

	}

}