package types

import "github.com/kukymbr/sqlamble/internal/utils"

type Query struct {
	GenericData

	Content string
}

func (q *Query) GetQuotedContent() string {
	return utils.GetQuotedContent(q.Content)
}
