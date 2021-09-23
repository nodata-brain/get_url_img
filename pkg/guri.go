package guri

import (
	"bufio"
	"bytes"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type txt struct {
	url []string
}

//New get url img
func New() error {

	//file read
	fp, err := os.Open("./source/url.txt")
	if err != nil {
		return err
	}
	defer fp.Close()

	var t txt

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		t.url = append(t.url, scanner.Text())
	}
	//request url
	for i, u := range t.url {
		resp, err := http.Get(u)
		if err != nil {
			log.Println(err)
			continue
		}
		imgByte, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			continue
		}
		img, _, _ := image.Decode(bytes.NewReader(imgByte))
		file, err := os.Create("./img/img_" + strconv.Itoa(i) + ".png")
		if err != nil {
			log.Println(err)
			continue
		}
		defer file.Close()

		err = png.Encode(file, img)
		if err != nil {
			log.Println(err)
			continue
		}

	}

	//byte to image

	//output image

	return nil
}
