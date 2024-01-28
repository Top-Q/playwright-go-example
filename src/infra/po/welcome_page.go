package po

import (
	"testing"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/require"
)

type WelcomePage struct {
	BasePage
	selectProjectLnk playwright.Locator
}

func NewWelcomePage(p playwright.Page, t *testing.T) *WelcomePage {
	w := &WelcomePage{BasePage: BasePage{page: p, t: *t, assertion: *require.New(t)}}
	w.selectProjectLnk = w.page.GetByRole("link", playwright.PageGetByRoleOptions{Name: *playwright.String("Select a project ")})
	return w
}

func (w *WelcomePage) GoTo() *WelcomePage {
	_, err := w.page.Goto("http://localhost:8080/")
	w.assertion.Nil(err)
	return w
}

func (w *WelcomePage) SelectProject(project string) *WelcomePage {
	err := w.selectProjectLnk.Click()
	w.assertion.Nil(err)
	err = w.page.Locator("li.ui-menu-item:has-text('" + project + "')").Click()
	w.assertion.Nil(err)
	return w
}
