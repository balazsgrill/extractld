package bc3

import (
	"context"
	"net/http"

	"github.com/balazsgrill/basecamp3"
	"github.com/balazsgrill/oauthenticator"
)

type tokenPersistence struct {
	context.Context
	oauthenticator.TokenPersistence
}

var _ basecamp3.ContextWithTokenPersistence = &tokenPersistence{}

func (tp *tokenPersistence) get(w http.ResponseWriter, r *http.Request) basecamp3.ContextWithTokenPersistence {
	return tp
}
