package common

import (
	"image"
	"image/color"
)

/**
 * @descript 自适应二值化算法
 * @param blockSize 块大小(块大小通常设置为 9-21 的奇数)
 * @param c 常数偏移(通常设置为 2-10)
 * @return image.Image
 */
func AdaptiveBinarize(img image.Image, blockSize int, c int) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	// 先转换为灰度图
	grayImg := image.NewGray(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			grayImg.Set(x, y, img.At(x, y))
		}
	}

	// 计算积分图
	integralImg := calculateIntegralImage(grayImg)

	result := image.NewGray(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// 定义局部窗口
			x1 := max(0, x-blockSize/2)
			y1 := max(0, y-blockSize/2)
			x2 := min(width-1, x+blockSize/2)
			y2 := min(height-1, y+blockSize/2)

			// 计算局部平均值
			area := (x2 - x1 + 1) * (y2 - y1 + 1)
			if area == 0 {
				continue
			}

			sum := integralImg[y2][x2]
			if x1 > 0 {
				sum -= integralImg[y2][x1-1]
			}
			if y1 > 0 {
				sum -= integralImg[y1-1][x2]
			}
			if x1 > 0 && y1 > 0 {
				sum += integralImg[y1-1][x1-1]
			}

			mean := sum / area

			// 应用自适应阈值
			threshold := mean - c
			if grayImg.GrayAt(x, y).Y > uint8(threshold) {
				result.SetGray(x, y, color.Gray{Y: 255}) // 白色
			} else {
				result.SetGray(x, y, color.Gray{Y: 0}) // 黑色
			}
		}
	}

	return result
}

// 灰度转二值图（阈值 threshold，0-255）
func ImgBinarize(img image.Image, threshold uint8) *image.Gray {
	bounds := img.Bounds()
	gray := image.NewGray(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			// 转灰度（加权平均，符合人眼感知）
			grayVal := uint8((r*299 + g*587 + b*114) >> 16) // >>16 相当于 /65535*255

			// 二值化
			if grayVal > threshold {
				gray.SetGray(x, y, color.Gray{255}) // 白色
			} else {
				gray.SetGray(x, y, color.Gray{0}) // 黑色
			}
		}
	}
	return gray
}

// 计算积分图
func calculateIntegralImage(img *image.Gray) [][]int {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	integralImg := make([][]int, height)
	for i := range integralImg {
		integralImg[i] = make([]int, width)
	}

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			currentValue := int(img.GrayAt(x, y).Y)

			if x > 0 {
				currentValue += integralImg[y][x-1]
			}
			if y > 0 {
				currentValue += integralImg[y-1][x]
			}
			if x > 0 && y > 0 {
				currentValue -= integralImg[y-1][x-1]
			}

			integralImg[y][x] = currentValue
		}
	}

	return integralImg
}
