package main

import (
	"image"
	"image/color"
	"os"
	"time"

	"github.com/llgcode/draw2d"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"

	"github.com/golang/freetype/truetype"
	"github.com/golang/glog"
"log"

	"golang.org/x/image/bmp"
	"golang.org/x/image/font/gofont/goregular"

	"github.com/wiless/waveshare"
)

var mono = true
var epd waveshare.EPD

func main() {
	waveshare.InitHW()
	draw2d.SetFontFolder(".")
	epd.Init(true)
	epdimg := ImageGenerate()
	kavimg := waveshare.LoadImage("kavishbw.jpg")
_=epdimg	
_=kavimg
	log.Println("Loading kavish..")
	UpdateImage(*kavimg)
	//UpdateImage(epdimg)
	time.Sleep(2*time.Second)
	UpdateImage(epdimg)

	epd.Init(false)
	epd.DrawLine(30,1,0x00)
	epd.DrawLine(50,2,0x00)
	epd.DisplayFrame()
return
	//	time.Sleep(2*time.Second)
//	log.Println("Loading Geometry.....")
//	UpdateImage(epdimg)
	// Toggling frames
//	for k:=0;k<4;k++{
//		time.Sleep(5*time.Second)
//		log.Println("Toggling Image...")
//		epd.DisplayFrame()
//	}
	

	log.Println("Partial Updating ...")
	for {
	PartialUpdate()
	time.Sleep(1*time.Second)
	}



}
func UpdateImage(epdimg image.Gray) {

	epd.SetFrame(epdimg) // set both frames with same image
}

func PartialUpdate() {

	epd.Init(false)
	timeimg := image.NewRGBA(image.Rect(0, 0, 48, 96))
	gc := draw2dimg.NewGraphicContext(timeimg)
	gc.ClearRect(0, 0, 48, 96)
	gc.SetFillColor(color.White)
	gc.SetStrokeColor(color.Black)
	draw2dkit.Rectangle(gc, 0, 0, 96, 48)
	gc.SetLineWidth(1.5)
	// gc.StrokeStringAt("Hey I am good", 0, 10)
	gc.FillStroke()
	gc.Save()
	draw2dimg.SaveToPngFile("subimage.png", timeimg)
	gimg := ConvertToGray(timeimg)
	SaveBMP("subimage.bmp", gimg)
	
	epd.SetSubFrame(8, 8, gimg)
}
func ConvertToGray(cimg image.Image) *image.Gray {
	b := cimg.Bounds()
	gimg := image.NewGray(b)
	RR := b.Max.Y
	CC := b.Max.X
	var cg color.Gray
	mono = true
	for r := 0; r < RR; r++ {
//		fmt.Println()

		for c := 0; c < CC; c++ {
			oldPixel := cimg.At(c, r)

			// gscale, _, _, _ := color.GrayModel.Convert(oldPixel).RGBA()
			cg = color.GrayModel.Convert(oldPixel).(color.Gray)

			// convert to monochrome
			if mono {
				str := ""
				_=str
				if cg.Y > 0 {
					cg.Y = 255
					str = "1"
				} else {
					cg.Y = 0
					str = "0"
				}
//				fmt.Print(str)
			}
			gimg.SetGray(r, c, cg)

		}
	}
	return gimg
}

func SaveBMP(fname string, img image.Image) {
	fp, fe := os.Create(fname)
	if fe != nil {
		glog.Errorln("Unable to Save ", fname)
		return
	}
	bmp.Encode(fp, img)
	fp.Close()
}

func ImageGenerate() (epdimg image.Gray) {
	// img := image.NewGray(image.Rect(0, 0, 200, 210))
	img := image.NewRGBA(image.Rect(0, 0, 200, 200))
	// for r := 0; r < 200; r++ {
	// 	for c := 0; c < 200; c++ {
	// 		img.Set(r, c, color.White)
	// 	}
	// }

	gc := draw2dimg.NewGraphicContext(img)
	gc.ClearRect(0, 0, 200, 200)
gc.Rotate(-3.141/2)
	gc.SetStrokeColor(color.Black)
	gc.SetFillColor(color.Black)
	gc.SetLineWidth(2)
	draw2dkit.Rectangle(gc, 30, 30, 100, 100)
	gc.Stroke()

	gc.SetStrokeColor(color.Black)
	gc.SetFillColor(color.White)
	gc.SetLineWidth(4)
	draw2dkit.Circle(gc, 100, 100, 30)
	gc.FillStroke()

	draw2dkit.RoundedRectangle(gc, 105, 105, 180, 180, 10, 10)
	gc.Stroke()

	gc.SetFillColor(color.Black)

	gc.SetStrokeColor(color.Black)
	// gc.Close()
	// gc.Restore()
	// gc.SetFillColor(color.Black)
	font, _ := truetype.Parse(goregular.TTF)
	// font, _ := truetype.Parse(gobold.TTF)

	gc.SetFont(font)
	gc.SetFontSize(14)
	gc.SetLineWidth(2.5)
	msg := "ABCDEFGHIJKLMNOPQ"
	// L, T, R, B := gc.GetStringBounds(msg)

	// fmt.Println("L T R B", L, T, R, B)
	gc.StrokeStringAt(msg, 0, 20)
	// gc.FillStroke()
	gc.SetFontSize(20)
	gc.SetLineWidth(4)
	datestr := time.Now().Format(time.Stamp)
	gc.StrokeStringAt(datestr, 0, 170)
	gc.FillStroke()
	

	gc.Close()
	
	draw2dimg.SaveToPngFile("hello.png", img)

	f1, _ := os.Create("input.bmp")

	bmp.Encode(f1, img)
	f1.Close()

	/// grayimage
	b := img.Bounds()

	gimg := image.NewGray(b)
	var cg color.Gray
	mono = true
	for r := 0; r < b.Max.Y; r++ {
		for c := 0; c < b.Max.X; c++ {
			oldPixel := img.At(c, r)

			// gscale, _, _, _ := color.GrayModel.Convert(oldPixel).RGBA()
			cg = color.GrayModel.Convert(oldPixel).(color.Gray)

			// convert to monochrome
			if mono {
				if cg.Y > 0 {
					cg.Y = 255
				} else {
					cg.Y = 0
				}

			}
			gimg.SetGray(r, c, cg)

		}
	}
	///

	////

	epdimg = waveshare.Mono2ByteImage(gimg)

	f, e := os.Create("output.bmp")
	glog.Errorln(e)
	bmp.Encode(f, gimg)
	f.Close()

	return epdimg

}
