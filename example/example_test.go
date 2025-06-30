package example_test

import (
	"fmt"
	"testing"

	"github.com/kukymbr/sqlamble/example/internal/queries"
)

func Test_PrintGeneratedQueries(t *testing.T) {
	fmt.Println(queries.Queries().VersionQuery())
	fmt.Println(queries.Queries().Users().GetListQuery())
	fmt.Println(queries.Queries().Users().SingleUser().GetUserDataQuery())
	fmt.Println(queries.Queries().Orders().ChangeStatusQuery())
	fmt.Println(queries.Queries().Orders().CreateOrderQuery())
}
