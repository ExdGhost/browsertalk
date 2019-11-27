# browsertalk
Simple browser interaction via rest api's using golang
Start: http://<server>/start?browser=chrome&url=http://example.com should start Google Chrome and open http://example.com in the same.
Stop: http://<server>/stop?browser=<browser> should stop the given browser if it is running.
Cleanup: http://<server>/cleanup?browser=<browser> should clean up the browsing session for the given browser if it has been stopped.
Get Active Tab:  http://<server>/geturl?browser=<browser> should get the current active tab URL for the given browser.