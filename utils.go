package main

import "net/url"

func isValidURL(urlInput string) bool {
	log(1, "Checking if %s is a valid URL...", bolden(urlInput))
	_, err := url.ParseRequestURI(urlInput)

	return err == nil
}
