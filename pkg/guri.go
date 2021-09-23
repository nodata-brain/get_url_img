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
	for i, u := range t.url {
		//request url
		resp, err := http.Get(u)
		if err != nil {
			log.Println(err)
			continue
		}

		//byte to image
		imgByte, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			continue
		}
		img, _, _ := image.Decode(bytes.NewReader(imgByte))

		//output image
		file, err := os.Create("./img/" + strconv.Itoa(i) + ".png")
		defer file.Close()
		if err != nil {
			log.Println(err)
			continue
		}
		err = png.Encode(file, img)
		if err != nil {
			log.Println(err)
			continue
		}

	}

	return nil
}
