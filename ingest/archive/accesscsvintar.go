package archive

import (
	"archive/tar"
	"fmt"
	"io"
	"path"
)

const cSVName = "data.csv"

// AccessCSVInTar provides an io.Reader that points to the CSV file inside
// the tar archive provided.
func AccessCSVInTar(tarFile io.Reader) (io.Reader, error) {

	tr := tar.NewReader(tarFile)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("Searching tar file: %v", err)
		}

		if path.Base(hdr.Name) == cSVName {
			return tr, nil
		}
	}
	return nil, fmt.Errorf("csv file (%s) not found in archive", cSVName)
}
