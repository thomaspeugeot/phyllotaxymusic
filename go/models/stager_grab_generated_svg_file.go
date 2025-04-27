package models

import (
	"log"
	"os"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func (stager *Stager) grabGeneratedSVGFile(imageFilePath string, err error) (*svg.SVG, bool) {
	var svg_ *svg.SVG
	for k := range *svg.GetGongstructInstancesSet[svg.SVG](stager.svgStage) {
		svg_ = k
		break
	}

	var svgText_ *svg.SvgText
	svgTexts := *svg.GetGongstructInstancesSet[svg.SvgText](stager.svgStage)
	log.Println("len of svgTexts", len(svgTexts))
	for k := range svgTexts {
		svgText_ = k
		break
	}
	_ = svgText_

	if svgText_ == nil {
		svgText_ = new(svg.SvgText).Stage(stager.svgStage)
	}

	svg_.IsSVGFileGenerated = true
	svgText_.Proxy = &SvgTextProxy{
		stager:        stager,
		svgText:       svgText_,
		imageFilePath: imageFilePath,
	}

	// the front will update svgText with the content of the svg file
	// the proxy will intercept the update
	stager.svgStage.Commit()

	// code is needed here to wait for SvgTextProxy.Updated() to proceed

	svg_.IsSVGFileGenerated = false

	stager.svgStage.Commit()

	return svg_, false
}

type SvgTextProxy struct {
	stager        *Stager
	svgText       *svg.SvgText
	imageFilePath string
}

// Updated implements models.SvgTextProxyInterface.
func (s *SvgTextProxy) Updated() {
	svgText_ := s.svgText
	imageFilePath := s.imageFilePath

	if svgText_.Text == "" {
		log.Fatalln("no image")
	}

	err := os.WriteFile(imageFilePath, []byte(svgText_.Text), 0644) // Use 0644 for standard file permissions
	if err != nil {
		log.Printf("Error writing root file '%s': %v\n", imageFilePath, err)
	}
	log.Printf("iumage file created successfully: %s\n", imageFilePath)

	// inform the main thread
}
