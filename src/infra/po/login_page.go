package po

import (
	"testing"

	"github.com/playwright-community/playwright-go"
)

type LoginPage struct {
	BasePage
	signInLnk playwright.Locator
	username  playwright.Locator
	password  playwright.Locator
	signInBtn playwright.Locator
}

func NewLoginPage(p playwright.Page, t *testing.T) *LoginPage {
	l := &LoginPage{BasePage: *NewBasePage(p, nil, t)}
	l.signInLnk = l.page.GetByRole("link", playwright.PageGetByRoleOptions{Name: *playwright.String("Sign in ")})
	l.username = l.page.GetByLabel("Username")
	l.password = l.page.GetByLabel("Password")
	l.signInBtn = l.page.GetByRole("button", playwright.PageGetByRoleOptions{Name: *playwright.String("Sign in")})
	return l
}

func (l *LoginPage) GoTo() *LoginPage {
	l.Navigate("http://localhost:8080/")
	return l
}

func (l *LoginPage) ClickSignInLink() *LoginPage {
	l.Click(l.signInLnk)
	return l
}

func (l *LoginPage) FillUsername(username string) *LoginPage {
	l.Fill(l.username, username)
	return l
}

func (l *LoginPage) ClickPassword() *LoginPage {
	l.Click(l.password)
	return l
}

func (l *LoginPage) FillPassword(password string) *LoginPage {
	l.Fill(l.password, password)
	return l
}

func (l *LoginPage) ClickSignInButton() *WelcomePage {
	l.Click(l.signInBtn)
	return NewWelcomePage(l.page, &l.t)
}
