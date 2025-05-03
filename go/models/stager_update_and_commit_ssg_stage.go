package models

import (
	"log"

	ssg "github.com/fullstack-lang/gong/lib/ssg/go/models"
)

func (stager *Stager) UpdateAndCommitSsgStage() {
	stager.ssgStage.Reset()

	parameters := GetGongstructInstancesSet[Parameter](stager.phyllotaxymusicStage)

	var parameter *Parameter
	if len(*parameters) != 1 {
		log.Fatalln("There should be one parameter, only one shall remains in scottland")
	}
	for k, _ := range *parameters {
		parameter = k
	}
	_ = parameter

	var content *Content
	contents := *GetGongstructInstancesSet[Content](stager.phyllotaxymusicStage)
	for k := range contents {
		content = k
	}

	var ssgChapters []*ssg.Chapter
	for _, chapter := range content.Chapters {
		ssgChapters = append(ssgChapters,
			&ssg.Chapter{
				Name:           chapter.Name,
				MardownContent: chapter.MardownContent,
			},
		)
	}

	ssg.StageBranch(
		stager.ssgStage,
		&ssg.Content{
			Name:           content.Name,
			MardownContent: content.MardownContent,

			LayoutPath: content.LayoutPath,
			StaticPath: content.StaticPath,

			ContentPath: content.ContentPath,
			OutputPath:  content.OutputPath,

			Target: ssg.Target(content.Target),

			Chapters: ssgChapters,
		},
	)

	stager.ssgStage.Commit()
}
