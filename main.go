package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"slices"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

var jsonFlag bool
var sortFlag string
var orderFlag string

func init() {
	flag.BoolVar(&jsonFlag, "json", false, "output JSON")
	flag.StringVar(&sortFlag, "sort", "depth", "sort on lines,exports or depth")
	flag.StringVar(&orderFlag, "order", "desc", "order asc or desc")
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s [directory]:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	dir, err := filepath.Abs(flag.Arg(0))

	if err != nil {
		log.Fatal(err)
	}

	p, err := parseDir(dir)

	if err != nil {
		log.Fatal(err)
	}

	pkgs := flattenPackages(p)
	sort(pkgs, orderFlag, sortFlag)

	if jsonFlag {
		printJSON(pkgs)
		return
	}

	printTable(pkgs, dir)
}

func printTable(pkgs []Package, dir string) {
	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		Headers("PACKAGE", "DIR", "LINES", "EXPORTS", "DEPTH")

	for _, p := range pkgs {
		rel, err := filepath.Rel(dir, p.Dir)

		if err != nil {
			log.Fatal(err)
		}

		t.Row(p.Name, rel, fmt.Sprintf("%d", p.Lines), fmt.Sprintf("%d", len(p.Exports)), fmt.Sprintf("%.2f", p.Depth))
	}

	fmt.Println(t)
}

func printJSON(pkgs []Package) {
	b, err := json.MarshalIndent(pkgs, "", "  ")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}

func flattenPackages(p Package) (pkgs []Package) {
	for _, subpkg := range p.SubPackages {
		pkgs = append(pkgs, flattenPackages(subpkg)...)
	}

	if p.Files > 0 && len(p.Exports) > 0 {
		p.SubPackages = nil
		pkgs = append(pkgs, p)
	}

	return pkgs
}

func sort(pkgs []Package, order, sort string) {
	slices.SortFunc(pkgs, func(x, y Package) int {
		var a, b Package
		if order == "asc" {
			a, b = x, y
		} else {
			a, b = y, x
		}

		switch sort {
		case "lines":
			return a.Lines - b.Lines
		case "exports":
			return len(a.Exports) - len(b.Exports)
		default:
			return int(a.Depth - b.Depth)
		}
	})
}
