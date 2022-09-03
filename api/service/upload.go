package service

import (
	"io"
	"mime/multipart"
	"os"
)

const UploadsDir = "./uploads/"

func WriteRequestFileToLocal(file *multipart.FileHeader) error {
	openedFile, err := file.Open()
	if err != nil {
		return err
	}
	defer openedFile.Close()
	// 4. create a new file for saving
	newFile, err := os.Create(UploadsDir + file.Filename)
	if err != nil {
		return err
	}
	defer newFile.Close()
	// 5. save the file
	_, err = io.Copy(newFile, openedFile)
	if err != nil {
		return err
	}
	return nil
}
