package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/misterunix/sniffle/hashing"
)

type fentry struct {
	name     string
	sidecard string
	hash     string
}

type recursive struct {
	name      string
	completed bool
}

var fentries []fentry
var fedelete []fentry
var rentries []recursive

func main() {

	r := recursive{name: "./", completed: false}
	rentries = append(rentries, r)

	for {
		done := godown()
		if done {
			break
		}
	}

}

func godown() bool {
	var doany bool = false

	for _, d := range rentries {
		if d.completed {
			continue
		}
		files, err := os.ReadDir(d.name)
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			if file.IsDir() {
				nd := recursive{name: d.name + file.Name() + "/", completed: false}
				rentries = append(rentries, nd)
				doany = true
				continue
			}
		}
	}

	return doany
}

func app() bool {

	files, err := os.ReadDir(d.name)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			nd := recursive{name: d.name + file.Name() + "/", completed: false}
			rentries = append(rentries, nd)
			continue
		}
		fe := fentry{name: file.Name()}

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
