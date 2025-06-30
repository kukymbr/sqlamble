package types

type Directory struct {
	GenericData

	IsRoot      bool
	Directories []*Directory
	Queries     []*Query
}
