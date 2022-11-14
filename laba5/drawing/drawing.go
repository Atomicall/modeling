package drawing

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	"github.com/Atomicall/Mod/laba5/request"
)

func DrawGraphic(r request.Requests, savePath string) image.Image {
	var (
		scale      = 20
		imgLength  = int(r.GetTotalTime()) * scale * 6 / 4
		imgHeight  = 100 * scale
		lineLength = int(r.GetTotalTime()) * scale
		canvas     = image.NewRGBA(image.Rect(0, 0, imgLength, imgHeight))
		colors     = []color.RGBA{
			{0, 100, 0, 255},   //green
			{50, 205, 50, 255}, //limegreen
			{237, 5, 5, 1},     //red

		}
		step      = 10
		lineWidth = 1 * scale / 4
		lineStart = (imgLength-lineLength)/2 - step
		lineEnd   = lineStart + lineLength + step
		rectWidth = imgHeight / 20
	)
	// draws background
	draw.Draw(canvas, image.Rect(0, 0, imgLength, imgHeight),
		&image.Uniform{color.White}, image.Point{}, draw.Src)
	// draws 1st timeline
	draw.Draw(canvas, image.Rect(lineStart, imgHeight/4, lineEnd, imgHeight/4+lineWidth),
		&image.Uniform{colors[0]}, image.Point{}, draw.Src)
	// draws 2nd timeline
	draw.Draw(canvas, image.Rect(lineStart, 3*imgHeight/4, lineEnd, 3*imgHeight/4+lineWidth),
		&image.Uniform{colors[0]}, image.Point{}, draw.Src)
	pos_x := lineStart
	pos_y := imgHeight / 4
	verticalLineYstart := imgHeight/4 - step - lineWidth
	verticalLineYend := 3*imgHeight/4 + step + 2*lineWidth
	directionFlag := false
	for i, item := range r {
		pos_x += (int(float64(scale) * (item.TimeToArrive)))
		if i%10 == 0 {
			directionFlag = !directionFlag
		}
		draw.Draw(canvas, image.Rect(pos_x+1, verticalLineYstart, pos_x+lineWidth+1, verticalLineYend),
			&image.Uniform{colors[2]}, image.Point{}, draw.Src)
		if i < len(r)-1 {
			draw.Draw(canvas, image.Rect(pos_x+int(item.TimeToServe*float64(scale))+lineWidth+1, pos_y,
				pos_x+int(item.TimeToServe*float64(scale))+(int(r[i+1].TimeToArrive*float64(scale)))+1, pos_y+rectWidth),
				&image.Uniform{colors[1]}, image.Point{}, draw.Src)
		}
		pos_x += (int(float64(scale) * (item.TimeToServe)))
		if directionFlag {
			pos_y += rectWidth

		} else {
			pos_y -= rectWidth
		}
	}

	file, err := os.Create(savePath) // ... now lets save output image
	if err != nil {
		panic(err)
	}
	defer file.Close()
	png.Encode(file, canvas)
	return canvas
}

func ShowGraphic(img image.Image) {

}
