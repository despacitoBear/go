package main

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

//creating temp dir for files
func main() {
	//создание папки для временного хранения в проекте
	createTempFolder()
	var ParentFolderPath string = "/home/ilya/Documents/test"
	//fmt.Println("Enter folder which you want to be ISOed, ZIPed")
	//fmt.Scanf("%s\n", &ParentFolderPath)
	_, err := os.Stat(ParentFolderPath)
	if err == nil {
		fmt.Println("Ready to go")
	} else {
		fmt.Println("Alredy exists")
	}
	compressToTar(ParentFolderPath, createTempFolder())
}
func createTempFolder() string {
	err := os.Mkdir("temp_dir", 0755)
	if err != nil {
		//log.Fatal(err)
		//panic(err)
	} else if os.IsExist(err) {
		fmt.Println("folder is already exists")
	}
	return "temp_dir"
}

//заводская функция сжатия  для tar
func compressToTar(source, target string) error {
	filename := filepath.Base(source)
	target = filepath.Join(target, fmt.Sprintf("%s.tar", filename))
	tarfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer tarfile.Close()

	tarball := tar.NewWriter(tarfile)
	defer tarball.Close()

	info, err := os.Stat(source)
	if err != nil {
		return nil
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(source)
	}

	return filepath.Walk(source,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			header, err := tar.FileInfoHeader(info, info.Name())
			if err != nil {
				return err
			}

			if baseDir != "" {
				header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
			}

			if err := tarball.WriteHeader(header); err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(tarball, file)
			return err
		})
}

//разархивация завод
func untar(victim, destination string) error {
	reader, err := os.Open(victim)
	if err != nil {
		return err
	}
	defer reader.Close()
	tarReader := tar.NewReader(reader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		path := filepath.Join(destination, header.Name)
		info := header.FileInfo()
		if info.IsDir() {
			if err = os.MkdirAll(path, info.Mode()); err != nil {
				return err
			}
			continue
		}

		file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode())
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(file, tarReader)
		if err != nil {
			return err
		}
	}
	return nil
}
