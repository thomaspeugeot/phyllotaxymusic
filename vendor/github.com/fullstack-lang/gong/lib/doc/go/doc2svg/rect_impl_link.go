package doc2svg

import (
	"log"

	gongdoc_models "github.com/fullstack-lang/gong/lib/doc/go/models"
	gongsvg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
)

// LinkImplLink
// it meets the interface of LinkImplInterface
type LinkImplLink struct {
	link         *gongdoc_models.Link
	gongdocStage *gongdoc_models.Stage
}

func NewLinkImplLink(
	link *gongdoc_models.Link,
	gongdocStage *gongdoc_models.Stage) (linkImplLink *LinkImplLink) {

	linkImplLink = new(LinkImplLink)
	linkImplLink.link = link
	linkImplLink.gongdocStage = gongdocStage

	return
}

func (linkImplLink *LinkImplLink) LinkUpdated(updatedLink *gongsvg_models.Link) {

	log.Println("LinkImplLink:LinkUpdated")

	// update the link
	linkImplLink.link.StartOrientation = gongdoc_models.OrientationType(updatedLink.StartOrientation)
	linkImplLink.link.StartRatio = updatedLink.StartRatio
	linkImplLink.link.EndOrientation = gongdoc_models.OrientationType(updatedLink.EndOrientation)
	linkImplLink.link.EndRatio = updatedLink.EndRatio
	linkImplLink.link.CornerOffsetRatio = updatedLink.CornerOffsetRatio

	linkImplLink.gongdocStage.CommitWithSuspendedCallbacks()
}
