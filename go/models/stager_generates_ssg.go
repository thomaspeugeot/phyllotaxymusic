package models

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

const bach2ndFugue string = "bach2ndFugue.png"

const growthCurveOnPineCone string = "growthCurveOnPineCone.png"

const firstVoiceSVGimage string = "fistVoiceSVGimage.svg"

const firstVoiceAndSecondSVGimage string = "fistVoiceAndSecondSVGimage.svg"

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
	parameter.OriginY = 450
	memoryOfShapeIsDisplayed := make(map[ShapeInterface]bool)
	for _, shape := range parameter.Shapes {
		memoryOfShapeIsDisplayed[shape] = shape.GetIsDisplayed()
		shape.SetIsDisplayed(false)
	}
	parameter.FirstVoice.IsDisplayed = true
	stager.UpdateSVGStage()

	// asks svg to generates an svg file
	imageFilePath := filepath.Join(pathToGeneratedSVG, firstVoiceSVGimage)
	stager.grabGeneratedSVGFile(imageFilePath, 15000*time.Millisecond)

	// // generates the second image
	// parameter.FirstVoiceShiftedRigth.IsDisplayed = true
	// stager.UpdateAndCommitSVGStage()

	// imageFilePath = filepath.Join(pathToGeneratedSVG, firstVoiceAndSecondSVGimage)
	// stager.grabGeneratedSVGFile(imageFilePath, 15000*time.Millisecond)

	/*
	 RESTORE SVG GENERATION
	*/
	parameter.OriginY = originY
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
