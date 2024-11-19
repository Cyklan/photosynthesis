package chars

import (
	"github.com/cyklan/photosynthesis/internal/models"
)

const SunChar = "🌞"

var sunChars = map[models.SunState]string{
	models.TopLeft:     "↘",
	models.TopRight:    "↙",
	models.Right:       "←",
	models.BottomRight: "↖",
	models.BottomLeft:  "↗",
	models.Left:        "→",
}

func GetSunChars(state models.SunState) string {
	return sunChars[state]
}
