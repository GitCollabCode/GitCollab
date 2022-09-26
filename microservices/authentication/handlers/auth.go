package handlers

import "github.com/go-chi/jwtauth"

type Claim struct {
	claim string
	val string
}

func CreateToken(tokenAuth *jwtauth.JWTAuth, claims ...Claim) (error, string) {
	var c = map[string]interface{}{}
	for _, claimData := range claims {
		c[claimData.claim] = claimData.val
	}
	// below returns a token object, maybe use?
	_, tokenString, err := tokenAuth.Encode(c)
	return err, tokenString
}