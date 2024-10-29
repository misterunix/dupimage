package main

import (
	"fmt"
	"log"

	"os"
	"path/filepath"
)

func main() {

	base := "."
	createLog()

	recursiveDecent(base)

	fmt.Println("Directories: ", len(rdirectories))
	log.Println("Directories: ", len(rdirectories))

	for _, d := range rdirectories {
		log.Println("Directory: ", d.name)
	}

	for _, d := range rdirectories {
		dp := d.name
		// log.Println("Directory: ", dp)
		app(dp)
	}

	for _, fe := range FileEntry {
		//fmt.Println(fe.name, fe.sidecard, fe.hash)
		log.Println(fe.name, fe.sidecard, fe.hash)
	}

	checkdups()

}

// do a recursive decent of the base directory and store the
// directories in slice 'rentries'
func recursiveDecent(base string) {
	err := filepath.WalkDir(base, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			p := path + "/"
			r := recursive{name: p, completed: false}
			rdirectories = append(rdirectories, r)
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
	}
}

/*
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
		fmt.Println(filename)
	}
}
*/
/*
		if strings.HasSuffix(fe.name, ".png") || strings.HasSuffix(fe.name, ".jpg") || strings.HasSuffix(fe.name, ".jpeg") {
			fe.hash, err = hashing.FileHash(hashing.SHA512, fe.name)
			if err != nil {
				log.Fatal(err)
			}

			fe.sidecard = fe.name + ".xmp"
			fentries = append(fentries, fe)
			count := 0
			for _, f := range fentries {
				if f.hash == fe.hash {
					count++
					if count > 1 {
						fedelete = append(fedelete, f)
					}
					continue
				}
			}
		}
	}

	fmt.Println("Count: ", len(fentries))
	fmt.Println("Delete: ", len(fedelete))

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Do you want to delete the duplicate images? (y/n): ")
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(strings.ToLower(response))

		if response == "y" || response == "yes" {
			fmt.Println("Continuing...")
			break
		} else if response == "n" || response == "no" {
			fmt.Println("Exiting...")
			return
		} else {
			fmt.Println("Invalid input. Please enter 'y' or 'n'.")
		}
	}

	for _, fe := range fedelete {

		//fmt.Printf("%s %s %s\n", fe.name, fe.sidecard, fe.hash)
		os.Remove(fe.name)
		os.Remove(fe.sidecard)
	}

	fmt.Println("Deleted: ", len(fedelete))

}

*/
