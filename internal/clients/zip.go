package clients

import (
	"archive/zip"
	"io"
	"log"
	"os"
)

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
