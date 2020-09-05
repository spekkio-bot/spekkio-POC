package gqlbuilder

import (
	"testing"
)

func TestMutationBuilder(t *testing.T) {
	testInput := &Mutations{
		Name:      "TestMutation",
		Mutations: []Mutation{},
	}

	testInput.Mutations = append(testInput.Mutations, Mutation{
		Name:  "testMutation",
		Alias: "testAlias",
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
	})

	testInput.Mutations = append(testInput.Mutations, Mutation{
		Name:  "testMutation",
		Alias: "testAnotherAlias",
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
	})

	testInput.Mutations = append(testInput.Mutations, Mutation{
		Name:  "testNoAliasMutation",
		Alias: "",
		Inputs: []MutationInput{
			MutationInput{
				Key:   "singleTestInput",
				Value: "980",
			},
		},
		Return: []string{"testReturnSingleValue"},
	})

	got, err := testInput.BuildQuery()
	want := "mutation TestMutation {testAlias: testMutation(input: {testFieldOne: \"1080\", testFieldTwo: \"2080\"}) { testReturnId testReturnValue } testAnotherAlias: testMutation(input: {testFieldOne: \"3080\", testFieldTwo: \"3090\"}) { testReturnAnotherId testReturnAnotherValue } testNoAliasMutation(input: {singleTestInput: \"980\"}) { testReturnSingleValue } }"

	if err != nil {
		t.Errorf("Failed with error %s", err.Error())
	} else if got != want {
		t.Errorf("BuildQuery() returned:\n\"%s\"\n\nExpected:\n\"%s\"\n", got, want)
	}
}
