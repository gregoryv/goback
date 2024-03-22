package http

// ServeMux is an HTTP request multiplexer. It matches the URL of each
// incoming request against a list of registered patterns and calls
// the handler for the pattern that most closely matches the URL.
type ServeMux struct {
	// Has unexported fields.
}
