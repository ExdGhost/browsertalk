package browser

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
)

var instance context.CancelFunc

// Start ...
func (browser *Chrome) Start() (err error) {
	opts := []chromedp.ExecAllocatorOption{chromedp.Flag("no-first-run", true)}

	allocContext, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancel := chromedp.NewContext(allocContext)
	// defer cancel()

	instance = cancel

	err = chromedp.Run(ctx, chromedp.Navigate("https://www.facebook.com/"))
	if err != nil {
		log.Println("ERROR@Chrome Run = ", err.Error())
		return err
	}
	// time.Sleep(15 * time.Second)

	// Returns
	return nil
}
