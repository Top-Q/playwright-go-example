package po

import (
	"testing"

	"github.com/playwright-community/playwright-go"
)

type BasePage struct {
	ActionBot
	// In case the page is inside iFrame, all elements will be searched inside the frame

	frameLocator playwright.FrameLocator
}

func NewBasePage(page playwright.Page, frameLocator playwright.FrameLocator, t *testing.T) *BasePage {
	b := &BasePage{ActionBot: *NewActionBot(page, t)}
	b.frameLocator = frameLocator
	b.TakeScreenshot("")
	return b
}
