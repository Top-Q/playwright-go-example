package po

import (
	"testing"

	"github.com/playwright-community/playwright-go"
)

type AllOpenWpPage struct {
	BasePage
	createBtn  playwright.Locator
	taskOption playwright.Locator
}

func NewAllOpenWpPage(p playwright.Page, t *testing.T) *AllOpenWpPage {
	l := &AllOpenWpPage{BasePage: *NewBasePage(p, nil, t)}
	l.createBtn = p.Locator("wp-create-button").GetByLabel("Create new work package")
	l.taskOption = p.GetByLabel("Task", playwright.PageGetByLabelOptions{Exact: playwright.Bool(true)})
	return l
}

func (l *AllOpenWpPage) ClickOnCreateTask() *AllOpenWpPage {
	l.Click(l.createBtn).Click(l.taskOption)
	return l
}
