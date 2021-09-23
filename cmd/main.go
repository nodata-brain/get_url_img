package main

import (
	"fmt"

	"github.com/nodata-brain/get_url_img/pkg"
)

func main() {

	err := guri.New()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("success")
}
