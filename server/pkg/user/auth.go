package user

import (
	"slices"
	"strings"

	"github.com/samber/lo"
)

const (
	ProviderPixland = "pixland"
	ProviderAuth0   = "auth0"
)

type Auth struct {
	Provider string
	Sub      string
}

func AuthFrom(sub string) Auth {
	s := strings.SplitN(sub, "|", 2)
	if len(s) != 2 {
		return Auth{Provider: "", Sub: sub}
	}
	return Auth{Provider: s[0], Sub: sub}
}

func NewPixlandAuth(sub string) Auth {
	return Auth{
		Provider: ProviderPixland,
		Sub:      "pixland|" + sub,
	}
}

func (a Auth) IsPixland() bool {
	return a.Provider == ProviderPixland
}

func (a Auth) IsAuth0() bool {
	return a.Provider == ProviderAuth0
}

func (a Auth) Ref() *Auth {
	a2 := a
	return &a2
}

func (a Auth) String() string {
	return a.Sub
}

func GenPixlandSub(userID string) *Auth {
	return &Auth{
		Provider: "pixland",
		Sub:      "pixland|" + userID,
	}
}

type Auths []Auth

func (a Auths) Has(sub string) bool {
	return lo.ContainsBy(a, func(a Auth) bool { return a.Sub == sub })
}

func (a Auths) HasProvider(p string) bool {
	return lo.ContainsBy(a, func(a Auth) bool { return a.Provider == p })
}

func (a Auths) GetByProvider(p string) *Auth {
	_, i, ok := lo.FindIndexOf(a, func(a Auth) bool { return a.Provider == p })
	if !ok {
		return nil
	}
	return &a[i]
}

func (a Auths) Get(sub string) *Auth {
	_, i, ok := lo.FindIndexOf(a, func(a Auth) bool { return a.Sub == sub })
	if !ok {
		return nil
	}
	return &a[i]
}

func (a Auths) Add(u Auth) Auths {
	if a.Has(u.Sub) {
		return a
	}
	return append(a, u)
}

func (a Auths) Remove(sub string) Auths {
	_, i, ok := lo.FindIndexOf(a, func(a Auth) bool { return a.Sub == sub })
	if !ok {
		return a
	}
	return slices.Delete(a, i, 1)
}

func (a Auths) RemoveByProvider(p string) Auths {
	_, i, ok := lo.FindIndexOf(a, func(a Auth) bool { return a.Provider == p })
	if !ok {
		return a
	}
	return slices.Delete(a, i, 1)
}

func PixlandSub(userID string) *Auth {
	return &Auth{
		Provider: "pixland",
		Sub:      "pixland|" + userID,
	}
}
