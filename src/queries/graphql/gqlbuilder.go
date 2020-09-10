package gqlbuilder

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/spekkio-bot/spekkio/src/app/model"
)

// Mutations is the top level GraphQL mutation structure
type Mutations struct {
	Name      string
	Mutations []Mutation
}

// Mutation is a structure for mutations
type Mutation struct {
	Name   string
	Alias  string
	Inputs []MutationInput
	Return []string
}

// MutationInput is a structure for the Inputs field of Mutation
type MutationInput struct {
	Key   string
	Value string
}

// Build will call BuildQuery and then return the JSON encoding of the GraphQL mutation
func (m *Mutations) Build() (io.Reader, error) {
	query, err := m.BuildQuery()
	if err != nil {
		var nilSlice []byte
		return bytes.NewBuffer(nilSlice), err
	}
	req := model.GraphQLRequest{
		Query: query,
	}

	var encoded []byte
	encoded, _ = json.Marshal(req)
	return bytes.NewBuffer(encoded), nil
}

// BuildQuery takes the fields of an Operation and builds a GraphQL mutation request string
func (m *Mutations) BuildQuery() (string, error) {
	req := fmt.Sprintf("mutation %s {", m.Name)
	for _, mutation := range m.Mutations {
		fragment := ""
		if mutation.Name == "" {
			err := errors.New("no mutation name specified")
			return "", err
		}
		if mutation.Alias != "" {
			fragment = fmt.Sprintf("%s: ", mutation.Alias)
		}
		fragment += mutation.Name + "(input: {"
		for i, input := range mutation.Inputs {
			if i > 0 {
				fragment += ", "
			}
			fragment += input.Key + ": "
			fragment += "\"" + input.Value + "\""
		}
		fragment += "}) { "
		for _, item := range mutation.Return {
			fragment += item + " "
		}
		fragment += "} "
		req += fragment
	}
	req += "}"

	return req, nil
}
