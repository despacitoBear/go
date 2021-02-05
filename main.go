package main

import (
	"fmt"
	"log"
	"os"
)

//creating temp dir for files
func main() {
	//создание папки для временного хранения в проекте
	CreateTempFolder()
	var ParentFolderPath string = ""
	fmt.Println("Enter folder which you want to be ISOed, ZIPed")
	fmt.Scanf(ParentFolderPath)
	_, err := os.Stat(ParentFolderPath)
	if err == nil {
		fmt.Println("Alredy exists")
	} else {
		fmt.Println("Ready to go")
	}
}
func CreateTempFolder() string {
	err := os.Mkdir("temp_dir", 0755)
	if err != nil {
		log.Fatal(err)
		panic(err)
	} else if os.IsExist(err) {
		return "already exist"
	}
	return "SomeThing went wrong"
}
