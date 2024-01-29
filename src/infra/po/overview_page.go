package po

import (
	"testing"

	"github.com/playwright-community/playwright-go"
)

type OverviewPage struct {
	BasePage
	MainMenuPage
	selectProjectLnk playwright.Locator
}

func NewOverviewPage(p playwright.Page, t *testing.T) *OverviewPage {
	w := &OverviewPage{BasePage: *NewBasePage(p, nil, t), MainMenuPage: *NewMainMenuPage(p, t)}
	w.selectProjectLnk = w.page.GetByRole("link", playwright.PageGetByRoleOptions{Name: *playwright.String("Select a project ")})
	return w
}
