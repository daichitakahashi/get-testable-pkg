package walk

type PackageInfo struct {
	TestFiles []string
	GoFiles   []string
}

func (i *PackageInfo) Testable() bool {
	return len(i.TestFiles) > 0
}
