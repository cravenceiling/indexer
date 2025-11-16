package document

import (
	"github.com/cravenceiling/indexer/cli/internal/email"
)

// A Document contains the path of the email and the email itself.
type Document struct {
	// Path is the path of the email file.
	Path string `json:"path"`
	// Email is the email itself.
	Email *email.Email `json:"email"`
}
