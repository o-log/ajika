package ajika

import (
	"net/http"
	"testing"

	"github.com/o-log/ajika/wraperr"
)

func TestSSRFNegative(t *testing.T) {
	ajika := Ajika{}
	_, err := ajika.Request(MockSalt{}, 0, http.MethodGet, "https://google.com", "", nil)
	if err == nil {
		t.Fatal("no error, should be")
	}
}

func TestSSRFPositive(t *testing.T) {
	ajika := Ajika{
		AllowedDomains: []string{
			"google.com",
			"policies.google.com",
		},
	}
	_, err := ajika.Request(MockSalt{}, 0, http.MethodGet, "https://google.com", "", nil)
	if err != nil {
		t.Fatal(wraperr.Wrap(err))
	}

	_, err = ajika.Request(MockSalt{}, 0, http.MethodGet, "https://policies.google.com/privacy", "", nil)
	if err != nil {
		t.Fatal(wraperr.Wrap(err))
	}
}
