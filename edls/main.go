package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"regexp"
	"runtime"
	"strings"
	"time"
)

func main() {
	// Falgs de filtrado de datos
	flagPattern := flag.String("p", "", "Filter by pattern")
	// flagAll := flag.Bool("a", false, "Sort all files, including hidden files")
	flagNumberRecord := flag.Int("n", 0, "Sort a spesific number of files")

	// Flags de orden de datos.
	/*
		hasOrderByTime := flag.Bool("t", false, "Order files by time, older first")
		hasOrderBySize := flag.Bool("s", false, "Order files by size, smallest first")
		hasOrderReverse := flag.Bool("r", false, "Reverse order while sorting")
	*/
	flag.Parse() // lectura de las flags creadas

	// path contiene la ruta de los datos
	path := flag.Arg(0)

	// asigno una ruta por defecto
	if path == "" {
		path = "."
	}
	dirs, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	// fs Slice de archivos
	fs := []file{}

	for _, dir := range dirs {
		f, err := getFile(dir, false)
		if err != nil {
			panic(err)
		}

		isMatched, err := regexp.MatchString("(?i)"+*flagPattern, f.name)
		if err != nil {
			panic(err)
		}

		if !isMatched {
			continue
		}
		fs = append(fs, f)
	}

	if *flagNumberRecord == 0 || *flagNumberRecord > len(fs) {
		*flagNumberRecord = len(fs)
	}
	printList(fs, *flagNumberRecord)
}

func printList(fs []file, nRecords int) {
	for _, file := range fs[:nRecords] {
		style := mapStyleFileType[file.fileType]
		fmt.Printf("%s %s %s %15d %s %s %s %s\n", file.mode, file.userName, file.groupName, file.size, file.modTime.Format(time.DateTime), style.icon, file.name, style.symbol)
	}
}

func getFile(dir fs.DirEntry, isHidden bool) (file, error) {
	info, err := dir.Info()
	if err != nil {
		return file{}, fmt.Errorf("dir.Info(): %v", err)
	}

	f := file{
		name:      dir.Name(),
		isDir:     dir.IsDir(),
		isHidden:  isHidden,
		userName:  "miguel",
		groupName: "users",
		size:      info.Size(),
		modTime:   info.ModTime(),
		mode:      info.Mode().String(),
	}
	setFile(&f)
	return f, nil
}

func setFile(f *file) {
	switch {
	case isLink(*f):
		f.fileType = fileLink
	case isExec(*f):
		f.fileType = fileExecutable
	case isCompress(*f):
		f.fileType = fileCompress
	case f.isDir:
		f.fileType = fileDirectory
	case isImage(*f):
		f.fileType = fileImage
	default:
		f.fileType = fileRegular
	}
}

func isLink(f file) bool {
	return strings.HasPrefix(strings.ToUpper(f.mode), "L")
}

func isExec(f file) bool {
	if runtime.GOOS == Windows {
		return strings.HasSuffix(f.name, executable)
	}

	return strings.Contains(f.mode, "x")
}

func isCompress(f file) bool {
	return strings.HasSuffix(f.name, zip) ||
		strings.HasSuffix(f.name, gz) ||
		strings.HasSuffix(f.name, tar) ||
		strings.HasSuffix(f.name, rar) ||
		strings.HasSuffix(f.name, archpkg)

}

func isImage(f file) bool {
	return strings.HasSuffix(f.name, png) ||
		strings.HasSuffix(f.name, jpg) ||
		strings.HasSuffix(f.name, jpeg) ||
		strings.HasSuffix(f.name, gif) ||
		strings.HasSuffix(f.name, bmp)

}
