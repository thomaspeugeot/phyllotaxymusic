// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	FreqencyAPIs []*FreqencyAPI

	NoteAPIs []*NoteAPI
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {

	// wait till backRepo is written by commit
	backRepo.rwMutex.RLock()
	defer backRepo.rwMutex.RUnlock()

	// insertion point for slices copies
	for _, freqencyDB := range backRepo.BackRepoFreqency.Map_FreqencyDBID_FreqencyDB {

		var freqencyAPI FreqencyAPI
		freqencyAPI.ID = freqencyDB.ID
		freqencyAPI.FreqencyPointersEncoding = freqencyDB.FreqencyPointersEncoding
		freqencyDB.CopyBasicFieldsToFreqency_WOP(&freqencyAPI.Freqency_WOP)

		backRepoData.FreqencyAPIs = append(backRepoData.FreqencyAPIs, &freqencyAPI)
	}

	for _, noteDB := range backRepo.BackRepoNote.Map_NoteDBID_NoteDB {

		var noteAPI NoteAPI
		noteAPI.ID = noteDB.ID
		noteAPI.NotePointersEncoding = noteDB.NotePointersEncoding
		noteDB.CopyBasicFieldsToNote_WOP(&noteAPI.Note_WOP)

		backRepoData.NoteAPIs = append(backRepoData.NoteAPIs, &noteAPI)
	}

}
