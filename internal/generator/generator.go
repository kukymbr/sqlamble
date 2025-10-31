package generator

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/kukymbr/sqlamble/internal/formatter"
	"github.com/kukymbr/sqlamble/internal/generator/templates"
	"github.com/kukymbr/sqlamble/internal/generator/types"
	"github.com/kukymbr/sqlamble/internal/logger"
	"github.com/kukymbr/sqlamble/internal/utils"
	"github.com/kukymbr/sqlamble/internal/version"
)

func New(opt Options) (*Generator, error) {
	if err := prepareOptions(&opt); err != nil {
		return nil, err
	}

	f, err := formatter.Factory(opt.Formatter)
	if err != nil {
		return nil, err
	}

	logger.Hellof("Hi, this is sqlamble generator.")
	logger.Debugf("Options: " + opt.Debug())

	return &Generator{
		opt:       opt,
		formatter: f,
	}, nil
}

type Generator struct {
	opt       Options
	formatter formatter.Formatter
}

func (g *Generator) Generate(ctx context.Context) error {
	_, dirs, err := g.scanDir(ctx, g.opt.SourceDir, nil)
	if err != nil {
		return err
	}

	for _, dir := range dirs {
		fc := &strings.Builder{}

		logger.Debugf("Writing %s...", dir.TargetPath)

		if err := templates.ExecuteDirTemplate(fc, dir); err != nil {
			return fmt.Errorf("%s: %w", dir.SourcePath, err)
		}

		content := g.format(ctx, []byte(fc.String()))

		if err := utils.WriteFile(content, dir.TargetPath); err != nil {
			return err
		}
	}

	logger.Successf("All done.")

	return nil
}

//nolint:funlen
func (g *Generator) scanDir(
	ctx context.Context,
	rootPath string,
	parent *types.Directory,
) (
	*types.Directory,
	[]*types.Directory,
	error,
) {
	logger.Debugf("Scanning dir: %s...", rootPath)

	name := ""
	if rootPath == g.opt.SourceDir {
		name = g.opt.PackageName
	}

	root := g.newDirectoryInstance(rootPath, name, parent)
	directories := make([]*types.Directory, 0)

	directories = append(directories, root)

	entries, err := os.ReadDir(rootPath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read dir %s: %w", rootPath, err)
	}

	for _, entry := range entries {
		if err := ctx.Err(); err != nil {
			return nil, nil, err
		}

		if strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		path := filepath.Join(rootPath, entry.Name())

		if entry.IsDir() {
			dir, dd, err := g.scanDir(ctx, path, root)
			if err != nil {
				return nil, nil, err
			}

			root.Directories = append(root.Directories, dir)
			directories = append(directories, dd...)

			continue
		}

		if !g.decideProcessFile(entry) {
			continue
		}

		query, err := g.readFile(root, path)
		if err != nil {
			return nil, nil, err
		}

		root.Queries = append(root.Queries, query)
	}

	return root, directories, nil
}

func (g *Generator) decideProcessFile(entry fs.DirEntry) bool {
	if len(g.opt.SourceFilesExt) == 0 {
		return true
	}

	for _, ext := range g.opt.SourceFilesExt {
		if strings.HasSuffix(entry.Name(), ext) {
			return true
		}
	}

	return false
}

func (g *Generator) readFile(parent *types.Directory, path string) (*types.Query, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", path, err)
	}

	return g.newQueryInstance(parent, path, string(content)), nil
}

func (g *Generator) newQueryInstance(parent *types.Directory, path string, content string) *types.Query {
	return &types.Query{
		GenericData: g.newGenericData(path, "", parent),
		Content:     content,
	}
}

func (g *Generator) newDirectoryInstance(path string, name string, parent *types.Directory) *types.Directory {
	return &types.Directory{
		GenericData: g.newGenericData(path, name, parent),
		IsRoot:      g.opt.SourceDir == path,
		Directories: make([]*types.Directory, 0),
		Queries:     make([]*types.Query, 0),
	}
}

func (g *Generator) newGenericData(path string, name string, parent *types.Directory) types.GenericData {
	d := types.GenericData{
		Package:           g.opt.PackageName,
		Version:           version.GetVersion(),
		SourcePath:        path,
		QueryGetterSuffix: g.opt.QueryGetterSuffix,
	}

	if name == "" {
		name = filepath.Base(path)
	}

	for _, ext := range g.opt.SourceFilesExt {
		if s, ok := strings.CutSuffix(name, ext); ok {
			name = s

			break
		}
	}

	g.initIdentifiers(&d, name, parent)

	d.TargetPath = filepath.Join(g.opt.TargetDir, d.Identifier+".go")

	return d
}

func (g *Generator) initIdentifiers(
	data *types.GenericData,
	name string,
	parent *types.Directory,
) {
	packageName := g.opt.PackageName

	parts := nameToWords(name)
	if len(parts) == 0 {
		parts = nameToWords(packageName)
	}

	data.Identifier = wordsToIdentifier(parts)
	data.PublicSlug = wordsToCapitalized(parts)
	data.PrefixedPublicSlug = data.PublicSlug

	if parent != nil && (!parent.IsRoot || parent.Identifier == data.Identifier) {
		if prefixParts := nameToWords(parent.Identifier); len(prefixParts) > 0 {
			data.Identifier = wordsToIdentifier(prefixParts) + "-" + data.Identifier
			data.PrefixedPublicSlug = wordsToCapitalized(prefixParts) + data.PrefixedPublicSlug
		}
	}

	data.PrivateSlug = firstLower(data.PublicSlug)
}

func (g *Generator) format(ctx context.Context, content []byte) []byte {
	logger.Debugf("Formatting generated code...")

	formatted, err := g.formatter.Format(ctx, content)
	if err != nil {
		logger.Warningf("Failed to format generated code: %s", err.Error())

		return content
	}

	return formatted
}
