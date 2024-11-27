package main

import (
	"context"

	"github.com/coreos/go-oidc/v3/oidc"
)

var token string = `{
  "aud": "2431dd52-...",
  "iss": "https://login.microsoftonline.com/374f8026-.../v2.0",
}`

func main() {
	p, err := oidc.NewProvider(context.Background(), "https://login.microsoftonline.com/374f8026-.../v2.0")
	if err != nil {
		panic(err)
	}

	verifier := p.Verifier(&oidc.Config{
		ClientID: "2431dd52-...", // the expected audience
	})

	tk, err := verifier.Verify(context.Background(), token)
	if err != nil {
		panic(err)
	}

	println(tk)
}
