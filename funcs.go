package main

import (
	"archive/zip"
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func displayPlugins(pkgNames []string, action string) {
	log(1, "Are you sure you would like to %s the following plugins:", bolden(action))

	for _, pkgToDisplay := range pkgNames {
		indent(pkgToDisplay)
	}

	confirm("y", "(y/n)")
}

func publishPlg(plgName string, pluginDir string, version string) {
	//Step1: Get Plugin Registry (Server expose an API to get plugin registry)
	//_ := isValidURL(pkgName)
	//Step2: Check duplicate
	//Step3: Push to Plugin Registry (using API create)
}

func removePlg(plgName string) {
	destination, err := readPluginDestination()
	if err != nil {
		errorLogRaw("Destination is not valid")
		return
	}
	childFolderPath := filepath.Join(destination, plgName)

	_, err = os.Stat(childFolderPath)
	if err != nil {
		errorLogRaw("Child folder is not existed: %v", err)
		return
	}

	err = os.RemoveAll(childFolderPath)
	if err != nil {
		errorLogRaw("Remove plugin error: %v", err)
		return
	}
}

func initPlg(plgName string) {
	chapLog("=>", "", fmt.Sprintf("Initializing %s", plgName))
	//Step1: Get Plugin Registry (Server expose an API to get plugin registry)
	//Step2: Get plugin URLs
	//_ := isValidURL(pkgName)
	//Step3: Download plugin -> Setup to local machine

	url := "https://www.3gpp.org/ftp/Specs/archive/29_series/29.512/29512-010.zip"
	downloadZipFile(url)
}

func setDest(dest string) {
	if !checkFilePathIsValid(dest) {
		errorLogRaw("File path is not valid")

	}
	err := setPluginDestination(dest)
	if err != nil {
		errorLogRaw("Set Plugin destination err: %s", bolden(err.Error()))
		return
	}

	chapLog("=>", "", fmt.Sprintf("Destionation folder: %s", dest))
}

func checkFilePathIsValid(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return true
}

type Config struct {
	PluginDir string `yaml:"plugin_dir"`
}

func readPluginDestination() (string, error) {

	yamlFile, err := ioutil.ReadFile(PluginDestinationPath)
	if err != nil {
		return "", err
	}

	// Unmarshal YAML into struct
	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return "", err
	}

	return config.PluginDir, nil
}
func setPluginDestination(dest string) error {

	var config Config
	config.PluginDir = dest

	updatedYAML, err := yaml.Marshal(&config)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(PluginDestinationPath, updatedYAML, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
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

	destination, err := readPluginDestination()
	if err != nil {
		errorLogRaw("Destination is not valid")
		return
	}
	err = unzip(zipName, destination)
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
