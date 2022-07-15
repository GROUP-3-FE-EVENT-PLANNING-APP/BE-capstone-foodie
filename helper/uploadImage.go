package helper

import (
	"fmt"
	"mime/multipart"
	"time"
)

func UploadImage(fileData multipart.File, fileInfo *multipart.FileHeader) (string, error) {
	// cek ekstension file upload
	extension, err_check_extension := CheckFileExtension(fileInfo.Filename)
	if err_check_extension != nil {
		return "", fmt.Errorf("file extension error")
	}

	// check file size
	err_check_size := CheckFileSize(fileInfo.Size)
	if err_check_size != nil {
		return "", fmt.Errorf("file size error")
	}

	// memberikan nama file
	fileName := time.Now().Format("2006-01-02 15:04:05") + "." + extension

	url, errUploadImg := UploadImageToS3(fileName, fileData)

	if errUploadImg != nil {
		fmt.Println(errUploadImg)
		return "", fmt.Errorf("failed to upload file")
	}

	return url, nil

}
