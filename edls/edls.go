package main

import (
	"time"

	"github.com/fatih/color"
)

// Windows os System
const Windows = "windows"

// File Types
const (
	fileRegular int = iota
	fileDirectory
	fileExecutable
	fileCompress
	fileImage
	fileLink
)

// File Extensions
const (
	executable = ".exe"
	debian     = ".deb"
	archpkg    = ".zst"
	bin        = ".bin"
	zip        = ".zip"
	rar        = ".rar"
	tar        = ".tar"
	zip7       = ".7z"
	tgz        = ".tgz"
	gz         = ".gz"
	png        = ".png"
	jpg        = ".jpg"
	jpeg       = ".jpeg"
	gif        = ".gif"
	bmp        = ".bmp"
)

// file Estructura del archivo
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
	color  color.Attribute
	symbol string
}

// mapStyleByFileType Es el mapa para los estilos de los tipos de archivos.
var mapStyleByFileType = map[int]styleFileType{
	fileRegular:    {icon: "📄 "},
	fileDirectory:  {icon: "📁 ", color: color.FgBlue, symbol: "/"},
	fileExecutable: {icon: "⚙️ ", color: color.FgGreen, symbol: "*"},
	fileCompress:   {icon: "🗃 ", color: color.FgRed},
	fileImage:      {icon: "📷 ", color: color.FgMagenta},
	fileLink:       {icon: "📎 ", color: color.FgCyan},
}

var (
	blue    = color.New(color.FgBlue).Add(color.Bold).SprintFunc()
	green   = color.New(color.FgGreen).Add(color.Bold).SprintFunc()
	red     = color.New(color.FgRed).Add(color.Bold).SprintFunc()
	magenta = color.New(color.FgMagenta).Add(color.Bold).SprintFunc()
	cyan    = color.New(color.FgCyan).Add(color.Bold).SprintFunc()
	yellow  = color.New(color.FgYellow).Add(color.Bold).SprintFunc()
)
