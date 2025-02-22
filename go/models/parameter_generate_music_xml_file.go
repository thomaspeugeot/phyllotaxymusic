package models

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	m "github.com/thomaspeugeot/phyllotaxymusic/go/musicxml"
)

func (parameter *Parameter) GenerateMusicXMLFile() bool {
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

	// Write the XML to a new file
	filename := generateTimestampedFilename("export", ".musicxml")
	err = os.WriteFile(
		filepath.Join("musicxml", filename),
		[]byte(xml.Header+string(output)), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return true
	}
	return false
}

// generateTimestampedFilename creates a filename using the current timestamp (YYYYMMDD_HHMMSS).
func generateTimestampedFilename(prefix, extension string) string {
	// Format current date and time up to the second
	timestamp := time.Now().Format("20060102_150405")
	// Construct the filename
	return fmt.Sprintf("%s_%s%s", prefix, timestamp, extension)
}
