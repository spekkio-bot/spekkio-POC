package model

// Ping dictates how the response to the index route of the app should be structured.
type Ping struct {
	Message string `json:"message"`
}

// Error dictates how the app's error response should be structured.
type Error struct {
	Message string `json:"message"`
}
