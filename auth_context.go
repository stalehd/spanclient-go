package spanclient

import "context"

// NewAuthContext is a convenience method to create an authentication context
// with an API key.
func NewAuthContext(apiToken string) context.Context {
	return context.WithValue(context.Background(),
		ContextAPIKey,
		APIKey{
			Key:    apiToken,
			Prefix: "",
		})
}
