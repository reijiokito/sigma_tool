package main

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
)

func displayPkgs(pkgNames []string, action string) {
	log(1, "Are you sure you would like to %s the following packages:", bolden(action))

	for _, pkgToDisplay := range pkgNames {
		indent(pkgToDisplay)
	}

	confirm("y", "(y/n)")
}

func setDest(dest string) {
	DestFolder = dest

	chapLog("=>", "", fmt.Sprintf("Destionation folder: %s", DestFolder))
}

func initPlugin(pkgName string) {
	chapLog("=>", "", fmt.Sprintf("Initializing %s", pkgName))
	//Step1: Get Plugin Registry (Server expose an API to get plugin registry)
	//Step2: Get plugin URLs
	//_ := isValidURL(pkgName)
	//Step3: Download plugin -> Setup to local machine
}

func downloadZipFile(pluginUrl string) {
	specUrl := pluginUrl
	resp, err := http.Get(specUrl)
	if err != nil {
		errorLogRaw("URL Zip file err: %s", bolden(err.Error()))
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		errorLogRaw("Get Zip file err: %s", bolden(err.Error()))
		return
	}

	zipName := "plugin.zip"

	out, err := os.Create(zipName)
	if err != nil {
		errorLogRaw("Create Zip err: %s", bolden(err.Error()))
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)

	err = unzip(zipName, DestFolder)
	if err != nil {
		errorLogRaw("UnZip err: %s", bolden(err.Error()))
	}
	err = os.Remove(zipName)
	if err != nil {
		errorLogRaw("Remove zip file err %s:", bolden(err.Error()))
		return
	}
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
