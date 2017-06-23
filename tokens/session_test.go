package tokens_test

import (
	"net/url"
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/keratin/authn/config"
	"github.com/keratin/authn/data/mock"
	"github.com/keratin/authn/tests"
	"github.com/keratin/authn/tokens"
)

func TestSessionJWT(t *testing.T) {
	store := mock.NewRefreshTokenStore()
	cfg := config.Config{
		AuthNURL:          &url.URL{Scheme: "http", Host: "authn.example.com"},
		SessionSigningKey: []byte("key-a-reno"),
	}

	session, err := tokens.NewSessionJWT(store, &cfg, 658908)
	if err != nil {
		t.Fatal(err)
	}
	tests.AssertEqual(t, "http://authn.example.com", session.Issuer)
	tests.AssertEqual(t, "http://authn.example.com", session.Audience)
	tests.AssertEqual(t, "RefreshToken:658908", session.Subject)
	tests.AssertEqual(t, "", session.Azp)
	tests.RefuteEqual(t, int64(0), session.IssuedAt)

	sessionString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, session).SignedString(cfg.SessionSigningKey)
	if err != nil {
		t.Fatal(err)
	}

	claims, err := tokens.ParseSessionJWT(sessionString, cfg.SessionSigningKey)
	if err != nil {
		t.Fatal(err)
	}
	tests.AssertEqual(t, "http://authn.example.com", claims.Issuer)
	tests.AssertEqual(t, "http://authn.example.com", claims.Audience)
	tests.AssertEqual(t, "RefreshToken:658908", claims.Subject)
	tests.AssertEqual(t, "", claims.Azp)
	tests.RefuteEqual(t, int64(0), claims.IssuedAt)
}