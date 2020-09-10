package gqlbuilder

import (
	"io/ioutil"
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
			{
				Key:   "testFieldOne",
				Value: "1080",
			},
			{
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
			{
				Key:   "testFieldOne",
				Value: "3080",
			},
			{
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
			{
				Key:   "singleTestInput",
				Value: "980",
			},
		},
		Return: []string{"testReturnSingleValue"},
	})

	gotEncoded, err := testInput.Build()

	if err != nil {
		t.Errorf("(Mutation) Build() failed with error %s\n", err.Error())
	}

	got, _ := ioutil.ReadAll(gotEncoded)
	want := "{\"query\":\"mutation TestMutation {testAlias: testMutation(input: {testFieldOne: \\\"1080\\\", testFieldTwo: \\\"2080\\\"}) { testReturnId testReturnValue } testAnotherAlias: testMutation(input: {testFieldOne: \\\"3080\\\", testFieldTwo: \\\"3090\\\"}) { testReturnAnotherId testReturnAnotherValue } testNoAliasMutation(input: {singleTestInput: \\\"980\\\"}) { testReturnSingleValue } }\"}"

	if string(got) != want {
		t.Errorf("(Mutation) Build() returned:\n\"%s\"\n\nExpected:\n\"%s\"\n", got, want)
	}
}

func TestNamelessMutationBuilder(t *testing.T) {
	testInput := &Mutations{
		Name:      "TestMutation",
		Mutations: []Mutation{},
	}

	testInput.Mutations = append(testInput.Mutations, Mutation{
		Name:  "",
		Alias: "shouldFail",
		Inputs: []MutationInput{
			{
				Key:   "thisShouldFail",
				Value: "666",
			},
		},
		Return: []string{"thisShouldNotReturn"},
	})

	_, err := testInput.Build()
	want := "no mutation name specified"

	if err == nil {
		t.Errorf("(Mutation) Build() should have failed with error %s but did not fail\n", want)
	}

	got := err.Error()

	if got != want {
		t.Errorf("(Mutation) Build() failed with unexpected error\n\ngot %s\nwant %s\n", got, want)
	}
}
