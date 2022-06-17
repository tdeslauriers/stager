package im

import (
	"image"
	"runtime"
	"sync"
	"sync/atomic"
)

// what all pics to be the same height for site grid
const height int = 200 // pixels?

// current impl is nearest-neighbor, no filter
func Resize(img image.Image) *image.NRGBA {

	src := newScanner(img)

	// preserve aspect ratio
	// height const
	width := int((float64(height) * float64(src.w) / float64(src.h)) + .5) // int rounds down

	dst := image.NewNRGBA(image.Rect(0, 0, width, height))
	dx := float64(img.Bounds().Dx()) / float64(width)
	dy := float64(img.Bounds().Dy()) / float64(height)

	parrallel(0, height, func(ys <-chan int) {
		for y := range ys {
			sy := int((float64(y) + .5) * dy)
			dstOff := y * dst.Stride
			for x := 0; x < width; x++ {
				sx := int((float64(x) + .5) * dx)
				src.scan(sx, sy, sx+1, sy+1, dst.Pix[dstOff:dstOff+4])
				dstOff += 4
			}
		}
	})

	return dst
}

// parrallel processing
var maxProcs int64

func parrallel(start, stop int, fn func(<-chan int)) {

	count := stop - start
	if count < 1 {
		return
	}

	// determine how many threads to use
	procs := runtime.GOMAXPROCS(0)
	limit := int(atomic.LoadInt64(&maxProcs))
	if procs > limit && limit > 0 {
		procs = limit
	}
	if procs > count {
		procs = count
	}

	c := make(chan int, count)
	for i := start; i < stop; i++ {
		c <- i
	}
	close(c)

	var wg sync.WaitGroup
	for i := 0; i < procs; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fn(c)
		}()
	}
	wg.Wait()
}

// image scanning
type scanner struct {
	image image.Image
	w, h  int
}

func newScanner(img image.Image) *scanner {

	return &scanner{
		image: img,
		w:     img.Bounds().Dx(),
		h:     img.Bounds().Dy(),
	}
}

func (s *scanner) scan(x1, y1, x2, y2 int, dst []uint8) {
	switch img := s.image.(type) {
	case *image.YCbCr:
		j := 0
		x1 += img.Rect.Min.X
		x2 += img.Rect.Min.X
		y1 += img.Rect.Min.Y
		y2 += img.Rect.Min.Y

		hy := img.Rect.Min.Y / 2
		hx := img.Rect.Min.X / 2
		for y := y1; y < y2; y++ {
			iy := (y-img.Rect.Min.Y)*img.YStride + (x1 - img.Rect.Min.X)

			var yBase int
			switch img.SubsampleRatio {
			case image.YCbCrSubsampleRatio444, image.YCbCrSubsampleRatio422:
				yBase = (y - img.Rect.Min.Y) * img.CStride
			case image.YCbCrSubsampleRatio420, image.YCbCrSubsampleRatio440:
				yBase = (y/2 - hy) * img.CStride
			}

			for x := x1; x < x2; x++ {
				var ic int
				switch img.SubsampleRatio {
				case image.YCbCrSubsampleRatio444, image.YCbCrSubsampleRatio440:
					ic = yBase + (x - img.Rect.Min.X)
				case image.YCbCrSubsampleRatio422, image.YCbCrSubsampleRatio420:
					ic = yBase + (x/2 - hx)
				default:
					ic = img.COffset(x, y)
				}

				yy1 := int32(img.Y[iy]) * 0x10101
				cb1 := int32(img.Cb[ic]) - 128
				cr1 := int32(img.Cr[ic]) - 128

				r := yy1 + 91881*cr1
				if uint32(r)&0xff000000 == 0 {
					r >>= 16
				} else {
					r = ^(r >> 31)
				}

				g := yy1 - 22554*cb1 - 46802*cr1
				if uint32(g)&0xff000000 == 0 {
					g >>= 16
				} else {
					g = ^(g >> 31)
				}

				b := yy1 + 116130*cb1
				if uint32(b)&0xff000000 == 0 {
					b >>= 16
				} else {
					b = ^(b >> 31)
				}

				d := dst[j : j+4 : j+4]
				d[0] = uint8(r)
				d[1] = uint8(g)
				d[2] = uint8(b)
				d[3] = 0xff

				iy++
				j += 4
			}
		}
	}
}
