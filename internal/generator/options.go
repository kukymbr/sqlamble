package generator

import (
	"fmt"
	"strings"

	"github.com/kukymbr/sqlamble/internal/formatter"
	"github.com/kukymbr/sqlamble/internal/utils"
)

const (
	DefaultPackageName          = "queries"
	DefaultSourceDir            = "."
	DefaultTargetDir            = "internal/" + DefaultPackageName
	DefaultFormatter            = formatter.GoFmt
	DefaultSourceFilesExtension = ".sql"
	DefaultQueryGetterSuffix    = "Query"
)

type Options struct {
	// PackageName is a target package name of the generated code.
	// Default is "queries".
	PackageName string

	// SourceDir is a directory of the SQL files.
	// Default is the current directory (most applicable for go:generate).
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
	return fmt.Sprintf("%#v", opt)
}

func prepareOptions(opt *Options) error {
	opt.PackageName = strings.TrimSpace(opt.PackageName)
	opt.QueryGetterSuffix = strings.TrimSpace(opt.QueryGetterSuffix)

	if opt.PackageName == "" {
		opt.PackageName = DefaultPackageName
	}

	if opt.SourceDir == "" {
		opt.SourceDir = DefaultSourceDir
	}

	if opt.TargetDir == "" {
		opt.TargetDir = DefaultTargetDir
	}

	if opt.Formatter == "" {
		opt.Formatter = DefaultFormatter
	}

	if opt.QueryGetterSuffix == "" {
		opt.QueryGetterSuffix = DefaultQueryGetterSuffix
	}

	opt.QueryGetterSuffix = utils.FirstUpper(opt.QueryGetterSuffix)

	if err := utils.ValidateIsDir(opt.SourceDir); err != nil {
		return err
	}

	if err := utils.ValidatePackageName(opt.PackageName); err != nil {
		return err
	}

	if err := utils.ValidateQueryGetterSuffix(opt.QueryGetterSuffix); err != nil {
		return err
	}

	if err := utils.EnsureDir(opt.TargetDir); err != nil {
		return err
	}

	opt.SourceFilesExt = prepareSourceFilesExt(opt.SourceFilesExt)

	return nil
}

func prepareSourceFilesExt(exts []string) []string {
	// Filtering in disabled.
	if len(exts) == 1 && exts[0] == "" {
		return nil
	}

	res := make([]string, 0, len(exts))

	for _, ext := range exts {
		ext = strings.TrimSpace(ext)
		if ext != "" {
			res = append(res, ext)
		}
	}

	if len(res) == 0 {
		return []string{DefaultSourceFilesExtension}
	}

	return res
}
