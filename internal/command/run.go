package command

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/kukymbr/sqlamble/internal/formatter"
	"github.com/kukymbr/sqlamble/internal/generator"
	"github.com/kukymbr/sqlamble/internal/logger"
	"github.com/kukymbr/sqlamble/internal/version"
	"github.com/spf13/cobra"
)

func Run() {
	if err := run(); err != nil {
		logger.Errorf("%s", err.Error())
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

			gen, err := generator.New(opt)
			if err != nil {
				return err
			}

			return gen.Generate(ctx)
		},
		Version: version.GetVersion(),
	}

	initFlags(cmd, &opt, &silent)

	cmd.PersistentPreRun = func(_ *cobra.Command, _ []string) {
		logger.SetSilentMode(silent)
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
		generator.DefaultFormatter,
		"Formatter used to format generated go files ("+formatter.GoFmt+"|"+formatter.Noop+")",
	)

	cmd.Flags().StringVar(
		&opt.QueryGetterSuffix,
		"query-suffix",
		generator.DefaultQueryGetterSuffix,
		"Suffix for query getter functions",
	)
}
