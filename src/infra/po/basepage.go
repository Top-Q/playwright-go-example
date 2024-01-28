package po

import (
	"testing"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/require"
)

type BasePage struct {
	page      playwright.Page
	t         testing.T
	assertion require.Assertions
}
