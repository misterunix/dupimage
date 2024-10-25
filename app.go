package main

import (
	"log"
	"os"
	"strings"

	"github.com/misterunix/sniffle/hashing"
)

func app(path string) {

	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		if file.IsDir() {
			continue
		}

		filename := path + file.Name()
		if strings.HasSuffix(filename, ".png") ||
			strings.HasSuffix(filename, ".jpg") ||
			strings.HasSuffix(filename, ".jpeg") {
			fe := fentry{name: filename}
			fe.hash, err = hashing.FileHash(hashing.MD5, fe.name)
			if err != nil {
				log.Fatal(err)
			}

			sidecard := fe.name + ".xmp"
			_, err := os.Stat(sidecard) // dont actually need the file info
			if err != nil {
				fe.sidecard = ""
			} else {
				fe.sidecard = sidecard
			}

			fentries = append(fentries, fe)
		}
	}
}

func checkdups() {

	for _, f1 := range fentries {
		for _, f2 := range fentries {
			if f1.name == f2.name {
				continue
			}
			if f1.hash == f2.hash {
				log.Println("Duplicate: ", f1.name, f2.name)
			}
		}
	}
}
