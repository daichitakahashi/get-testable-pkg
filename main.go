package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/daichitakahashi/get-testable-pkg/walk"
	"github.com/fatih/color"
)

func main() {
	pkgInfo, err := walk.Walk(os.DirFS("."), "go.mod", nil)
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
			_, _ = fmt.Fprintln(os.Stderr, color.HiRedString("no test file found: %s", pkg))
		}
	}

	// shuffle

	if len(testablePackages) == 0 {
		log.Fatal(color.HiRedString("no testable packages"))
	} else if len(noTestPackages) == 0 {
		_, _ = fmt.Fprintln(os.Stderr, color.HiGreenString("all packages are testable"))
	}
	fmt.Println(strings.Join(testablePackages, " "))
}
