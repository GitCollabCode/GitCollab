package jwt

import (
	"net/http"
	"os"
	"time"

	goJwt "github.com/golang-jwt/jwt"
)

const ACCESS_TOKEN_STR = "access_token"
const REFRESH_TOKEN_STR = "refresh_token"

const (
	TOKEN_VERIFIED = iota
	TOKEN_INVALID
	TOKEN_EXPIRED
	NO_COOKIE
)

const ACCESS_TOKEN_EXP = 15 // minutes
const REFRESH_TOKEN_EXP = 6 // days

type AccessTokenContents struct {
	Role       string        // role of the user
	Authorized bool          // is user authorized
	UserName   string        // username, or email?
	Expiration time.Duration // time to expiration in minutes
}

type RefreshTokenContents struct {
	UserName   string
	Email      string
	Expiration time.Duration // time to expiration in days
}

type jwtTokenPair struct {
	AccessToken     *goJwt.Token
	AccessTokenErr  int
	RefreshToken    *goJwt.Token
	RefreshTokenErr int
}

// generate a new keypair
func NewTokenPair(ac *AccessTokenContents, rc *RefreshTokenContents) *jwtTokenPair {
	tokenPair := jwtTokenPair{
		AccessToken:     GenerateAccessToken(ac),
		RefreshToken:    GenerateRefreshToken(rc),
		AccessTokenErr:  TOKEN_VERIFIED,
		RefreshTokenErr: TOKEN_VERIFIED,
	}
	return &tokenPair
}

func getJWTFromCookie(cookieName string, r *http.Request) (*goJwt.Token, int) {
	jwtCookie, err := r.Cookie(cookieName)
	if err != nil {
		return nil, NO_COOKIE
	}
	token, err := goJwt.Parse(jwtCookie.Value, getEncodingSecret)
	if err != nil {
		if isTokenExpired(token) {
			return nil, TOKEN_EXPIRED
		}
		return nil, TOKEN_INVALID
	}
	return token, TOKEN_VERIFIED
}

func GenerateAccessToken(contents *AccessTokenContents) *goJwt.Token {
	// populate claims
	claims := goJwt.MapClaims{}
	claims["authorized"] = contents.Authorized
	claims["aud"] = contents.Role
	claims["userName"] = contents.UserName
	claims["exp"] = time.Now().Add(time.Minute * contents.Expiration).Unix()
	// generate token
	accessToken := goJwt.NewWithClaims(goJwt.SigningMethodHS256, claims)
	return accessToken
}

func GenerateRefreshToken(contents *RefreshTokenContents) *goJwt.Token {
	// populate claims
	claims := goJwt.MapClaims{}
	claims["userName"] = contents.UserName
	claims["exp"] = time.Now().Add(time.Hour * 24 * contents.Expiration).Unix()
	//generate token
	refreshToken := goJwt.New(goJwt.SigningMethodHS256)
	return refreshToken
}

/*
 * Encode a token using secret
 * return a string representation of the JWT
 */
func encodeTokenString(token *goJwt.Token, secret string) (string, error) {
	jwtSecret := []byte(secret)
	return token.SignedString(jwtSecret)
}

func ServejwtCookie(w http.ResponseWriter, name string, token *goJwt.Token) {
	tokenString, _ := encodeTokenString(token, os.Getenv("JWT_SECRET"))
	atCookie := http.Cookie{
		Name:    name,
		Path:    "/",
		Value:   tokenString,
		Expires: time.Now().Add(time.Hour * 168), // 1 week
	}
	http.SetCookie(w, &atCookie)
}

/*
 * Check if a token is expired
 * return true if token is expired, false if valid JWT
 */
func isTokenExpired(token *goJwt.Token) bool {
	claims := token.Claims.(goJwt.MapClaims)
	return !claims.VerifyExpiresAt(time.Now().Unix(), true)
}

func (kp jwtTokenPair) GetLatestTokenPair(r *http.Request) {
	kp.AccessToken, kp.AccessTokenErr = getJWTFromCookie(ACCESS_TOKEN_STR, r)
	kp.RefreshToken, kp.RefreshTokenErr = getJWTFromCookie(REFRESH_TOKEN_STR, r)
}

func (kp jwtTokenPair) GetAccessTokenContents() *AccessTokenContents {
	//claims, ok := kp.accessToken.Claims.(jwt.MapClaims)
	return nil
}

func (kp jwtTokenPair) GetRefreshTokenContents() *RefreshTokenContents {
	//claims, ok := kp.accessToken.Claims.(jwt.MapClaims)
	return nil
}

// Get secret used to initially encode the token
func getEncodingSecret(token *goJwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_SECRET")), nil
}

func (kp jwtTokenPair) ServeTokenPair(w http.ResponseWriter) {
	ServejwtCookie(w, ACCESS_TOKEN_STR, kp.AccessToken)
	ServejwtCookie(w, REFRESH_TOKEN_STR, kp.RefreshToken)
}

/*
 * Refresh access token
 */
func (kp jwtTokenPair) RefreshAccesstoken(w http.ResponseWriter, r *http.Request) int {
	if kp.RefreshToken.Valid || !isTokenExpired(kp.RefreshToken) {
		contents := AccessTokenContents{
			Authorized: true,
			Expiration: ACCESS_TOKEN_EXP,
			UserName:   "asd", // get from db
			Role:       "asd", // get from db
		}
		acessToken := GenerateAccessToken(&contents)
		ServejwtCookie(w, "refresh_token", acessToken)
		return TOKEN_VERIFIED
	}
	return NO_COOKIE
}
