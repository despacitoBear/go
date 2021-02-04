package main

import (
	"fmt"
	"log"
	"os"
)

//creating temp dir for files
func main() {
	var ParentFolderPath string = ""
	fmt.Println("Enter folder which you want to be ISOed, ZIPed")
	fmt.Scanf(ParentFolderPath)
	if os.IsExist(ParentFolderPath) {
		//smth
	} else {
		fmt.Println("Folder doesn't exist, try again")
		panicfmt.Scanf(ParentFolderPath))
	}

}
func CreateTempFolder() {
	err := os.Mkdir("temp_dir", 0755)
	if err {
		log.Fatal(err)
		panic(err)
	}

}

func CompressionZip(string path) string {

	return //path to compressed file
}
