// generated code - do not edit

//insertion point for imports
import { FreqencyAPI } from './freqency-api'

import { NoteAPI } from './note-api'


export class BackRepoData {
	// insertion point for declarations
	FreqencyAPIs = new Array<FreqencyAPI>()

	NoteAPIs = new Array<NoteAPI>()



	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.FreqencyAPIs = data?.FreqencyAPIs || [];

		this.NoteAPIs = data?.NoteAPIs || [];

	}

}