package model

// Ping dictates how the response to the index route of the app should be structured.
type Ping struct {
	Message string `json:"message"`
}

// GraphQLRequest dictates how GraphQL requests should be structured (duh).
type GraphQLRequest struct {
	Query string `json:"query"`
}

// ScrumifyRequest dictates how the request body of the Scrumify resource should be structured.
type ScrumifyRequest struct {
	Token  string `json:"token"`
	RepoID string `json:"repo_id"`
}

// ScrumifyResponse dictates how the response body of the Scrumify resource should be structured.
type ScrumifyResponse struct {
	Message  string `json:"message"`
	Response string `json:"github_response"`
}

// ScrumifyQueryResponse dictates how the ScrumifyLabels query response should be structured.
type ScrumifyQueryResponse struct {
	ID          int
	Name        string
	Color       string
	Description string
}

// Error dictates how the app's error response should be structured.
type Error struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}
