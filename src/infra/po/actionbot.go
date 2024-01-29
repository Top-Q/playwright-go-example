package po

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"testing"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/require"
)

type ActionBot struct {
	page      playwright.Page
	t         testing.T
	assertion require.Assertions
}

func NewActionBot(p playwright.Page, t *testing.T) *ActionBot {
	a := &ActionBot{page: p, t: *t, assertion: *require.New(t)}
	return a
}

func (a *ActionBot) Click(locator playwright.Locator) *ActionBot {
	err := locator.Click()
	a.assertion.NoError(err, "Failed to click on element: %v", locator)
	return a
}

func (a *ActionBot) Fill(locator playwright.Locator, value string) *ActionBot {
	err := locator.Fill(value)
	a.assertion.NoError(err, "failed to fill value in element: %v", locator)
	return a
}

func (a *ActionBot) Navigate(url string) *ActionBot {
	_, err := a.page.Goto(url)
	a.assertion.NoError(err, "failed to go to url: %v", url)
	return a
}

func (a *ActionBot) Check(t *testing.T, locator playwright.Locator) *ActionBot {
	err := locator.Check()
	a.assertion.NoError(err, "Failed to check element: %v", locator)
	return a
}

func (a *ActionBot) IsEnabled(t *testing.T, locator playwright.Locator) bool {
	isEnabled, err := locator.IsEnabled()
	a.assertion.NoError(err, "Element isn't enabled: %v", locator)
	return isEnabled
}

func (a *ActionBot) IsVisible(t *testing.T, locator playwright.Locator) bool {
	isVisible, err := locator.IsVisible()
	a.assertion.NoError(err, "Element is invisible: %v", locator)
	return isVisible
}

func (a *ActionBot) TakeScreenshot(description string) {
	if description == "" {
		description = "screenshot"
	}
	testName := strings.Split(a.t.Name(), "/")[1]
	_, err := os.Stat("screenshots/")
	if os.IsNotExist(err) {
		// Directory does not exist, create it
		err := os.Mkdir("screenshots/", 0777)
		a.assertion.NoError(err, "Failed to create screenshots directory at screenshots")
	}
	testScreenshotDir := filepath.Join("screenshots", testName)
	_, err = os.Stat(testScreenshotDir)
	if os.IsNotExist(err) {
		// Directory does not exist, create it
		err := os.Mkdir(testScreenshotDir, 0777)
		a.assertion.NoError(err, "Failed to create screenshots directory at "+testScreenshotDir)
	}
	screenshotName := fmt.Sprintf("%s_%s_%d.png", filepath.Base(a.t.Name()), description, time.Now().Unix())
	screenshotPath := filepath.Join(testScreenshotDir, screenshotName)

	//screenshot, err := page.Screenshot()
	_, err = a.page.Screenshot(playwright.PageScreenshotOptions{
		Path: playwright.String(screenshotPath),
	})
	a.assertion.NoError(err, "Failed to take screenshot at "+screenshotPath)

}
