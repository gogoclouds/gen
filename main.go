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
		Name       string
		StructName string
		FileName   string
		PkgName    string
	}{
		Name: "UserInfo",
	}
	module.PkgName = strings.ToLower(module.Name)

	filepaths := []string{
		"./tmpl/repo.tmpl",
		"./tmpl/service.tmpl",
		"./tmpl/server.tmpl",
		"./tmpl/api.tmpl",
	}
	for _, fp := range filepaths {
		tp, err := template.ParseFiles(fp)
		if err != nil {
			log.Println(err)
			return
		}

		fType := strings.ToLower(strings.TrimSuffix(filepath.Base(fp), filepath.Ext(fp)))
		f, err := pkg.MustOpen(pkg.CamelCaseToUdnderscore(module.Name)+"_"+fType+".go", strings.ToLower(module.Name))
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
