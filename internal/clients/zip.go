package clients

import (
	"archive/zip"
	"io"
	"log"
	"os"
)

//go:generate mockgen -source=./zip.go -destination=./zip_mock.go -package=clients

type ZipInterface interface {
	UnZip(filename string) error
	Remove(filename string) error
}

type Zip struct {
}

func (z *Zip) UnZip(filename string) error {
	log.Printf("Unzipping %s", filename)
	r, err := zip.OpenReader(filename)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		fpath := f.Name
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, f.Mode())
		} else {

			f, err := os.OpenFile(
				fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (z *Zip) Remove(filename string) error {
	err := os.Remove(filename)
	if err != nil {
		return err
	}

	return nil
}
