package response

// Dummy is a dummy response to test client-server communication
type Dummy struct {
	Message string `json:"message"`
}
