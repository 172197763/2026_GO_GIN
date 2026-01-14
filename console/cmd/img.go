package cmd

import (
	"fmt"
	"gin_test/api/common"
	"image"
	"image/jpeg"
	"image/png"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var path string
var imgCmd = &cobra.Command{
	Use:   "imgCmd",
	Short: "图片处理",
	Long:  `图片处理`,
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func run() {
	// 检查路径是否存在
	info, err := os.Stat(path)
	if err != nil {
		fmt.Printf("错误: 路径不存在 - %s\n", err)
		return
	}

	if !info.IsDir() {
		fmt.Println("错误: 指定路径不是目录")
		return
	}
	// 遍历目录中的所有文件
	err = filepath.WalkDir(path, func(filePath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// 只处理文件，跳过子目录
		if d.IsDir() {
			return nil
		}

		// 检查文件扩展名是否为 jpg 或 png
		ext := strings.ToLower(filepath.Ext(filePath))
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
			return nil
		}

		fmt.Printf("正在处理文件: %s\n", filePath)

		// 打开输入图片
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("无法打开文件 %s: %v\n", filePath, err)
			return nil // 继续处理其他文件
		}
		defer file.Close()

		img, _, err := image.Decode(file)
		if err != nil {
			fmt.Printf("无法解码图片 %s: %v\n", filePath, err)
			return nil // 继续处理其他文件
		}
		blockSize := 5
		c := 2
		// 二值化（阈值设为 128，可调整）
		binaryImg := common.AdaptiveBinarize(img, blockSize, c)

		// 生成输出文件名
		dir := filepath.Dir(filePath)
		resDir := filepath.Join(dir, "res")

		// 确保 res 目录存在
		if err := os.MkdirAll(resDir, 0755); err != nil {
			fmt.Printf("无法创建输出目录 %s: %v\n", resDir, err)
			return nil // 继续处理其他文件
		}

		baseName := strings.TrimSuffix(filepath.Base(filePath), ext)
		outputPath := filepath.Join(resDir, fmt.Sprintf("%s_%d_%d_binary.png", baseName, blockSize, c))

		// 创建输出文件
		outFile, err := os.Create(outputPath)
		if err != nil {
			fmt.Printf("无法创建输出文件 %s: %v\n", outputPath, err)
			return nil // 继续处理其他文件
		}
		defer outFile.Close()

		// 根据原始文件格式选择编码器
		switch ext {
		case ".jpg", ".jpeg":
			err = jpeg.Encode(outFile, binaryImg, &jpeg.Options{Quality: 100})
		case ".png":
			err = png.Encode(outFile, binaryImg)
		}

		if err != nil {
			fmt.Printf("无法编码图片 %s: %v\n", outputPath, err)
			return nil // 继续处理其他文件
		}

		fmt.Printf("二值化完成，已保存为 %s\n", outputPath)
		return nil
	})

	if err != nil {
		fmt.Printf("遍历目录时出错: %v\n", err)
	} else {
		fmt.Println("所有图片处理完成")
	}
}
func init() {
	imgCmd.Flags().StringVarP(&path, "path", "r", "", "folder path (required)")
	imgCmd.MarkFlagRequired("path")
}
