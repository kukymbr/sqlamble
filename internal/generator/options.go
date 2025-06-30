package generator

import (
	"fmt"
	"strings"
)

const (
	DefaultPackageName          = "queries"
	DefaultSourceDir            = "."
	DefaultTargetDir            = "internal/" + DefaultPackageName
	DefaultSourceFilesExtension = ".sql"
)

type Options struct {
	// PackageName is a target package name of the generated code.
	PackageName string

	// SourceDir is a directory of the SQL files.
	SourceDir string

	// TargetDir is a target go code directory.
	TargetDir string

	SourceFilesExt []string
}

func (opt Options) Debug() string {
	return fmt.Sprintf(
		"package=%s; "+
			"source=%s; "+
			"target=%s; "+
			"ext=%s",
		opt.PackageName,
		opt.SourceDir,
		opt.TargetDir,
		strings.Join(opt.SourceFilesExt, ","),
	)
}
