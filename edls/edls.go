package main

import "time"

// File Types
const (
	fileRegular int = iota
	fileDirectory
	fileExecutable
	fileCompress
	fileImage
	fileLink
)

// File Extension
const (
	executable = ".exe"
	debian     = ".deb"
	archpkg    = ".zst"
	zip        = ".zip"
	gz         = ".gz"
	tar        = ".tar"
	rar        = ".rar"
	png        = ".png"
	jpg        = ".jpg"
	jpeg       = ".jpeg"
	gif        = ".gif"
)

// File Structure
type file struct {
	name      string
	fileType  int
	isDir     bool
	isHidden  bool
	userName  string
	groupName string
	size      int64
	modTime   time.Time
	mode      string
}

// styleFileType Opciones para dar estilos a los archivos desde el mapa.
type styleFileType struct {
	icon   string
	color  string
	symbol string
}

// mapStyleFileType Es el mapa para los estilos de los tipos de archivos.
var mapStyleFileType = map[int]styleFileType{
	fileRegular:    {icon: "docicon"},
	fileDirectory:  {icon: "FolderIcon", color: "Blue", symbol: "/"},
	fileExecutable: {icon: "execIcon", color: "Green", symbol: "*"},
	fileCompress:   {icon: "CompressIcon", color: "red"},
	fileImage:      {icon: "ImageIcon", color: "Magenta"},
	fileLink:       {icon: "LinkIcon", color: "Cyan"},
}
