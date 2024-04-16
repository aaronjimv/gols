package main

import "time"

// Windows os system
const Windows = "windows"

// file types
const (
	fileRegular int = iota
	fileDirectory
	fileExecutable
	fileCompress
	fileImage
	fileLink
)

// file extensions
const (
	exe = ".exe"
	deb = ".deb"
	zip = ".zip"
	gz  = ".gz"
	tar = ".tar"
	rar = ".rar"
	png = ".png"
	jpg = ".jpg"
	gif = ".gif"
)

type file struct {
	name             string
	fileType         int
	isDir            bool
	isHidden         bool
	userName         string
	groupName        string
	size             int64
	modificationTime time.Time
	mode             string
}

type styleFileType struct {
	symbol string
	color  string
	icon   string
}

var mapStyleByFileType = map[int]styleFileType{
	fileRegular:    {icon: "📄"},
	fileDirectory:  {icon: "📂", color: "color.FgBlue", symbol: "/"},
	fileExecutable: {icon: "🚀", color: "color.FgGreen", symbol: "*"},
	fileCompress:   {icon: "📦", color: "color.FgRed"},
	fileImage:      {icon: "📸", color: "color.FgMagenta"},
	fileLink:       {icon: "🔗", color: "color.FgCyan"},
}
