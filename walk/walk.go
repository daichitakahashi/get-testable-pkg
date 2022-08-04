package walk

import (
	"io/fs"
	"log"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/mod/modfile"
)

func Walk(fsys fs.FS, modFile string, excludes []string) (map[string]*PackageInfo, error) {
	e, err := parseExcludes(excludes)
	if err != nil {
		return nil, err
	}

	data, err := fs.ReadFile(fsys, modFile)
	if err != nil {
		return nil, err
	}
	mod, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		log.Fatal(err)
	}
	base := mod.Module.Mod.Path

	pkgInfo := map[string]*PackageInfo{}
	err = fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			sp := filepath.ToSlash(p)
			if ignoreDir(sp) || e.Excluded(sp) {
				return fs.SkipDir
			}
			return nil
		}

		if !strings.HasSuffix(p, ".go") || strings.HasPrefix(p, "_") {
			return nil // ignore
		}

		dir := path.Join(base, filepath.ToSlash(filepath.Dir(p)))
		i, ok := pkgInfo[dir]
		if !ok {
			i = &PackageInfo{}
			pkgInfo[dir] = i
		}
		if strings.HasSuffix(p, "_test.go") {
			i.TestFiles = append(i.TestFiles, p)
		} else {
			i.GoFiles = append(i.GoFiles, p)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return pkgInfo, nil
}

type exclusion []*regexp.Regexp

func parseExcludes(excludes []string) (r exclusion, _ error) {
	for _, e := range excludes {
		rr, err := regexp.Compile(e)
		if err != nil {
			return nil, err
		}
		r = append(r, rr)
	}
	return r, nil
}

func (e exclusion) Excluded(p string) bool {
	for _, r := range e {
		if r.MatchString(p) {
			return true
		}
	}
	return false
}

func ignoreDir(p string) bool {
	base := filepath.Base(p)
	if base == "testdata" || strings.HasPrefix(base, "_") {
		return true
	}
	return false
}
