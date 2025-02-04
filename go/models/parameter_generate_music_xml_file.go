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

		{
			parameter.addMeasure(&part)
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

func (parameter *Parameter) addMeasure(part *m.A_part) {
	var measure m.A_measure
	part.Measure = append(part.Measure, &measure)

	measure.Number = "1"
	measure.Width = "269.19"

	{
		firstVoiceNotes := parameter.FirstVoiceNotes
		notes := firstVoiceNotes.Circles
		for _, noteIn := range notes {
			log.Println(noteIn.Pitch)

			if !noteIn.isKept {
				continue
			}

			var note m.Note
			measure.Note = append(measure.Note, &note)

			pitch := generateNote(noteIn.Pitch)
			note.Pitch = pitch

			note.Default_x = "82.98"
			note.Default_y = "-25"

			note.Duration = "1"
			note.Voice = "1"

			var type_ m.Note_type
			note.Type = &type_

			type_.EnclosedText = "eighth"

			var stem m.Stem
			note.Stem = &stem
			stem.EnclosedText = "down"

			var beam m.Beam
			note.Beam = &beam

			beam.Number = 1
			beam.EnclosedText = "begin"
		}

		// parameter.addNote(&measure, m.Enum_Step_A)
		// parameter.addNote(&measure, m.Enum_Step_B)
	}
}

func (parameter *Parameter) addNote(measure *m.A_measure, step m.Enum_Step) {
	var note m.Note
	measure.Note = append(measure.Note, &note)

	note.Default_x = "82.98"
	note.Default_y = "-25"

	var pitch m.Pitch
	note.Pitch = &pitch

	pitch.Step = step
	pitch.Alter = "-1"
	pitch.Octave = 4

	note.Duration = "1"
	note.Voice = "1"

	var type_ m.Note_type
	note.Type = &type_

	type_.EnclosedText = "eighth"

	var stem m.Stem
	note.Stem = &stem
	stem.EnclosedText = "down"

	var beam m.Beam
	note.Beam = &beam

	beam.Number = 1
	beam.EnclosedText = "begin"
}

// generateTimestampedFilename creates a filename using the current timestamp (YYYYMMDD_HHMMSS).
func generateTimestampedFilename(prefix, extension string) string {
	// Format current date and time up to the second
	timestamp := time.Now().Format("20060102_150405")
	// Construct the filename
	return fmt.Sprintf("%s_%s%s", prefix, timestamp, extension)
}
