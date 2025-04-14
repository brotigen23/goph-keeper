package util

import (
	"strconv"

	"github.com/brotigen23/goph-keeper/client/internal/ui/style"
	"github.com/charmbracelet/lipgloss"
	"github.com/mbndr/figlet4go"
)

func WarningResolutions(w, h, minW, minH int) (string, bool) {
	if w < minW || h < minH {
		warning := style.ColorRed.Render("not enough space")
		warning += style.Gap

		warning += style.ColorGreen.Render(strconv.Itoa(minW) + " x " + strconv.Itoa(minH))
		warning += "\n"
		if w < minW {
			warning += style.ColorRed.Render(strconv.Itoa(w))
		} else {
			warning += style.ColorGreen.Render(strconv.Itoa(w))
		}
		warning += " x "
		if h < minH {
			warning += style.ColorRed.Render(strconv.Itoa(h))
		} else {
			warning += style.ColorGreen.Render(strconv.Itoa(h))
		}
		return lipgloss.Place(w, h, lipgloss.Center, lipgloss.Center, warning), false
	}
	return "", true
}

func RenderASCII(str string) string {
	ascii := figlet4go.NewAsciiRender()
	options := figlet4go.NewRenderOptions()
	options.FontColor = []figlet4go.Color{
		figlet4go.ColorGreen,
		figlet4go.ColorYellow,
		figlet4go.ColorCyan,
	}
	options.FontName = "larry3d"
	renderStr, _ := ascii.RenderOpts(str, options)
	return renderStr
}
