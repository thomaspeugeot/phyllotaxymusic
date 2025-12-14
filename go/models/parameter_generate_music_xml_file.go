package models

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	load "github.com/fullstack-lang/gong/lib/load/go/models"
	m "github.com/thomaspeugeot/phyllotaxymusic/go/musicxml"
)

type GenerationMode string

const (
	DOWNLOAD        GenerationMode = "DOWNLOAD"
	STATIC_WEB_SITE GenerationMode = "STATIC_WEB_SITE"
)

func (parameter *Parameter) GenerateMusicXMLFile(generationMode GenerationMode) bool {
	log.Println("Export to Music xml requested")

	var scorePartwise m.Score_partwise

	parameter.generateHeaders(&scorePartwise)

	{
		var part m.A_part
		scorePartwise.Part = append(scorePartwise.Part, &part)

		part.Id = "P1"

		// the pitch computation is based on the geometry of the curve interpolated with the pitch grid
		// the interpolation is not perfect, therefore, one recompute the pitch of the
		// second voice
		shift := parameter.ActualBeatsTemporalShift
		_ = shift
		pitchDiff := parameter.PitchDifference
		_ = pitchDiff
		firstVoiceCircleNotes := parameter.FirstVoiceNotes.Circles
		_ = firstVoiceCircleNotes
		secondVoiceCircleNotes := parameter.SecondVoiceNotes.Circles
		_ = secondVoiceCircleNotes

		nbFullMesaures := 10
		var lastCirclePartSecondVoiceNote *Circle
		var lastFirstPartSecondVoiceNote *m.Note
		{
			lastCirclePartSecondVoiceNote,
				lastFirstPartSecondVoiceNote =
				parameter.addMeasure(&part, parameter.FirstVoiceNotes, parameter.SecondVoiceNotes, 0, nil, nil)
			for range nbFullMesaures {
				lastCirclePartSecondVoiceNote,
					lastFirstPartSecondVoiceNote =
					parameter.addMeasure(&part, parameter.FirstVoiceNotes, parameter.SecondVoiceNotes, 1,
						lastCirclePartSecondVoiceNote, lastFirstPartSecondVoiceNote)
			}
			parameter.addMeasure(&part, parameter.FirstVoiceNotes, parameter.SecondVoiceNotes, 2,
				lastCirclePartSecondVoiceNote, lastFirstPartSecondVoiceNote)
		}

	}

	// Marshal the books struct back into XML
	output, err := xml.MarshalIndent(scorePartwise, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling to XML:", err)
		return true
	}

	if generationMode == DOWNLOAD {

		// Write the XML to a new file
		filename := generateTimestampedFilename("export", ".musicxml")

		downloadStage := parameter.stager.loadStage
		downloadStage.Reset()
		fileToDownlad := (&load.FileToDownload{Name: filename}).Stage(downloadStage)
		fileToDownlad.Base64EncodedContent = base64.StdEncoding.EncodeToString(output)
		downloadStage.Commit()

	} else {

		pathToGeneratedScore := parameter.PathToGeneratedScore

		// --- Start: Remove existing pathToGeneratedScore directory ---
		// log.Printf("Attempting to remove existing directory: %s", pathToGeneratedScore)
		err := os.RemoveAll(pathToGeneratedScore)
		if err != nil {
			// Log the error but continue, as MkdirAll below might still succeed if the path didn't exist
			// or if the error was related to something already gone.
			// If MkdirAll fails later, that error will be caught.
			log.Printf("Warning: Error removing directory '%s': %v. Attempting to continue.", pathToGeneratedScore, err)
		} else {
			// log.Printf("Successfully removed existing directory: %s", pathToGeneratedScore)
		}
		// --- End: Remove existing pathToGeneratedScore directory ---

		// --- Start: Generate root _index.md for the Content ---
		// 1. Create the root content directory (MkdirAll handles existing directories gracefully)
		err = os.MkdirAll(pathToGeneratedScore, 0755) // Use 0755 for standard directory permissions
		if err != nil {
			log.Printf("Error creating root content directory '%s': %v\n", pathToGeneratedScore, err)
			// Decide if this is fatal or if chapter generation should still proceed.
			// For now, let's return if the root directory cannot be created.
		}

		filepath := filepath.Join(parameter.PathToGeneratedScore, MusicXMLFile)

		writeErr := os.WriteFile(filepath, []byte(output), 0644)
		if writeErr != nil {
			log.Printf("Error writing score file '%s': %v\n", filepath, writeErr)
			// Return the write error. The defer will handle cleanup.
		}
		log.Printf("Music file created successfully: %s\n", filepath)
	}

	// err = os.WriteFile(
	// 	filepath.Join("musicxml", filename),
	// 	[]byte(xml.Header+string(output)), 0644)
	// if err != nil {
	// 	fmt.Println("Error writing to file:", err)
	// 	return true
	// }
	return false
}

// generateTimestampedFilename creates a filename using the current timestamp (YYYYMMDD_HHMMSS).
func generateTimestampedFilename(prefix, extension string) string {
	// Format current date and time up to the second
	timestamp := time.Now().Format("20060102_150405")
	// Construct the filename
	return fmt.Sprintf("%s_%s%s", prefix, timestamp, extension)
}
