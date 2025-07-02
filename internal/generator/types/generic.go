package types

type GenericData struct {
	Package string
	Version string

	SourcePath string
	TargetPath string

	Identifier         string
	PublicSlug         string
	PrefixedPublicSlug string
	PrivateSlug        string

	QueryGetterSuffix string
}
