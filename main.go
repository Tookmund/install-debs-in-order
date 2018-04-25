package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/starkandwayne/install-debs-in-order/debpkg"
)

func checkDebian() {
	_, err := exec.LookPath("dpkg-deb")
	if err != nil {
		log.Fatal("Please run inside Debian/Ubuntu")
	}
}

func main() {
	checkDebian()

	targetPath := os.Args[1]

	fmt.Println("Looking for files", targetPath)
	folder, err := debpkg.NewDebianPackagesFromFolder(targetPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", folder)
	folder.RemovePreinstalledPackages()
	fmt.Printf("%#v\n", folder)
	for _, pkg := range folder.Packages {
		fmt.Printf("%#v\n", pkg)
	}
}
