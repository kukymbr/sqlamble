package command

import (
	"context"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/kukymbr/sqlamble/internal/formatter"
	"github.com/kukymbr/sqlamble/internal/generator"
	"github.com/kukymbr/sqlamble/internal/utils"
	"github.com/kukymbr/sqlamble/internal/version"
	"github.com/spf13/cobra"
)

func Run() {
	if err := run(); err != nil {
		utils.PrintErrorf("%s", err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}

func run() error {
	opt := generator.Options{}
	silent := false

	var cmd = &cobra.Command{
		Use:   "sqlamble",
		Short: "Embed SQL in go code",
		Long: `Generates structured SQL getters in go code.
See https://github.com/kukymbr/sqlamble for info.`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
			defer cancel()

			if err := prepareOptions(&opt); err != nil {
				return err
			}

			gen, err := generator.New(opt)
			if err != nil {
				return err
			}

			utils.PrintHellof("Hi, this is sqlamble generator.")
			utils.PrintDebugf("Options: " + opt.Debug())

			return gen.Generate(ctx)
		},
		Version: version.GetVersion(),
	}

	initFlags(cmd, &opt, &silent)

	cmd.PersistentPreRun = func(_ *cobra.Command, _ []string) {
		utils.SetSilentMode(silent)
	}

	return cmd.Execute()
}

func initFlags(cmd *cobra.Command, opt *generator.Options, silent *bool) {
	cmd.PersistentFlags().BoolVarP(silent, "silent", "s", false, "Silent mode")

	cmd.Flags().StringVar(
		&opt.PackageName,
		"package",
		generator.DefaultPackageName,
		"Target package name of the generated code",
	)

	cmd.Flags().StringVar(
		&opt.SourceDir,
		"source",
		generator.DefaultSourceDir,
		"Directory containing SQL files",
	)

	cmd.Flags().StringVar(
		&opt.TargetDir,
		"target",
		generator.DefaultTargetDir,
		"Directory for the generated Go files",
	)

	cmd.Flags().StringSliceVar(
		&opt.SourceFilesExt,
		"ext",
		[]string{generator.DefaultSourceFilesExtension},
		"If set, source files will be filtered by these suffixes in names",
	)

	cmd.Flags().StringVar(
		&opt.Formatter,
		"fmt",
		formatter.DefaultFormatter,
		"Formatter used to format generated go files (gofmt|noop)",
	)

	cmd.Flags().StringVar(
		&opt.QueryGetterSuffix,
		"query-suffix",
		generator.DefaultQueryGetterSuffix,
		"Suffix for query getter functions",
	)
}

func prepareOptions(opt *generator.Options) error {
	opt.PackageName = strings.TrimSpace(opt.PackageName)
	opt.QueryGetterSuffix = strings.TrimSpace(opt.QueryGetterSuffix)

	if opt.PackageName == "" {
		opt.PackageName = generator.DefaultPackageName
	}

	if opt.SourceDir == "" {
		opt.SourceDir = generator.DefaultSourceDir
	}

	if opt.TargetDir == "" {
		opt.TargetDir = generator.DefaultTargetDir
	}

	if opt.Formatter == "" {
		opt.Formatter = formatter.DefaultFormatter
	}

	if opt.QueryGetterSuffix == "" {
		opt.QueryGetterSuffix = generator.DefaultQueryGetterSuffix
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
		return []string{generator.DefaultSourceFilesExtension}
	}

	return res
}
