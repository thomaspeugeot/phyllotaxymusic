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
			Name: "Phyllotaxy Music",
			MardownContent: `
This site describes phyllotaxy generated music.

Precisely, phyllotaxy helps to generates a music
theme. A music theme is a small suite 
of music notes, a melodic material.

A wellknown music theme is at the start of Bach's 2nd 
fugue in C minor.

<img src="./images/` + bach2ndFugue + `" style="height: 100px; width: auto;">

### Phylotaxy

Phylotaxy means "shape of leaves". 
It is the 
science developped by botanist, 
mathematicians and physicists 
to understand why plants
have the shape they have.

In the office of Stephane Douady,
a "spiral plants" specialist I 
visited, the shelves are full of 
beautiful 
dried specimen that people 
sent to him
from all over the world. 
Each plant has spectacular
spiral arragenements of leaves
that run upward the trunks.

A common spiral plant is the pinecone.

Stephane taught me that when 
the trunk grows, 
the cells
at the stem divide themselves. New cells
 end up building
the trunk but regularly, 
one of them differentiates and
becomes a leave. 
Because this happens on a regular tempo, 
the leaves arrange themselves along 
the trunk
in spiral patterns. 
Stephane pointed out that if you draw
a curve around the trunk to isolate 
one leave cell from
newer leave cells, you end up drawing 
the "front curve"
of the plant growth.

<img src="./images/` + growthCurveOnPineCone + `" style="height: 300px; width: auto;">

Stephane then compared a new leave cell 
to the last one.
Because leave cells appear regularly, 
a new cell is 
always located with the same distance
 above and the
same angle sideway, 
a rotation near 168 degrees, 
related to the golden ratio. 
Therefore, the "front curve" has 
this interesting geometric feature : 
but on the new leave cell,
it overlaps itself
after an upward translation 
and a rotation. Let's demonstrate this.

Use a grey pen and draw the unfolded the front curve on a piece 
of paper.

<img src="./images/` + SVGfirstVoice + `" style="height: 200px; width: auto;">

Use a green pen and draw the same curve at the right of the first.

<img src="./images/` + SVGfirstVoiceAndFirstVoiceShiftedRight + `" style="height: 200px; width: auto;">

Draw it again
on a tracing paper with a red pen and put the tracing 
paper above and
shift it a bit to the right and 
a bit to the top.

<img src="./images/` + SVGfirstVoiceAndFirstVoiceShiftedRightAndSecondVoice + `" style="height: 250px; width: auto;">

We did it ! The overlapping 
is perfect 
except for one space that have
the shape of an eye (
the place for the leave).

### Generating a theme form the front curve 

Now take take the first drawing and add a canevas
of vertical measure lines and horizontal pitch lines.

For the pitch lines, the spacing between pitch line
follows the scale of your choice. If you
choose a C minor scale, the tone differences between
notes are 1, 1/2, 1, 1, 1/2, 1 and 3/2.

<img src="./images/` + SVGfirstVoiceAndFirstVoiceShiftedRightWithGrid + `" style="height: 600px; width: auto;">

Now, along the front curve, when it intersect the 
measure lines, 
draw some notes of your choosing.


<img src="./images/` + SVGfirstVoiceAndFirstVoiceShiftedRightWithGridWithNotes + `" style="height: 600px; width: auto;">


			`,

			LayoutPath: "../../../vendor/github.com/fullstack-lang/gong/lib/ssg/go/defaults/layouts",
			StaticPath: "../../../static",

			ContentPath: "content",
			OutputPath:  "../../../docs",

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
