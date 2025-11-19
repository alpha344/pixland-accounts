package user

import (
	"testing"

	"github.com/alpha344/pixland-accounts/server/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestAuthFrom(t *testing.T) {
	assert.Equal(t, Auth{
		Provider: "xx",
		Sub:      "xx|yy",
	}, AuthFrom("xx|yy"))

	assert.Equal(t, Auth{
		Provider: "",
		Sub:      "yy",
	}, AuthFrom("yy"))

	assert.Equal(t, Auth{
		Provider: "",
		Sub:      "",
	}, AuthFrom(""))
}

func TestNewPixlandAuth(t *testing.T) {
	assert.Equal(t, Auth{
		Provider: "pixland",
		Sub:      "pixland|xx|yy",
	}, NewPixlandAuth("xx|yy"))
}
func TestIsPixland(t *testing.T) {
	a := Auth{
		Provider: "pixland",
		Sub:      "yy",
	}
	assert.True(t, a.IsPixland())

	a = Auth{
		Provider: "xx",
		Sub:      "yy",
	}
	assert.False(t, a.IsPixland())
}

func TestIsAuth0(t *testing.T) {
	a := Auth{
		Provider: "auth0",
		Sub:      "yy",
	}
	assert.True(t, a.IsAuth0())

	a = Auth{
		Provider: "xx",
		Sub:      "yy",
	}
	assert.False(t, a.IsAuth0())
}

func TestPixlandSub(t *testing.T) {
	uid := NewID().String()

	assert.Equal(t, &Auth{
		Provider: "pixland",
		Sub:      "pixland|" + uid,
	}, PixlandSub(uid))
}

func TestAuth_Ref(t *testing.T) {
	a := Auth{
		Provider: "auth0",
		Sub:      "yy",
	}
	assert.Equal(t, &a, a.Ref())
}

func TestAuth_String(t *testing.T) {
	a := Auth{
		Provider: "auth0",
		Sub:      "yy",
	}
	assert.Equal(t, "yy", a.String())
}

func TestAuths_Has(t *testing.T) {
	a := []Auth{
		{
			Provider: "auth0",
			Sub:      "xxx",
		},
	}
	assert.True(t, Auths(a).Has("xxx"))
	assert.False(t, Auths(a).Has("yyy"))
}

func TestAuths_HasProvider(t *testing.T) {
	a := []Auth{
		{
			Provider: "auth0",
			Sub:      "xxx",
		},
	}
	assert.True(t, Auths(a).HasProvider("auth0"))
	assert.False(t, Auths(a).HasProvider("xxx"))
}

func TestAuths_GetByProvider(t *testing.T) {
	a := []Auth{
		{
			Provider: "auth0",
			Sub:      "xxx",
		},
	}
	assert.Equal(t, &Auth{
		Provider: "auth0",
		Sub:      "xxx",
	}, Auths(a).GetByProvider("auth0"))

	assert.Nil(t, Auths(a).GetByProvider("yyy"))
}

func TestAuths_Get(t *testing.T) {
	a := []Auth{
		{
			Provider: "auth0",
			Sub:      "xxx",
		},
	}
	assert.Equal(t, &Auth{
		Provider: "auth0",
		Sub:      "xxx",
	}, Auths(a).Get("xxx"))

	assert.Nil(t, Auths(a).Get("yyy"))
}

func TestAuths_Add(t *testing.T) {
	a := []Auth{
		{
			Provider: "auth0",
			Sub:      "xxx",
		},
	}
	assert.Equal(t, Auths([]Auth{
		{
			Provider: "auth0",
			Sub:      "xxx",
		},
		{
			Provider: "p",
			Sub:      "s",
		},
	}), Auths(a).Add(Auth{Provider: "p", Sub: "s"}))

	a = []Auth{
		{
			Provider: "auth0",
			Sub:      "xxx",
		},
	}
	assert.Equal(t, Auths([]Auth{{
		Provider: "auth0",
		Sub:      "xxx",
	}}), Auths(a).Add(Auth{
		Provider: "auth0",
		Sub:      "xxx",
	}))
}

func TestAuths_Remove(t *testing.T) {
	a := []Auth{
		{
			Provider: "auth0",
			Sub:      "xxx",
		},
	}
	assert.Equal(t, Auths([]Auth{}), Auths(a).Remove("xxx"))

	a = []Auth{
		{
			Provider: "auth0",
			Sub:      "xxx",
		},
	}
	assert.Equal(t, Auths([]Auth{{
		Provider: "auth0",
		Sub:      "xxx",
	}}), Auths(a).Remove("foo"))
}

func TestAuths_RemoveByProvider(t *testing.T) {
	a := []Auth{
		{
			Provider: "auth0",
			Sub:      "xxx",
		},
	}
	assert.Equal(t, Auths([]Auth{}), Auths(a).RemoveByProvider("auth0"))

	a = []Auth{
		{
			Provider: "auth0",
			Sub:      "xxx",
		},
	}
	assert.Equal(t, Auths([]Auth{{
		Provider: "auth0",
		Sub:      "xxx",
	}}), Auths(a).RemoveByProvider("foo"))
}

func TestGenPixlandSub(t *testing.T) {
	uid := id.NewUserID()

	tests := []struct {
		name  string
		input string
		want  *Auth
	}{
		{
			name:  "should return pixland sub",
			input: uid.String(),
			want: &Auth{
				Provider: "pixland",
				Sub:      "pixland|" + uid.String(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenPixlandSub(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
