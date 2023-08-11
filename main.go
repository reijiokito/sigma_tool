package main

import (
	"archive/zip"
	"os"
)
import "io"

func main() {
	parseInput()
	//specUrl := "https://www.3gpp.org/ftp/Specs/archive/29_series/29.512/29512-000.zip"
	//resp, err := http.Get(specUrl)
	//if err != nil {
	//	fmt.Printf("err: %s", err)
	//}
	//
	//defer resp.Body.Close()
	//if resp.StatusCode != 200 {
	//	return
	//}
	//
	//// Create the file
	//out, err := os.Create("example.zip")
	//if err != nil {
	//	fmt.Printf("err: %s", err)
	//}
	//defer out.Close()
	//
	//// Write the body to file
	//_, err = io.Copy(out, resp.Body)
	//fmt.Printf("err: %s", err)
	//// Unzip the downloaded ZIP file
	//err = unzip("example.zip", "/home/congnt/Documents")
	//if err != nil {
	//	fmt.Printf("err: %s", err)
	//}
	//err = os.Remove("example.zip")
	//if err != nil {
	//	fmt.Printf("Remove err: %s", err)
	//	return
	//}
}

func unzip(zipFile, destFolder string) error {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer reader.Close()

	err = os.MkdirAll(destFolder, 0755)
	if err != nil {
		return err
	}

	for _, file := range reader.File {
		filePath := destFolder + "/" + file.Name
		if file.FileInfo().IsDir() {
			os.MkdirAll(filePath, file.Mode())
			continue
		}

		fileWriter, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer fileWriter.Close()

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		_, err = io.Copy(fileWriter, fileReader)
		if err != nil {
			return err
		}
	}

	return nil
}
