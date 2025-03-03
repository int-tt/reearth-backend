package gqlmodel

import (
	"github.com/reearth/reearth-backend/pkg/user"
)

func ToUser(user *user.User) *User {
	if user == nil {
		return nil
	}
	auths := user.Auths()
	authsgql := make([]string, 0, len(auths))
	for _, a := range auths {
		authsgql = append(authsgql, a.Provider)
	}
	return &User{
		ID:       user.ID().ID(),
		Name:     user.Name(),
		Email:    user.Email(),
		Lang:     user.Lang(),
		Theme:    Theme(user.Theme()),
		MyTeamID: user.Team().ID(),
		Auths:    authsgql,
	}
}

func ToSearchedUser(u *user.User) *SearchedUser {
	if u == nil {
		return nil
	}
	return &SearchedUser{
		UserID:    u.ID().ID(),
		UserName:  u.Name(),
		UserEmail: u.Email(),
	}
}

func ToTheme(t *Theme) *user.Theme {
	th := user.ThemeDefault

	if t == nil {
		return nil
	}

	switch *t {
	case ThemeDark:
		th = user.ThemeDark
	case ThemeLight:
		th = user.ThemeLight
	}
	return &th
}
