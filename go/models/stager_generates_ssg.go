package models

import (
	"log"
	"os"
	"path/filepath"
	"time"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

const bach2ndFugue string = "bach2ndFugue.png"

const growthCurveOnPineCone string = "growthCurveOnPineCone.png"

const SVGfirstVoice string = "fistVoiceSVGimage.svg"

const SVGfirstVoiceAndFirstVoiceShiftedRight string = "SVGfirstVoiceAndFirstVoiceShiftedRight.svg"

const SVGfirstVoiceAndFirstVoiceShiftedRightAndSecondVoice string = "SVGfirstVoiceAndFirstVoiceShiftedRightAndSecondVoice.svg"

const SVGfirstVoiceAndFirstVoiceShiftedRightWithGrid string = "SVGfirstVoiceAndFirstVoiceShiftedRightWithGrid.svg"

func (stager *Stager) GenerateSSG() {
	stager.UpdatePhyllotaxyStage()

	// change the parameter
	parameter := stager.parameter
	pathToGeneratedSVG := parameter.PathToGeneratedSVG

	// start by copying the static directory
	_, shouldReturn := stager.prepareStaticDic(pathToGeneratedSVG)
	if shouldReturn {
		return
	}

	/*
		KEEP THE ORIGINAL SVG GENERATION CONFIGURATION
	*/

	originY := parameter.OriginY
	pitchDifference := parameter.PitchDifference
	parameter.OriginY = 450
	memoryOfShapeIsDisplayed := make(map[ShapeInterface]bool)
	for _, shape := range parameter.Shapes {
		memoryOfShapeIsDisplayed[shape] = shape.GetIsDisplayed()
		shape.SetIsDisplayed(false)
	}
	parameter.FirstVoice.IsDisplayed = true
	stager.UpdateSVGStage() // beware to not commit because the grabGeneratedSVGFile will commit the svg

	// asks svg to generates an svg file
	imageFilePath := filepath.Join(pathToGeneratedSVG, SVGfirstVoice)
	svg.GrabGeneratedSVGFile(stager.svgStage, imageFilePath, 5000*time.Millisecond)

	// generates the second image
	parameter.FirstVoiceShiftedRigth.IsDisplayed = true
	stager.UpdateSVGStage()

	imageFilePath = filepath.Join(pathToGeneratedSVG, SVGfirstVoiceAndFirstVoiceShiftedRight)
	svg.GrabGeneratedSVGFile(stager.svgStage, imageFilePath, 5000*time.Millisecond)

	// generates the third image
	parameter.SecondVoice.IsDisplayed = true
	parameter.PitchDifference = 0
	stager.UpdateSVGStage()

	imageFilePath = filepath.Join(pathToGeneratedSVG, SVGfirstVoiceAndFirstVoiceShiftedRightAndSecondVoice)
	svg.GrabGeneratedSVGFile(stager.svgStage, imageFilePath, 5000*time.Millisecond)

	// generates the fourth image
	parameter.SecondVoice.IsDisplayed = false
	parameter.PitchDifference = 0
	parameter.PitchLines.IsDisplayed = true
	parameter.BeatLines.IsDisplayed = true
	stager.UpdateSVGStage()

	imageFilePath = filepath.Join(pathToGeneratedSVG, SVGfirstVoiceAndFirstVoiceShiftedRightWithGrid)
	svg.GrabGeneratedSVGFile(stager.svgStage, imageFilePath, 5000*time.Millisecond)

	/*
	 RESTORE SVG GENERATION
	*/
	parameter.OriginY = originY
	parameter.PitchDifference = pitchDifference
	for _, shape := range parameter.Shapes {
		shape.SetIsDisplayed(memoryOfShapeIsDisplayed[shape])
	}
	stager.UpdateAndCommitSVGStage()

	stager.ssgStage.Generation()
}

func (*Stager) prepareStaticDic(pathToGeneratedSVG string) (error, bool) {
	CopyDirectory("../../../vendor/github.com/fullstack-lang/gong/lib/ssg/go/defaults/static",
		"../../../static",
	)

	// --- Start: Remove existing pathToGeneratedSVG directory ---
	// log.Printf("Attempting to remove existing directory: %s", pathToGeneratedSVG)
	err := os.RemoveAll(pathToGeneratedSVG)
	if err != nil {
		// Log the error but continue, as MkdirAll below might still succeed if the path didn't exist
		// or if the error was related to something already gone.
		// If MkdirAll fails later, that error will be caught.
		log.Printf("Warning: Error removing directory '%s': %v. Attempting to continue.", pathToGeneratedSVG, err)
	} else {
		// log.Printf("Successfully removed existing directory: %s", pathToGeneratedSVG)
	}
	// --- End: Remove existing pathToGeneratedSVG directory ---

	// --- Start: Generate root _index.md for the Content ---
	// 1. Create the root content directory (MkdirAll handles existing directories gracefully)
	err = os.MkdirAll(pathToGeneratedSVG, 0755) // Use 0755 for standard directory permissions
	if err != nil {
		log.Printf("Error creating root content directory '%s': %v\n", pathToGeneratedSVG, err)
		// Decide if this is fatal or if chapter generation should still proceed.
		// For now, let's return if the root directory cannot be created.
		return nil, true
	}
	// log.Printf("Root content directory created or already exists: %s\n", pathToGeneratedSVG)

	/* copy necessary images for the geenration */
	CopyFile("../../../images/"+bach2ndFugue, filepath.Join(pathToGeneratedSVG, bach2ndFugue))
	CopyFile("../../../images/"+growthCurveOnPineCone, filepath.Join(pathToGeneratedSVG, growthCurveOnPineCone))

	return err, false
}
