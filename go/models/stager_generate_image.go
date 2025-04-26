package models

import (
	"log"
	"os"
	"time"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func (stager *Stager) generateImage(imageFilePath string, err error) (*svg.SVG, bool) {
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

	err = os.WriteFile(imageFilePath, []byte(svgText_.Text), 0644) // Use 0644 for standard file permissions
	if err != nil {
		log.Printf("Error writing root file '%s': %v\n", imageFilePath, err)
		// Decide if this is fatal or if chapter generation should still proceed.
		// For now, let's return if the root index file cannot be written.
		return nil, true
	}
	log.Printf("iumage file created successfully: %s\n", imageFilePath)

	svg_.IsSVGFileGenerated = false

	stager.svgStage.Commit()
	return svg_, false
}
