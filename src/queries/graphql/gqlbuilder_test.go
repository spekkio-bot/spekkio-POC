package gqlbuilder

import (
	"testing"
)

func TestMutationBuilder(t *testing.T) {
	testInput := &Operation{
		Name:       "TestMutation",
		Type:       "mutation",
		Operations: make([]interface{}, 3),
	}

	testInput.Operations[0] = Mutation{
		Name:   "testMutation",
		Alias:  "testAlias",
		Inputs: []MutationInput{
			MutationInput{
				Key:   "testFieldOne",
				Value: "1080",
			},
			MutationInput{
				Key:   "testFieldTwo",
				Value: "2080",
			},
		},
		Return: []string{"testReturnId", "testReturnValue"},
	}

	testInput.Operations[1] = Mutation{
		Name:   "testMutation",
		Alias:  "testAnotherAlias",
		Inputs: []MutationInput{
			MutationInput{
				Key:   "testFieldOne",
				Value: "3080",
			},
			MutationInput{
				Key:   "testFieldTwo",
				Value: "3090",
			},
		},
		Return: []string{"testReturnAnotherId", "testReturnAnotherValue"},
	}

	testInput.Operations[2] = Mutation{
		Name:   "testNoAliasMutation",
		Alias:  "",
		Inputs: []MutationInput{
			MutationInput{
				Key:   "singleTestInput",
				Value: "980",
			},
		},
		Return: []string{"testReturnSingleValue"},
	}

	got, err := testInput.BuildQuery()
	want := "mutation TestMutation {testAlias: testMutation(input: {testFieldOne: \"1080\", testFieldTwo: \"2080\"}) { testReturnId testReturnValue } testAnotherAlias: testMutation(input: {testFieldOne: \"3080\", testFieldTwo: \"3090\"}) { testReturnAnotherId testReturnAnotherValue } testNoAliasMutation(input: {singleTestInput: \"980\"}) { testReturnSingleValue } }"

	if err != nil {
		t.Errorf("Failed with error %s", err.Error())
	} else if got != want {
		t.Errorf("BuildQuery() returned:\n\"%s\"\n\nExpected:\n\"%s\"\n", got, want)
	}
}

func TestInvalidOperationType(t *testing.T) {
	testInput := &Operation{
		Name:       "TestInvalid",
		Type:       "invalid",
		Operations: make([]interface{}, 0),
	}

	got, err := testInput.BuildQuery()
	wantErr := "invalid gqlbuilder operation type"

	if err.Error() != wantErr {
		t.Errorf("BuildQuery() returned error:\n\"%s\"\nExpected error:\n\"%s\"\n", err.Error(), wantErr)
	}
	if got != "" {
		t.Errorf("BuildQuery() returned:\n\"%s\"\n\nExpected:\n\"\"\n", got)
	}
}
