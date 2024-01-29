package po

import (
	"testing"

	"github.com/playwright-community/playwright-go"
)

type MainMenuPage struct {
	BasePage
	workPackagesItm playwright.Locator
}

func NewMainMenuPage(p playwright.Page, t *testing.T) *MainMenuPage {
	m := &MainMenuPage{BasePage: *NewBasePage(p, nil, t)}
	m.workPackagesItm = m.page.Locator("#main-menu-work-packages")
	return m
}

func (m *MainMenuPage) ClickWorkPackagesItm() *AllOpenWpPage {
	m.Click(m.workPackagesItm)
	return NewAllOpenWpPage(m.page, &m.t)
}
