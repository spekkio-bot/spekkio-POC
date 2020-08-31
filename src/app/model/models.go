package model

// Ping dictates how the response to the index route of the app should be structured.
type Ping struct {
	Message string `json:"message"`
}

// ScrumifyRequest dictates how the request body of the Scrumify resource should be structured.
type ScrumifyRequest struct {
	Token  string `json:"token"`
	RepoID string `json:"repo_id"`
}

// Error dictates how the app's error response should be structured.
type Error struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}
