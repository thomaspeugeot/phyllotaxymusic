package models

import (
	"log"
	"os"
	"path/filepath"
	"time"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func (stager *Stager) generateSSG() {
	stager.UpdatePhyllotaxyStage()

	// change the parameter
	parameter := stager.parameter

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
	stager.UpdateAndCommitSVGStage()

	// asks svg to generates an svg file
	var svg_ *svg.SVG
	for k := range *svg.GetGongstructInstancesSet[svg.SVG](stager.svgStage) {
		svg_ = k
		break
	}

	var svgText_ *svg.SvgText
	for k := range *svg.GetGongstructInstancesSet[svg.SvgText](stager.svgStage) {
		svgText_ = k
		break
	}
	_ = svgText_

	if svgText_ == nil {
		svgText_ = new(svg.SvgText).Stage(stager.svgStage)
	}

	svg_.IsSVGFileGenerated = true
	stager.svgStage.Commit()

	// check out the svg stage that must contains the svg text
	// after waiting for 500 ms
	time.Sleep(500 * time.Millisecond)

	stager.svgStage.Checkout()

	for k := range *svg.GetGongstructInstancesSet[svg.SvgText](stager.svgStage) {
		svgText_ = k
		break
	}
	_ = svgText_

	// start by copying the static directory
	CopyDirectory("../../../vendor/github.com/fullstack-lang/gong/lib/ssg/go/defaults/static",
		"../../../static",
	)

	pathToGeneratedSVG := parameter.PathToGeneratedSVG

	// --- Start: Remove existing pathToGeneratedSVG directory ---
	log.Printf("Attempting to remove existing directory: %s", pathToGeneratedSVG)
	err := os.RemoveAll(pathToGeneratedSVG)
	if err != nil {
		// Log the error but continue, as MkdirAll below might still succeed if the path didn't exist
		// or if the error was related to something already gone.
		// If MkdirAll fails later, that error will be caught.
		log.Printf("Warning: Error removing directory '%s': %v. Attempting to continue.", pathToGeneratedSVG, err)
	} else {
		log.Printf("Successfully removed existing directory: %s", pathToGeneratedSVG)
	}
	// --- End: Remove existing pathToGeneratedSVG directory ---

	// --- Start: Generate root _index.md for the Content ---
	// 1. Create the root content directory (MkdirAll handles existing directories gracefully)
	err = os.MkdirAll(pathToGeneratedSVG, 0755) // Use 0755 for standard directory permissions
	if err != nil {
		log.Printf("Error creating root content directory '%s': %v\n", pathToGeneratedSVG, err)
		// Decide if this is fatal or if chapter generation should still proceed.
		// For now, let's return if the root directory cannot be created.
		return
	}
	log.Printf("Root content directory created or already exists: %s\n", pathToGeneratedSVG)

	imageFilePath := filepath.Join(pathToGeneratedSVG, "image1.svg")

	err = os.WriteFile(imageFilePath, []byte(svgText_.Text), 0644) // Use 0644 for standard file permissions
	if err != nil {
		log.Printf("Error writing root file '%s': %v\n", imageFilePath, err)
		// Decide if this is fatal or if chapter generation should still proceed.
		// For now, let's return if the root index file cannot be written.
		return
	}
	log.Printf("iumage file created successfully: %s\n", imageFilePath)
	// --- End: Generate iumage for the Content ---

	svg_.IsSVGFileGenerated = false
	stager.svgStage.Commit()

	/* copy necessary images for the geenration */
	CopyFile("../../../images/bach2ndFugue.png", filepath.Join(pathToGeneratedSVG, "bach2ndFugue.png"))

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
