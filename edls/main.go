package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"time"

	"github.com/AJRDRGZ/fileinfo"
	"github.com/fatih/color"
	"golang.org/x/exp/constraints"
)

func main() {
	// Flags de Filtrado
	flagPattern := flag.String("p", "", "Fitra por patrón.")
	flagAll := flag.Bool("a", false, "Muestra todos los archivos, hasta los ocultos.")
	flagNumberRecords := flag.Int("n", 0, "Numero de archivos a mostrar.")

	// Flags de ordenamiento
	hasOrderByTime := flag.Bool("t", false, "Ordena por tiempo, primero el más antiguo.")
	hasOrderBySize := flag.Bool("s", false, "Ordena por tamaño, primero el más pequeño.")
	hasOrderReverse := flag.Bool("r", false, "Ordena la salida reversada")

	//  Parseo para capturar todas las flags
	flag.Parse()
	// obtener el argumento de la ruta
	path := flag.Arg(0)

	if path == "" {
		path = "."
	}
	dirs, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	fs := []file{}
	for _, dir := range dirs {
		isHidden := isHidden(dir.Name(), path)

		if isHidden && !*flagAll {
			continue
		}

		if *flagPattern != "" {
			isMatched, err := regexp.MatchString("(?i)"+*flagPattern, dir.Name())
			if err != nil {
				panic(err)
			}
			if !isMatched {
				continue
			}
		}

		f, err := getFile(dir, isHidden)
		if err != nil {
			panic(err)
		}

		fs = append(fs, f)
	}

	if !*hasOrderBySize || !*hasOrderByTime {
		orderByName(fs, *hasOrderReverse)
	}

	if *hasOrderBySize && !*hasOrderByTime {
		orderBySize(fs, *hasOrderReverse)
	}

	if *hasOrderByTime {
		orderByTime(fs, *hasOrderReverse)
	}

	if *flagNumberRecords == 0 || *flagNumberRecords > len(fs) {
		*flagNumberRecords = len(fs)
	}
	printList(fs, *flagNumberRecords)
}

func mySort[T constraints.Ordered](i, j T, isReverse bool) bool {
	if isReverse {
		return i > j
	}
	return i < j
}

func orderByTime(files []file, isReverse bool) {
	sort.SliceStable(files, func(i, j int) bool {
		return mySort(
			files[i].modTime.Unix(),
			files[j].modTime.Unix(),
			isReverse,
		)
	})
}

func orderByName(files []file, isReverse bool) {
	sort.SliceStable(files, func(i, j int) bool {
		return mySort(
			strings.ToLower(files[i].name),
			strings.ToLower(files[j].name),
			isReverse,
		)
	})
}

func orderBySize(files []file, IsReverse bool) {
	sort.Slice(files, func(i, j int) bool {
		return mySort(
			files[i].size,
			files[j].size,
			IsReverse,
		)
	})
}

func printList(fs []file, nRecords int) {
	for _, file := range fs[0:nRecords] {
		style := mapStyleByFileType[file.fileType]
		fmt.Printf("%s %s %s %10d %s %s %s %s %s\n", file.mode,
			file.userName, file.groupName, file.size, file.modTime.Format(time.DateTime),
			style.icon, setColor(file.name, style.color), style.symbol, markHiddenFiles(file.isHidden))
	}
}
func getFile(dir fs.DirEntry, isHidden bool) (file, error) {
	info, err := dir.Info()
	if err != nil {
		return file{}, fmt.Errorf("dir.info(): %v", err)

	}

	userName, groupName := fileinfo.GetUserAndGroup(info.Sys())

	f := file{
		name:      dir.Name(),
		isDir:     dir.IsDir(),
		isHidden:  isHidden,
		userName:  userName,
		groupName: groupName,
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
	case f.isDir:
		f.fileType = fileDirectory
	case isExec(*f):
		f.fileType = fileExecutable
	case isCompress(*f):
		f.fileType = fileCompress
	case isImage(*f):
		f.fileType = fileImage
	default:
		f.fileType = fileRegular
	}
}

func setColor(nameFile string, styleColor color.Attribute) string {
	switch styleColor {
	case color.FgBlue:
		return blue(nameFile)
	case color.FgGreen:
		return green(nameFile)
	case color.FgRed:
		return red(nameFile)
	case color.FgMagenta:
		return magenta(nameFile)
	case color.FgCyan:
		return cyan(nameFile)
	default:
		return nameFile
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
		strings.HasSuffix(f.name, rar) ||
		strings.HasSuffix(f.name, tar) ||
		strings.HasSuffix(f.name, zip7) ||
		strings.HasSuffix(f.name, tgz) ||
		strings.HasSuffix(f.name, gz) ||
		strings.HasSuffix(f.name, debian) ||
		strings.HasSuffix(f.name, archpkg)

}

func isImage(f file) bool {
	return strings.HasSuffix(f.name, png) ||
		strings.HasSuffix(f.name, jpg) ||
		strings.HasSuffix(f.name, jpeg) ||
		strings.HasSuffix(f.name, gif) ||
		strings.HasSuffix(f.name, bmp)

}

func isHidden(fileName, basePath string) bool {
	filePath := fileName
	if runtime.GOOS == Windows {
		filePath = path.Join(basePath, fileName)
	}
	return fileinfo.IsHidden(filePath)
}

func markHiddenFiles(isHidden bool) string {
	if !isHidden {
		return ""
	}
	return yellow("‡")
}
