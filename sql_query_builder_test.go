package mysql

import (
	"fmt"
	"testing"
)

const TABLE_NAME = "demo"

var sqb SqlQueryBuilder

func TestSQBInsert(t *testing.T) {
	sqb.Insert(TABLE_NAME, "id", "add_time", "edit_time", "name")

	printQueryAndArgs()
}

func TestSQBValues(t *testing.T) {
	sqb.Values(
		[]interface{}{1, "2016-06-23 09:00:00", "2016-06-23 09:00:00", "a"},
		[]interface{}{2, "2016-06-23 09:10:00", "2016-06-23 09:10:00", "b"},
	)

	printQueryAndArgs()
}

func TestSQBDelete(t *testing.T) {
	sqb.Delete(TABLE_NAME)

	printQueryAndArgs()
}

func TestSQBUpdate(t *testing.T) {
	sqb.Update(TABLE_NAME)

	printQueryAndArgs()
}

func TestSQBSet(t *testing.T) {
	sqb.Set(
		NewSqlColQueryItem("name", "", "d"),
		NewSqlColQueryItem("edit_time", "", "2016-06-24 09:00:00"),
	)

	printQueryAndArgs()
}

func TestSQBSelect(t *testing.T) {
	sqb.Select("*", TABLE_NAME)
	printQueryAndArgs()

	sqb.Select("name, count(*)", TABLE_NAME)
	printQueryAndArgs()
}

func TestSQBWhere(t *testing.T) {
	sqb.WhereConditionAnd(
		//NewSqlColQueryItem("id", SQL_COND_IN, []int64{1, 2}),
		//NewSqlColQueryItem("add_time", SQL_COND_BETWEEN, []string{"2016-06-23 00:00:00", "2016-06-25 00:00:00"}),
		NewSqlColQueryItem("edit_time", SQL_COND_LESS_EQUAL, "2016-06-24 09:00:00"),
		//NewSqlColQueryItem("name", SQL_COND_LIKE, "%a%"),
	)
	printQueryAndArgs()
}

func TestSQBGroupBy(t *testing.T) {
	sqb.GroupBy("name ASC")
	printQueryAndArgs()
}

func TestSQBHaving(t *testing.T) {
	sqb.HavingConditionAnd(
		NewSqlColQueryItem("id", SQL_COND_GREATER, 3),
	)
	printQueryAndArgs()
}

func TestSQBOrderBy(t *testing.T) {
	sqb.OrderBy("id DESC")
	printQueryAndArgs()
}

func TestSQBLimit(t *testing.T) {
	sqb.Limit(0, 10)
	printQueryAndArgs()
}

func printQueryAndArgs() {
	fmt.Println(sqb.Query(), sqb.Args())
}
