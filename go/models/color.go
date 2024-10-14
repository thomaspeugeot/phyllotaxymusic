package models

// Distinct colors corresponding to integers between 0 and 20
var colors = []string{
	"#FF0000", // Red
	"#00FF00", // Green
	"#0000FF", // Blue
	"#FFFF00", // Yellow
	"#FF00FF", // Magenta
	"#00FFFF", // Cyan
	"#800000", // Maroon
	"#008000", // DarkGreen
	"#000080", // Navy
	"#808000", // Olive
	"#800080", // Purple
	"#008080", // Teal
	"#C0C0C0", // Silver
	"#FFA500", // Orange
	"#A52A2A", // Brown
	"#7FFF00", // Chartreuse
	"#DC143C", // Crimson
	"#4682B4", // SteelBlue
	"#DA70D6", // Orchid
	"#ADFF2F", // GreenYellow
	"#F08080", // LightCoral
}

// GenerateColor generates a distinct SVG color for integers between 0 and 20.
func GenerateColor(n int) string {
	if n < 0 || n > 20 {
		return "#000000" // Return black for out of range numbers
	}
	return colors[n]
}
