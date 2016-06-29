[![Travis build status](http://img.shields.io/travis/mdreizin/chrome-bookmarks-alfred-workflow/master.svg?style=flat-square)](https://travis-ci.org/mdreizin/chrome-bookmarks-alfred-workflow)
[![Coverage Status](https://img.shields.io/coveralls/mdreizin/chrome-bookmarks-alfred-workflow/master.svg?style=flat-square)](https://coveralls.io/r/mdreizin/chrome-bookmarks-alfred-workflow?branch=master)

<h1 id="chrome-bookmarks-alfred-workflow">chrome-bookmarks-alfred-workflow</h1>

> Chrome/Canary/Chromium bookmarks search workflow for Alfred

- [x] Google Chrome (`chrome`, `chrome-profiles`)
- [x] Google Chrome Canary (`canary`, `canary-profiles`)
- [x] Chromium (`chromium`, `chromium-profiles`)

![](https://raw.github.com/mdreizin/chrome-bookmarks-alfred-workflow/master/.gitdown/bookmarks.gif)
![](https://raw.github.com/mdreizin/chrome-bookmarks-alfred-workflow/master/.gitdown/profiles.gif)

<h2 id="chrome-bookmarks-alfred-workflow-setup">Setup</h2>

* Run `brew install go`
* Run `go get github.com/tools/godep`
* Run `make restore`

<h2 id="chrome-bookmarks-alfred-workflow-develop">Develop</h2>

* Run `make workflow`
* Open `./dist` folder

<h2 id="chrome-bookmarks-alfred-workflow-test">Test</h2>

* Run `make test`

<h2 id="chrome-bookmarks-alfred-workflow-cover">Cover</h2>

* Run `make cover` or `make cover-html`
