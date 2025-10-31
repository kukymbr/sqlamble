package generator_test

import (
	"context"
	"io"
	"io/fs"
	"os"
	"testing"

	"github.com/kukymbr/sqlamble/internal/formatter"
	"github.com/kukymbr/sqlamble/internal/generator"
	"github.com/stretchr/testify/suite"
)

type generatorGenerateTestCase struct {
	Name                  string
	GetOptFunc            func() generator.Options
	GetContextFunc        func() context.Context
	AssertConstructorFunc func(err error)
	AssertFunc            func(targetFS fs.FS, err error)
}

func TestGenerator(t *testing.T) {
	suite.Run(t, &GeneratorSuite{})
}

type GeneratorSuite struct {
	suite.Suite

	targets map[string]struct{}
}

func (s *GeneratorSuite) SetupSuite() {
	s.targets = make(map[string]struct{})
}

func (s *GeneratorSuite) TearDownSuite() {
	for path := range s.targets {
		_ = os.RemoveAll(path)
	}
}

func (s *GeneratorSuite) TestGenerator_PositiveCases() {
	tests := []generatorGenerateTestCase{
		{
			Name: "sql",
			GetOptFunc: func() generator.Options {
				return generator.Options{
					SourceDir: "testdata/source/sql",
				}
			},
			AssertConstructorFunc: func(err error) {
				s.Require().NoError(err)
			},
			AssertFunc: func(targetFS fs.FS, err error) {
				s.Require().NoError(err)
				s.Require().NotNil(targetFS)
				s.assertFile(
					targetFS,
					"queries.go",
					"package queries",
					"VersionQuery()",
					"SELECT version FROM app_info;",
					"Users()",
				)
				s.assertFile(
					targetFS,
					"users.go",
					"GetListQuery()",
					"SingleUser()",
				)
				s.assertFile(
					targetFS,
					"users-single-user.go",
					"GetUserDataQuery()",
					"SELECT * FROM users WHERE id = $1;",
				)
				s.assertNoFile(targetFS, "users-ignored.go")
			},
		},
		{
			Name: "yaml",
			GetOptFunc: func() generator.Options {
				return generator.Options{
					PackageName:       "configs",
					SourceDir:         "testdata/source/yaml",
					SourceFilesExt:    []string{".yml", ".yaml"},
					QueryGetterSuffix: "YAML",
				}
			},
			AssertConstructorFunc: func(err error) {
				s.Require().NoError(err)
			},
			AssertFunc: func(targetFS fs.FS, err error) {
				s.Require().NoError(err)
				s.Require().NotNil(targetFS)
				s.assertFile(
					targetFS,
					"configs.go",
					"package configs",
					"ConfigYAML()",
					"descriptions: Yes, in fact, you can use sqlamble to embed any type of content info the go code.",
					"Nested()",
				)
				s.assertFile(
					targetFS,
					"nested.go",
					"OtherYAML() string",
				)
			},
		},
	}

	for _, test := range tests {
		s.Run(test.Name, func() {
			s.runGeneratorGenerateTest(test)
		})
	}
}

func (s *GeneratorSuite) TestGenerator_NegativeCases() {
	tests := []generatorGenerateTestCase{
		{
			Name: "when source dir does not exist",
			GetOptFunc: func() generator.Options {
				return generator.Options{
					SourceDir: "testdata/unknown_dir",
				}
			},
			AssertConstructorFunc: func(err error) {
				s.Require().Error(err)
			},
		},
		{
			Name: "when query getter suffix is invalid",
			GetOptFunc: func() generator.Options {
				return generator.Options{
					SourceDir:         "testdata/source/sql",
					QueryGetterSuffix: "- -",
				}
			},
			AssertConstructorFunc: func(err error) {
				s.Require().Error(err)
			},
		},
		{
			Name: "when formatter is invalid",
			GetOptFunc: func() generator.Options {
				return generator.Options{
					SourceDir: "testdata/source/sql",
					Formatter: "unknown formatter",
				}
			},
			AssertConstructorFunc: func(err error) {
				s.Require().Error(err)
			},
		},
		{
			Name: "when context is canceled",
			GetContextFunc: func() context.Context {
				ctx, cancel := context.WithCancel(context.Background())
				cancel()

				return ctx
			},
			GetOptFunc: func() generator.Options {
				return generator.Options{
					SourceDir: "testdata/source/sql",
				}
			},
			AssertConstructorFunc: func(err error) {
				s.Require().NoError(err)
			},
			AssertFunc: func(targetFS fs.FS, err error) {
				s.Require().ErrorIs(err, context.Canceled)
			},
		},
	}

	for _, test := range tests {
		s.Run(test.Name, func() {
			s.runGeneratorGenerateTest(test)
		})
	}
}

func (s *GeneratorSuite) runGeneratorGenerateTest(test generatorGenerateTestCase) {
	targetPath := s.prepareTarget(test.Name)
	targetFS := os.DirFS(targetPath)

	opt := test.GetOptFunc()
	opt.TargetDir = targetPath

	if opt.Formatter == "" {
		opt.Formatter = formatter.Noop
	}

	gen, err := generator.New(opt)
	test.AssertConstructorFunc(err)

	if err != nil {
		return
	}

	ctx := s.T().Context()
	if test.GetContextFunc != nil {
		ctx = test.GetContextFunc()
	}

	err = gen.Generate(ctx)
	if test.AssertFunc != nil {
		test.AssertFunc(targetFS, err)
	}
}

func (s *GeneratorSuite) assertFile(fsys fs.FS, path string, expectContains ...string) {
	s.T().Helper()

	statFS := fsys.(fs.StatFS)

	stat, err := statFS.Stat(path)

	s.Require().NoError(err)
	s.Require().NotNil(stat)
	s.Require().False(stat.IsDir(), "expected file, got a directory")

	if len(expectContains) == 0 {
		return
	}

	f, err := fsys.Open(path)
	s.Require().NoError(err)

	content, err := io.ReadAll(f)
	s.Require().NoError(err)

	for _, substr := range expectContains {
		s.Require().Contains(string(content), substr)
	}
}

func (s *GeneratorSuite) assertNoFile(fsys fs.FS, path string) {
	s.T().Helper()

	statFS := fsys.(fs.StatFS)
	_, err := statFS.Stat(path)

	s.Require().True(os.IsNotExist(err), "expected file to not exist")
}

func (s *GeneratorSuite) prepareTarget(name string) string {
	s.T().Helper()

	targetPath := "testdata/target/" + name

	_ = os.RemoveAll(targetPath)
	s.Require().NoError(os.MkdirAll(targetPath, 0755))

	s.targets[targetPath] = struct{}{}

	return targetPath
}
