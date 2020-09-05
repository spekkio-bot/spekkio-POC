package sqlbuilder

import (
	"testing"
)

func TestSelectQueryBuilder(t *testing.T) {
	testProps := &SelectQueryProps{
		BaseTable: "TestTable",
		Columns: []Column{
			{
				Name:  "TestColumn",
				Alias: "",
			},
			{
				Name:  "TestColumnWithAlias",
				Alias: "AliasWorks",
			},
		},
	}

	got, err := testProps.BuildQuery()
	want := "SELECT TestColumn, TestColumnWithAlias AS AliasWorks FROM TestTable;"

	if err != nil {
		t.Errorf("Failed with error:\n%s", err.Error())
	} else if got != want {
		t.Errorf("BuildQuery() returned:\n\"%s\"\n\nExpected:\n\"%s\"\n", got, want)
	}
}

func TestSelectQueryBuilderEmptyColumnError(t *testing.T) {
	testProps := &SelectQueryProps{
		BaseTable: "TestTableEmptyColumn",
		Columns: []Column{
			{
				Name:  "",
				Alias: "",
			},
		},
	}

	_, err := testProps.BuildQuery()
	wantErr := "empty Column.name not allowed"

	if err.Error() != wantErr {
		t.Errorf("Failed with error:\n%s\n\nExpected error:\n%s\n", err.Error(), wantErr)
	} else if err == nil {
		t.Errorf("Did not fail with error, expected error:\n%s\n", wantErr)
	}
}
