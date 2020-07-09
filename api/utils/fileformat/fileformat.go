package fileformat

import (
	"github.com/twinj/uuid"
	"path"
	"strings"
)

func UniqueFormat(fileName string) string {
	newFileName := strings.TrimSuffix(fileName, path.Ext(fileName))
	extension := path.Ext(fileName)
	u := uuid.NewV4()
	newFileName = newFileName + "-" + u.String() + extension
	return newFileName
}
