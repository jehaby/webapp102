package auth

import (
	"context"
	"encoding/json"
	"time"

	"github.com/davecgh/go-spew/spew"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"

	"github.com/jehaby/webapp102/config"
	"github.com/jehaby/webapp102/entity"
)

var (
	StrTokenCtxKey = &contextKey{"jwt"}
	TokenCtxKey    = &contextKey{"token"}
	ErrorCtxKey    = &contextKey{"error"}
)

var (
	ErrUnauthorized = errors.New("jwtauth: token is unauthorized")
	ErrExpired      = errors.New("jwtauth: token is expired")
)

type JwtAuth struct {
	signKey   interface{}
	verifyKey interface{}
	signer    jwt.SigningMethod
	parser    *jwt.Parser
}

// New creates a JwtAuth authenticator instance that provides middleware handlers
// and encoding/decoding functions for JWT signing.
func New(cfg config.Auth) *JwtAuth {
	return NewWithParser(cfg.Alg, &jwt.Parser{}, []byte(cfg.Secret), nil)
}

// NewWithParser is the same as New, except it supports custom parser settings
// introduced in jwt-go/v2.4.0.
//
// We explicitly toggle `SkipClaimsValidation` in the `jwt-go` parser so that
// we can control when the claims are validated - in our case, by the Verifier
// http middleware handler.
func NewWithParser(alg string, parser *jwt.Parser, signKey interface{}, verifyKey interface{}) *JwtAuth {
	parser.SkipClaimsValidation = true
	return &JwtAuth{
		signKey:   signKey,
		verifyKey: verifyKey,
		signer:    jwt.GetSigningMethod(alg),
		parser:    parser,
	}
}

// Verify
func (ja *JwtAuth) Verify(ctx context.Context) (*jwt.Token, error) {
	strToken, err := tknFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	// Verify the token
	token, err := ja.parser.Parse(strToken, ja.keyFunc)
	if err != nil {
		switch err.Error() {
		case "token is expired":
			err = ErrExpired
		}
		return token, err
	}
	if token == nil || !token.Valid || token.Method != ja.signer {
		err = ErrUnauthorized
		return token, err
	}

	// Check expiry via "exp" claim
	if IsExpired(token) {
		err = ErrExpired
		return token, err
	}

	spew.Dump("in auth service", token)

	// Valid!
	return token, nil
}

func tknFromCtx(ctx context.Context) (string, error) {
	tknRow := ctx.Value(StrTokenCtxKey)
	if tknRow == nil {
		return "", ErrUnauthorized
	}
	if strToken, _ := tknRow.(string); strToken != "" {
		return strToken, nil
	}
	return "", ErrUnauthorized
}

func IsExpired(t *jwt.Token) bool {
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		panic("jwtauth: expecting jwt.MapClaims")
	}

	if expv, ok := claims["exp"]; ok {
		var exp int64
		switch v := expv.(type) {
		case float64:
			exp = int64(v)
		case int64:
			exp = v
		case json.Number:
			exp, _ = v.Int64()
		default:
		}

		if exp < epochNow() {
			return true
		}
	}

	return false
}

func (ja *JwtAuth) TokenFromUser(user *entity.User, expiryTime time.Duration) (string, error) {
	claims := Claims{
		"user": userResponseFormUser(*user),
	}.SetExpiryIn(expiryTime)

	_, tkn, err := ja.Encode(claims)
	if err != nil {
		return "", errors.Wrapf(err, "auth.TokenFromUser: couldn't Encode claims (%v)", claims)
	}
	return tkn, nil
}

type userResponse struct {
	UUID  uuid.UUID `json:"uuid"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

func userResponseFormUser(u entity.User) userResponse {
	return userResponse{
		UUID:  u.UUID,
		Name:  u.Name,
		Email: u.Email,
	}
}

func (ja *JwtAuth) Encode(claims Claims) (t *jwt.Token, tokenString string, err error) {
	t = jwt.New(ja.signer)
	t.Claims = claims
	tokenString, err = t.SignedString(ja.signKey)
	t.Raw = tokenString
	return
}

// Claims is a convenience type to manage a JWT claims hash.
type Claims map[string]interface{}

// NOTE: as of v3.0 of jwt-go, Valid() interface method is called to verify
// the claims. However, the current design we test these claims in the
// Verifier middleware, so we skip this step.
func (c Claims) Valid() error {
	return nil
}

func (c Claims) Set(k string, v interface{}) Claims {
	c[k] = v
	return c
}

func (c Claims) Get(k string) (interface{}, bool) {
	v, ok := c[k]
	return v, ok
}

// Set issued at ("iat") to specified time in the claims
func (c Claims) SetIssuedAt(tm time.Time) Claims {
	c["iat"] = tm.UTC().Unix()
	return c
}

// Set issued at ("iat") to present time in the claims
func (c Claims) SetIssuedNow() Claims {
	c["iat"] = EpochNow()
	return c
}

// Set expiry ("exp") in the claims and return itself so it can be chained
func (c Claims) SetExpiry(tm time.Time) Claims {
	c["exp"] = tm.UTC().Unix()
	return c
}

// Set expiry ("exp") in the claims to some duration from the present time
// and return itself so it can be chained
func (c Claims) SetExpiryIn(tm time.Duration) Claims {
	c["exp"] = ExpireIn(tm)
	return c
}

// Helper function that returns the NumericDate time value used by the spec
func EpochNow() int64 {
	return time.Now().UTC().Unix()
}

// Helper function to return calculated time in the future for "exp" claim.
func ExpireIn(tm time.Duration) int64 {
	return EpochNow() + int64(tm.Seconds())
}

// Helper function that returns the NumericDate time value used by the spec
func epochNow() int64 {
	return time.Now().UTC().Unix()
}

func (ja *JwtAuth) keyFunc(t *jwt.Token) (interface{}, error) {
	if ja.verifyKey != nil {
		return ja.verifyKey, nil
	} else {
		return ja.signKey, nil
	}
}

// a pointer so it fits in an interface{} without allocation. This technique
// for defining context keys was copied from Go 1.7's new use of context in net/http.
type contextKey struct {
	name string
}

func (k *contextKey) String() string {
	return "jwtauth context value " + k.name
}
