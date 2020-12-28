package checker

import (
	sitemap "github.com/oxffaa/gopher-parse-sitemap"
)

func Parse(filename string) ([]sitemap.Entry, error) {
	var entryArray []sitemap.Entry

	err := sitemap.ParseFromFile(filename, func(e sitemap.Entry) error {
		entryArray = append(entryArray, e)
		return nil
	})
	if err != nil {
		return entryArray, err
	}
	return entryArray, nil
}
