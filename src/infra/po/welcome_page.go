package po

import (
	"testing"

	"github.com/playwright-community/playwright-go"
)

type WelcomePage struct {
	BasePage
	selectProjectLnk playwright.Locator
}

func NewWelcomePage(p playwright.Page, t *testing.T) *WelcomePage {
	w := &WelcomePage{BasePage: *NewBasePage(p, nil, t)}
	w.selectProjectLnk = w.page.GetByRole("link", playwright.PageGetByRoleOptions{Name: *playwright.String("Select a project ")})
	return w
}

func (w *WelcomePage) GoTo() *WelcomePage {
	w.Navigate("http://localhost:8080/")
	return w
}

func (w *WelcomePage) SelectProject(project string) *OverviewPage {
	w.Click(w.selectProjectLnk)
	w.Click(w.page.Locator("li.ui-menu-item:has-text('" + project + "')"))
	return NewOverviewPage(w.page, &w.t)
}
