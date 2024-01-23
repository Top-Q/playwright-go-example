package tests

import (
	"fmt"
	"testing"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	playwright *playwright.Playwright
	browser    playwright.Browser
	context    playwright.BrowserContext
	page       playwright.Page
}

func (s *Suite) SetupTest() {
	fmt.Println("In the setup test")
	require := require.New(s.T())

	// Run playwright
	pw, err := playwright.Run()
	require.Nil(err)
	s.playwright = pw

	// Launch browser
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false), // Set headless to false
	})
	require.Nil(err)
	s.browser = browser // Assign browser to the Browser field

	// Create context
	context, err := browser.NewContext(playwright.BrowserNewContextOptions{
		RecordVideo: &playwright.RecordVideo{
			Dir: "videos/",
		},
		// StorageStatePath: playwright.String("state/state.json"),
	})
	require.Nil(err)
	context.Tracing().Start(playwright.TracingStartOptions{
		Snapshots:   playwright.Bool(true),
		Sources:     playwright.Bool(true),
		Screenshots: playwright.Bool(true),
	})
	s.context = context // Assign context to the Context field
	// defer context.Tracing().Stop("test-results/trace.zip")
	page, err := context.NewPage()
	require.Nil(err)
	s.page = page // Assign page to the Page field

}

// cleanup will run after each test in the suite
func (s *Suite) TearDownTest() {
	fmt.Println("Cleaning up after test")
	s.context.Tracing().Stop("./test-results/trace.zip")
	s.page.Close()
	s.context.Close()
	s.browser.Close()
	s.playwright.Stop()

}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
