package po

import (
	"testing"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/require"
)

type LoginPage struct {
	BasePage
	username  playwright.Locator
	password  playwright.Locator
	signInBtn playwright.Locator
}

func NewLoginPage(p playwright.Page, t *testing.T) *LoginPage {
	l := &LoginPage{BasePage: BasePage{page: p, t: *t, assertion: *require.New(t)}}
	l.username = l.page.GetByLabel("Username")
	l.password = l.page.GetByLabel("Password")
	l.signInBtn = l.page.GetByRole("button", playwright.PageGetByRoleOptions{Name: *playwright.String("Sign in")})
	return l
}

func (l *LoginPage) GoTo() *LoginPage {
	_, err := l.page.Goto("http://localhost:8080/")
	l.assertion.Nil(err)
	return l
}

func (l *LoginPage) ClickSignInLink() *LoginPage {
	err := l.page.GetByRole("link", playwright.PageGetByRoleOptions{Name: *playwright.String("Sign in ")}).Click()
	l.assertion.Nil(err)
	return l
}

func (l *LoginPage) FillUsername(username string) *LoginPage {
	err := l.username.Fill(username)
	l.assertion.Nil(err)
	return l
}

func (l *LoginPage) ClickPassword() *LoginPage {
	err := l.password.Click()
	l.assertion.Nil(err)
	return l
}

func (l *LoginPage) FillPassword(password string) *LoginPage {
	err := l.password.Fill(password)
	l.assertion.Nil(err)
	return l
}

func (l *LoginPage) ClickSignInButton() *WelcomePage {
	l.signInBtn.Click()
	return NewWelcomePage(l.page, &l.t)
}
