package generator

import (
	"strings"
)

const (
	DefaultPackageName          = "queries"
	DefaultSourceDir            = "."
	DefaultTargetDir            = "internal/" + DefaultPackageName
	DefaultSourceFilesExtension = ".sql"
	DefaultQueryGetterSuffix    = "Query"
)

type Options struct {
	// PackageName is a target package name of the generated code.
	// Default is "queries".
	PackageName string

	// SourceDir is a directory of the SQL files.
	// Default is current directory (most applicable for go:generate).
	SourceDir string

	// TargetDir is a target go code directory.
	// Default is "internal/queries".
	TargetDir string

	// SourceFilesExt is a list of source files extensions (or suffixes).
	// If defined, files not matching any of these suffixes, will be ignored.
	// Default is [".sql"].
	SourceFilesExt []string

	// Formatter is a name of the formatter for the generated code files.
	// Available options: gofmt (default), none.
	Formatter string

	// QueryGetterSuffix is a suffix for a queries getters names.
	// Default is "Query".
	QueryGetterSuffix string
}

func (opt Options) Debug() string {
	values := []string{
		"package_name" + "=" + opt.PackageName,
		"source_dir" + "=" + opt.SourceDir,
		"target_dir" + "=" + opt.TargetDir,
		"source_files_ext" + "=" + strings.Join(opt.SourceFilesExt, ","),
		"formatter" + "=" + opt.Formatter,
		"query_getter_suffix" + "=" + opt.QueryGetterSuffix,
	}

	return strings.Join(values, "; ")
}
