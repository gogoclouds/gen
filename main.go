package main

import (
	"github.com/gogoclouds/gen/pkg"
	"log"
	"path/filepath"
	"strings"
	"text/template"
)

func main() {
	var module = struct {
		Project string // github.com/gogoclouds/gen
		Struct  string // User
		// 可选
		Path       string // ./internal
		ApiVersion string // v1
		File       string // user
		Model      string // sql table
	}{
		Path:    "./internal",
		Project: "github.com/gogoclouds/gen",
		Struct:  "User",
		Model:   "SysUser",
	}
	if module.ApiVersion == "" {
		module.ApiVersion = "v1"
	}
	if module.File == "" {
		module.File = strings.ToLower(module.Struct)
	}
	if module.Model == "" {
		module.Model = module.Struct
	}

	filepaths := []string{
		"./internal/tmpl/mvc/router.tmpl",
		"./internal/tmpl/mvc/api.tmpl",
		"./internal/tmpl/mvc/server.tmpl",
		"./internal/tmpl/mvc/service.tmpl",
		"./internal/tmpl/mvc/repo.tmpl",
	}
	for _, fp := range filepaths {
		tp, err := template.ParseFiles(fp)
		if err != nil {
			log.Println(err)
			return
		}

		filename := module.File + ".go"
		dir := filepath.Join(module.Path, module.File) // ./internal/user
		fType := strings.ToLower(strings.TrimSuffix(filepath.Base(fp), filepath.Ext(fp)))
		switch fType {
		case "router":
			filename = "router.go"
		case "api":
			dir = filepath.Join(dir, fType, module.ApiVersion)
		case "server":
			dir = filepath.Join(dir, "handler")
		default:
			dir = filepath.Join(dir, fType)
		}
		f, err := pkg.MustOpen(filename, dir)
		if err != nil {
			panic(err)
		}
		err = tp.Execute(f, module)
		f.Close()
		if err != nil {
			log.Println(err)
		}
	}
}
