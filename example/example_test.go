package example_test

import (
	"fmt"
	"testing"

	"github.com/kukymbr/sqlamble/example/internal/queries"
)

func Test_PrintGeneratedQueries(t *testing.T) {
	fmt.Println(queries.VersionQuery())
	fmt.Println(queries.Users().GetListQuery())
	fmt.Println(queries.Users().SingleUser().GetUserDataQuery())
	fmt.Println(queries.Orders().ChangeStatusQuery())
	fmt.Println(queries.Orders().CreateOrderQuery())
}
