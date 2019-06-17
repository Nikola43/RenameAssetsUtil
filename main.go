package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

func main() {

	var sizes [4]string
	sizes[0] = "mdpi"
	sizes[1] = "hdpi"
	sizes[2] = "xhdpi"
	sizes[3] = "xxhdpi"

	var fileNameWithOutExt string
	var fileName string

	files, readError := ioutil.ReadDir("./")
	if readError != nil {
		log.Fatal(readError)
	}

	for _, f := range files {
		fileName = f.Name()
		if !strings.Contains(fileName, "dpi") {
			if path.Ext(fileName) == ".png" {
				fileNameWithOutExt = strings.Split(fileName, ".")[0]
			}
		}
	}

	CreateDirIfNotExist(fileNameWithOutExt)
	CreateDirIfNotExist(fileNameWithOutExt + "/mipmap-" + sizes[0]) // mdpi
	CreateDirIfNotExist(fileNameWithOutExt + "/mipmap-" + sizes[1]) // hdpi
	CreateDirIfNotExist(fileNameWithOutExt + "/mipmap-" + sizes[2]) // xhdpi
	CreateDirIfNotExist(fileNameWithOutExt + "/mipmap-" + sizes[3]) // xxhdpi

	for _, f := range files {
		fileName = f.Name()
		// if is image
		if path.Ext(fileName) == ".png" {

			if !strings.Contains(fileName, "dpi") {
				fmt.Println(fileName + " -> " + fileNameWithOutExt + "/mipmap-" + sizes[0] + "/" + fileNameWithOutExt + ".png")
				_, err := copy(fileName, fileNameWithOutExt+"/mipmap-"+sizes[0]+"/"+fileNameWithOutExt+".png")
				if err != nil {
					fmt.Println("ERROR: " + err.Error())
					return
				}
			}

			if strings.Contains(fileName, "hdpi") {
				if !strings.Contains(fileName, "xhd") {
					fmt.Println(fileName + " -> " + fileNameWithOutExt + "/mipmap-" + sizes[1] + "/" + fileNameWithOutExt + ".png")
					_, err := copy(fileName, fileNameWithOutExt+"/mipmap-"+sizes[1]+"/"+fileNameWithOutExt+".png")
					if err != nil {
						fmt.Println("ERROR: " + err.Error())
						return
					}
				}
			}

			if strings.Contains(fileName, "xhdpi") {
				if !strings.Contains(fileName, "xxhd") {
					fmt.Println(fileName + " -> " + fileNameWithOutExt + "/mipmap-" + sizes[2] + "/" + fileNameWithOutExt + ".png")
					_, err := copy(fileName, fileNameWithOutExt+"/mipmap-"+sizes[2]+"/"+fileNameWithOutExt+".png")
					if err != nil {
						fmt.Println("ERROR: " + err.Error())
						return
					}
				}
			}

			if strings.Contains(fileName, "xxhdpi") {
				fmt.Println(fileName + " -> " + fileNameWithOutExt + "/mipmap-" + sizes[3] + "/" + fileNameWithOutExt + ".png")
				_, err := copy(fileName, fileNameWithOutExt+"/mipmap-"+sizes[3]+"/"+fileNameWithOutExt+".png")
				if err != nil {
					fmt.Println("ERROR: " + err.Error())
					return
				}
			}
		}
	}

	fmt.Println("DONE")
}

func contains(s string, subs string) {
	var size = len(s)
	var i int
	// var j int

	for i = 0; i < size; i++ {
		if s[i] == subs[0] {

		}
		fmt.Println(i)
	}
}

func GetSize(name string) string {
	var size = ""
	if strings.Contains(name, "dpi") {
		size += string(name[len(name)-10])
		size += string(name[len(name)-9])
		size += string(name[len(name)-8])
		size += string(name[len(name)-7])
		size += string(name[len(name)-6])
		size += string(name[len(name)-5])
		return size
	}

	return "mdpi"
}

func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err.Error())
		}
	}
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
func RenameFile(before string, after string) error {
	err := os.Rename(before, after)
	if err != nil {
		return err
	}
	return nil
}
