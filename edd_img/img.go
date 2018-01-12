package edd_img

import (
	"encoding/base64"
	"errors"
	"image"
	"image/png"
	"io/ioutil"
	"os"
)

//OpenImageToBase64 OpenImageToBase64
func OpenImageToBase64(filename string) (string, error) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(f), nil
}

//OpenPNG 打开png图片
func OpenPNG(filename string) (image.Image, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return png.Decode(f)
}

//SavePNG 保存png图片
func SavePNG(filename string, pic image.Image) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	return png.Encode(f, pic)
}

//CutImage 裁剪图片
func CutImage(src image.Image, x, y, w, h int) (image.Image, error) {
	var subImg image.Image

	if rgbImg, ok := src.(*image.YCbCr); ok {
		subImg = rgbImg.SubImage(image.Rect(x, y, x+w, y+h)).(*image.YCbCr) //图片裁剪x0 y0 x1 y1
	} else if rgbImg, ok := src.(*image.RGBA); ok {
		subImg = rgbImg.SubImage(image.Rect(x, y, x+w, y+h)).(*image.RGBA) //图片裁剪x0 y0 x1 y1
	} else if rgbImg, ok := src.(*image.NRGBA); ok {
		subImg = rgbImg.SubImage(image.Rect(x, y, x+w, y+h)).(*image.NRGBA) //图片裁剪x0 y0 x1 y1
	} else {
		return subImg, errors.New("图片解码失败")
	}

	return subImg, nil
}

