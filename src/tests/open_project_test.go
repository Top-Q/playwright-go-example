package tests

import (
	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/require"
)

func (s *Suite) TestCreateNewIssue() {
	require := require.New(s.T())
	p := s.page
	taskName := "Test task"
	_, err := p.Goto("http://localhost:8080/")
	require.Nil(err)
	err = p.GetByRole("link", playwright.PageGetByRoleOptions{Name: *playwright.String("Sign in ")}).Click()
	require.Nil(err)
	err = p.GetByLabel("Username").Fill("admin")
	require.Nil(err)
	err = p.GetByLabel("Password").Click()
	require.Nil(err)
	err = p.GetByLabel("Password").Fill("adminadmin")
	require.Nil(err)
	err = p.GetByRole("button", playwright.PageGetByRoleOptions{Name: *playwright.String("Sign in")}).Click()
	require.Nil(err)
	err = p.GetByRole("link", playwright.PageGetByRoleOptions{Name: *playwright.String("Select a project ")}).Click()
	require.Nil(err)
	err = p.Locator("li.ui-menu-item:has-text('Demo project')").Click()
	require.Nil(err)
	err = p.Locator("#main-menu-work-packages").Click()
	require.Nil(err)
	err = p.Locator("wp-create-button").GetByLabel("Create new work package").Click()
	require.Nil(err)
	err = s.page.GetByLabel("Task", playwright.PageGetByLabelOptions{Exact: playwright.Bool(true)}).Click()
	require.Nil(err)
	err = p.GetByLabel("Subject", playwright.PageGetByLabelOptions{Exact: playwright.Bool(true)}).Fill(taskName)
	require.Nil(err)
	err = p.Locator("#wp-new-inline-edit--field-assignee").GetByRole("combobox").Click()
	require.Nil(err)
	err = p.GetByLabel("Estimated time", playwright.PageGetByLabelOptions{Exact: playwright.Bool(true)}).Click()
	require.Nil(err)
	err = p.GetByLabel("Estimated time", playwright.PageGetByLabelOptions{Exact: playwright.Bool(true)}).Fill("10")
	require.Nil(err)
	err = p.GetByRole("button", playwright.PageGetByRoleOptions{Name: *playwright.String("Save")}).Click()
	require.Nil(err)
	err = p.GetByRole("link", playwright.PageGetByRoleOptions{Name: *playwright.String("Up")}).Click()
	require.Nil(err)
	err = p.GetByRole("link", playwright.PageGetByRoleOptions{Name: *playwright.String("Work packages")}).Click()
	require.Nil(err)
	// await page.getByRole('combobox').click();
	// await page.getByRole('combobox').fill(randomTaskName);
	// await page.getByRole('button', { name: 'Search' }).click();
	// await expect(page.getByRole('table')).toContainText(randomTaskName);
	// await page.getByRole('link', { name: 'OA', exact: true }).click();
	// await page.getByRole('link', { name: 'Sign out' }).click();

}
