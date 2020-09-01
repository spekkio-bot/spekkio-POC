package gqlbuilder

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/spekkio-bot/spekkio/src/app/model"
)

// Operation is the top level GraphQL query / mutation structure
type Operation struct {
	Name       string
	Type       string // "query" or "mutation"
	Operations []interface{}
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

// Build will call BuildQuery and then return the JSON encoding of the GraphQL query
func (o *Operation) Build() (io.Reader, error) {
	query, err := o.BuildQuery()
	if err != nil {
		var nilSlice []byte
		return bytes.NewBuffer(nilSlice), err
	}
	req := model.GraphQLRequest{
		Query: query,
	}

	var encoded []byte
	encoded, err = json.Marshal(req)
	if err != nil {
		var nilSlice []byte
		return bytes.NewBuffer(nilSlice), err
	}

	return bytes.NewBuffer(encoded), nil
}

// BuildQuery takes the fields of an Operation and builds a GraphQL request string
func (o *Operation) BuildQuery() (string, error) {
	var req string

	switch o.Type {
	case "query":
		req = fmt.Sprintf("query %s {", o.Name)
	case "mutation":
		req = fmt.Sprintf("mutation %s {", o.Name)
	default:
		err := errors.New("invalid gqlbuilder operation type")
		return "", err
	}

	for _, operation := range o.Operations {
		fragment := ""
		switch operation.(type) {
		case Mutation:
			if operation.(Mutation).Name == "" {
				err := errors.New("no operation name specified")
				return "", err
			}
			if operation.(Mutation).Alias != "" {
				fragment = fmt.Sprintf("%s: ", operation.(Mutation).Alias)
			}
			fragment += operation.(Mutation).Name + "(input: {"
			for i, input := range operation.(Mutation).Inputs {
				if i > 0 {
					fragment += ", "
				}
				fragment += input.Key + ": "
				fragment += "\"" + input.Value + "\""
			}
			fragment += "}) { "
			for _, item := range operation.(Mutation).Return {
				fragment += item + " "
			}
			fragment += "} "
		default:
			err := errors.New("invalid operation type")
			return "", err
		}
		req += fragment
	}

	req += "}"

	return req, nil
}
