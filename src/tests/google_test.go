package tests

import (
	"fmt"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/require"
)

func (s *Suite) TestFillCheeseInGoogle() {
	require := require.New(s.T())
	fmt.Println("In the test")
	_, err := s.page.Goto("https://google.com")
	require.Nil(err)
	locator := s.page.GetByLabel("חיפוש", playwright.PageGetByLabelOptions{Exact: playwright.Bool(true)})
	locator.Click()
	err = s.page.GetByLabel("חיפוש", playwright.PageGetByLabelOptions{Exact: playwright.Bool(true)}).Click()
	require.Nil(err)
	err = s.page.GetByLabel("חיפוש", playwright.PageGetByLabelOptions{Exact: playwright.Bool(true)}).Fill("Cheese")
	require.Nil(err)
	err = s.page.GetByLabel("חיפוש ב-Google").First().Click()
	require.Nil(err)
}
