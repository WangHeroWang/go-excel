package main

import (
	"fmt"
	_ "image/png"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	f, err := excelize.OpenFile("./Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 插入图片
	err = f.AddPicture("Sheet1", "F10", "img.png", "")
	if err != nil {
		fmt.Println(err)
	}
	if err = f.Save(); err != nil {
		fmt.Println(err)
	}
}
