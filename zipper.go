package main

import (
	"archive/zip"
	"io"
	"os"

	"path/filepath"
	"strings"
)

type zipIt struct {
	rawPath   string
	trimPath  string
	listFiles []string
	baseDir   string
}

func zipThis(sourceDir, target string) error {

	exitFile, err := os.Create(target)
	if err != nil {
		loger("zipThis(), os.Create() error", err)
	}
	defer exitFile.Close()

	archive := zip.NewWriter(exitFile)
	defer archive.Close()

	var z zipIt
	z.initStruct(sourceDir)

	for _, singlFile := range z.listFiles {
		err = z.fileZiper(singlFile, archive)
		if err != nil {
			loger("zipThis(), z.fileZiper() error", err)
		}
	}
	return nil
}
func (z zipIt) fileZiper(singlFile string, archive *zip.Writer) error {
	info, err := os.Stat(singlFile)
	if err != nil {
		loger("zipThis(), os.Stat() error", err)
		return err
	}
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		loger("zipThis(), zip.FileInfoHeader() error", err)
		return err
	}

	if z.baseDir != "" {
		header.Name = filepath.Join(z.baseDir, strings.TrimPrefix(singlFile, z.trimPath))
	}

	if info.IsDir() {
		header.Name += string(os.PathSeparator)
	} else {
		header.Method = zip.Deflate
	}
	writer, err := archive.CreateHeader(header)
	if err != nil {
		loger("zipThis(), archive.CreateHeader() error", err)
		return err
	}
	file, err := os.Open(singlFile)
	if err != nil {
		loger("zipThis(), os.Open() error", err)
		return err
	}
	defer file.Close()
	_, _ = io.Copy(writer, file)
	return nil
}

func getBaseDir(adress string) string {
	info, err := os.Stat(adress)
	if err != nil {
		loger("getBaseDir, os.Stat", err)
	}
	if info.IsDir() {
		baseDir := filepath.Base(adress)
		return baseDir
	}
	return ""
}
func (z *zipIt) initStruct(rawPath string) {

	z.rawPath = rawPath
	z.trimPath, _, _ = wldChk(rawPath)
	z.baseDir = getBaseDir(z.trimPath)
	z.listFiles = dirlist(rawPath)

}
