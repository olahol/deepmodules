package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type Package struct {
	Name        string    `json:"name"`
	Dir         string    `json:"dir"`
	Files       int       `json:"files"`
	Lines       int       `json:"lines"`
	Depth       float64   `json:"depth"`
	Exports     []string  `json:"exports"`
	SubPackages []Package `json:"-"`
}

func declNames(decls []ast.Decl) (x []string) {
	for _, d := range decls {
		switch v := d.(type) {
		case *ast.FuncDecl:
			x = append(x, v.Name.String())
		case *ast.GenDecl:
			for _, s := range v.Specs {
				switch w := s.(type) {
				case *ast.TypeSpec:
					x = append(x, w.Name.String())
				case *ast.ValueSpec:
					for _, n := range w.Names {
						x = append(x, n.String())
					}
				}
			}
		}
	}

	return x
}

func parsePackage(dir string) (p Package, err error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, func(fi fs.FileInfo) bool {
		return !strings.HasSuffix(fi.Name(), "_test.go")
	}, parser.AllErrors)

	if err != nil {
		return
	}

	p.Dir = dir

	for name, pkg := range pkgs {
		p.Name = name
		p.Files += 1

		for _, file := range pkg.Files {
			ast.FileExports(file)

			p.Lines += fset.Position(file.FileEnd).Line
			p.Exports = append(p.Exports, declNames(file.Decls)...)
		}
	}

	p.Depth = float64(p.Lines) / float64(len(p.Exports))

	return p, nil
}

func parseDir(dir string) (Package, error) {
	p, err := parsePackage(dir)

	if err != nil {
		return Package{}, err
	}

	entries, err := os.ReadDir(dir)

	if err != nil {
		return Package{}, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		c, err := parseDir(filepath.Join(dir, entry.Name()))

		if err != nil {
			continue
		}

		if len(c.Exports) == 0 && len(c.SubPackages) == 0 {
			continue
		}

		p.SubPackages = append(p.SubPackages, c)
	}

	return p, nil
}
