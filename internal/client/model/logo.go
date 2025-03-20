package model

import "github.com/mbndr/figlet4go"

const logoStr = "goph-keeper"

func renderLogo() string {
	ascii := figlet4go.NewAsciiRender()
	options := figlet4go.NewRenderOptions()
	options.FontColor = []figlet4go.Color{
		figlet4go.ColorGreen,
		figlet4go.ColorYellow,
		figlet4go.ColorCyan,
	}
	options.FontName = "larry3d"
	renderStr, _ := ascii.RenderOpts(logoStr, options)
	return renderStr
}
