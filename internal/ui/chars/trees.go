package chars

import (
	"github.com/cyklan/photosynthesis/internal/models"
)

var treeChars = map[models.TreeState]string{
	models.Empty:   "⬡",
	models.Sapling: "P",
	models.Small:   "S",
	models.Medium:  "M",
	models.Large:   "L",
}

func GetTreeChar(tree *models.Tree) string {
	return treeChars[tree.TreeState]
}
