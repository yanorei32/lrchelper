package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/yanorei32/lrchelper/def"
	"github.com/yanorei32/lrchelper/edtbeat"
	"github.com/yanorei32/lrchelper/edtlegacy"
	"github.com/yanorei32/lrchelper/lyric"
)

func youtubeUrlParse(url string) string {
	return url[len(url)-12:]
}

func init_(wd string, def def.Def) {
	if def.Mode != "legacy" {
		log.Fatal("init works with legacy mode only")
	}

	edtlrcPath := filepath.Join(wd, "_EDIT.lrc")

	if _, err := os.Stat(edtlrcPath); err == nil {
		log.Fatal("_EDIT.lrc is exist")
	}

	len, err := time.ParseDuration(def.Length)
	if err != nil {
		log.Fatal("Parse length: " + err.Error())
	}

	lrcF, err := os.Create(edtlrcPath)
	if err != nil {
		log.Fatal("Create file: " + err.Error())
	}

	lrcF.WriteString(
		edtlegacy.GenerateBase(len, def.Timing),
	)
}

func build(wd string, def def.Def) {
	lyric := lyric.Lyric{}
	lyric.Artist = def.Head.Artist
	lyric.Album = def.Head.Album
	lyric.Title = def.Head.Title
	lyric.Author = def.Head.Author
	lyric.By = def.Head.By
	lyric.Version = def.Head.Version

	switch def.Mode {
	case "legacy":
		edtLrcF, err := os.Open(filepath.Join(wd, "_EDIT.lrc"))
		if err != nil {
			log.Fatal("Open _EDIT.lrc: " + err.Error())
		}

		lyric.Lines, err = edtlegacy.Parse(bufio.NewReader(edtLrcF))
		if err != nil {
			log.Fatal("Parse _EDIT.lrc: " + err.Error())
		}

	case "beat":
		edtBeatF, err := os.Open(filepath.Join(wd, "_EDIT.txt"))
		if err != nil {
			log.Fatal("Open _EDIT.txt: " + err.Error())
		}

		lyric.Lines, err = edtbeat.Parse(bufio.NewReader(edtBeatF), def.Timing)
		if err != nil {
			log.Fatal("Parse _EDIT.txt: " + err.Error())
		}

	default:
		log.Fatal("Unknown mode: " + def.Mode)
	}

	for _, export := range def.Exports {
		lyric.Length = export.Length

		lrcF, err := os.Create(filepath.Join(wd, export.Filename))
		if err != nil {
			log.Fatal("Create file: " + err.Error())
		}

		lrcF.WriteString(lyric.Format(export.Offset))
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please set init or build")
		os.Exit(1)
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Get working directory: " + err.Error())
	}

	defYmlF, err := os.Open(filepath.Join(wd, "_DEF.yml"))
	if err != nil {
		log.Fatal("Open _DEF.yml: " + err.Error())
	}

	defYmlB, err := ioutil.ReadAll(defYmlF)
	if err != nil {
		log.Fatal("Read _DEF.yml: " + err.Error())
	}

	def := def.Def{}
	err = yaml.Unmarshal(defYmlB, &def)
	if err != nil {
		log.Fatal("Parse _DEF.yml: " + err.Error())
	}

	switch os.Args[1] {
	case "init":
		init_(wd, def)
		break

	case "build":
		build(wd, def)
		break

	default:
		fmt.Println("Please set init or build")
		os.Exit(1)
	}
}
