package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/daichitakahashi/get-testable-pkg/walk"
	"github.com/fatih/color"
)

var (
	verbose bool
	shuffle bool
)

func init() {
	flag.BoolVar(&verbose, "v", false, "output non-testable packages and other info")
	flag.BoolVar(&shuffle, "shuffle", true, "shuffle testable package")
	flag.Parse()
	log.SetFlags(0)
}

func main() {
	pkgInfo, err := walk.Walk(os.DirFS("."), "go.mod", flag.Args())
	if err != nil {
		log.Fatal(color.HiRedString(err.Error()))
	}

	noTestPackages := make([]string, 0, len(pkgInfo))
	testablePackages := make([]string, 0, len(pkgInfo))

	for pkg, info := range pkgInfo {
		if !info.Testable() {
			noTestPackages = append(noTestPackages, pkg)
			continue
		}
		testablePackages = append(testablePackages, pkg)
	}

	if len(noTestPackages) > 0 {
		sort.Strings(noTestPackages)
		for _, pkg := range noTestPackages {
			optionalLog(verbose, color.HiRedString("no test file found: %s", pkg))
		}
	}

	if len(testablePackages) == 0 {
		log.Fatal(color.HiRedString("no testable packages"))
	} else if len(noTestPackages) == 0 {
		optionalLog(verbose, color.HiGreenString("all packages are testable"))
	}

	if !shuffle {
		sort.Strings(testablePackages)
	}
	fmt.Println(strings.Join(testablePackages, " "))
}

func optionalLog(v bool, s string) {
	if v {
		_, _ = fmt.Fprintln(os.Stderr, s)
	}
}
