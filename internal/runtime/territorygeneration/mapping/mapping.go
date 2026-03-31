package mapping

import (
	"image"
	"image/draw"
	"image/png"
	"log"
	"minifarm/internal/runtime/territorygeneration/matrices"
	"minifarm/internal/runtime/territorygeneration/worldobjects"
	"os"
	"strconv"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func loadImage(imgName string) image.Image {
	imgFile, err := os.Open(imgName)
	checkErr(err)
	defer imgFile.Close()

	img, _, err := image.Decode(imgFile)
	checkErr(err)

	return img
}

func saveImage(imgName string, img image.Image) {
	outFile, err := os.Create(imgName)
	checkErr(err)
	defer outFile.Close()

	png.Encode(outFile, img)
}

func getWidth(img image.Image) int {
	return img.Bounds().Dx()
}

func getHeight(img image.Image) int {
	return img.Bounds().Dy()
}

func newRGBAImage(width, height int) *image.RGBA {
	return image.NewRGBA(image.Rect(0, 0, width, height))
}

func pasteImage(dst draw.Image, dp image.Point, src image.Image) {
	rect := image.Rect(dp.X, dp.Y, dp.X+getWidth(src), dp.Y+getHeight(src))
	draw.Draw(dst, rect, src, image.Point{0, 0}, draw.Src)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// чанки
var GroundMapData matrices.Matrix
var PlantMapData matrices.Matrix

// PaintMap генерирует новую карту;
// записывает изображение
func PaintMap() {
	width := 0
	height := 0

	var dstImg *image.RGBA
	var nameF string
	var img image.Image

	used := make([]worldobjects.RiverPoint, 0)
	for l := 0; l < 1; l++ {
		GroundMapData = matrices.NewMatrix(100, 100)
		worldobjects.GenerateRiver(&GroundMapData, &used)

		PlantMapData = matrices.NewMatrix(100, 100)
		worldobjects.GenerateTree(&PlantMapData, 50)

		width = GroundMapData.M() * 16
		height = GroundMapData.N() * 16
		dstImg = newRGBAImage(width, height)

		for m := range GroundMapData {
			for n := range GroundMapData[m] {
				if GroundMapData[m][n] == 0 {
					nameF = "assets/tiles/трава.png"
				} else {
					nameF = "assets/tiles/" + strconv.Itoa(int(GroundMapData[m][n])) + ".png"
				}
				img = loadImage(nameF)
				pasteImage(dstImg, image.Point{m * 16, n * 16}, img)

			}
		}
		saveImage("assets/map/map.png", dstImg)
	}
}
