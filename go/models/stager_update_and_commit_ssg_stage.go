package models

import (
	"fmt"
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

	ssg.StageBranch(
		stager.ssgStage,
		&ssg.Content{
			Name:           "Phyllotaxie Music",
			MardownContent: "This the is plyllotaxy music generated site",

			LayoutPath: "../../../vendor/github.com/fullstack-lang/gong/lib/ssg/go/defaults/layouts",
			StaticPath: "../../../vendor/github.com/fullstack-lang/gong/lib/ssg/go/defaults/static",

			ContentPath: "content",
			OutputPath:  "../../../public",

			Target: ssg.FILE,

			Chapters: []*ssg.Chapter{
				{
					Name:           "Introduction to phyllotaxy music",
					MardownContent: "",
				},

				{
					Name: "Parameters",
					MardownContent: "Parameters\n\n" +
						"| Parameter | Value |\n" +
						"| -------- | -------- |\n" +
						fmt.Sprintf("| N | %d |\n", parameter.N) +
						fmt.Sprintf("| M | %d |\n", parameter.M) +
						"",
				},
			},
		},
	)

	stager.ssgStage.Commit()
}
