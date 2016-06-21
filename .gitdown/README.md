{"gitdown": "badge", "name": "travis"}
{"gitdown": "badge", "name": "coveralls"}

# {"gitdown": "gitinfo", "name": "name"}

> Chrome/Canary/Chromium bookmarks search workflow for Alfred

Google Chrome (`chrome`, `chrome-profiles`)

![](https://raw.github.com/{"gitdown": "gitinfo", "name": "username"}/{"gitdown": "gitinfo", "name": "name"}/{"gitdown": "gitinfo", "name": "branch"}/.gitdown/chrome.gif)
![](https://raw.github.com/{"gitdown": "gitinfo", "name": "username"}/{"gitdown": "gitinfo", "name": "name"}/{"gitdown": "gitinfo", "name": "branch"}/.gitdown/chrome-profiles.gif)

Google Chrome Canary (`canary`, `canary-profiles`)

![](https://raw.github.com/{"gitdown": "gitinfo", "name": "username"}/{"gitdown": "gitinfo", "name": "name"}/{"gitdown": "gitinfo", "name": "branch"}/.gitdown/canary.gif)
![](https://raw.github.com/{"gitdown": "gitinfo", "name": "username"}/{"gitdown": "gitinfo", "name": "name"}/{"gitdown": "gitinfo", "name": "branch"}/.gitdown/canary-profiles.gif)

Chromium (`chromium`, `chromium-profiles`)

![](https://raw.github.com/{"gitdown": "gitinfo", "name": "username"}/{"gitdown": "gitinfo", "name": "name"}/{"gitdown": "gitinfo", "name": "branch"}/.gitdown/chromium.gif)
![](https://raw.github.com/{"gitdown": "gitinfo", "name": "username"}/{"gitdown": "gitinfo", "name": "name"}/{"gitdown": "gitinfo", "name": "branch"}/.gitdown/chromium-profiles.gif)

## Setup

* Run `brew install go`
* Run `go get github.com/tools/godep`
* Run `godep restore`

## Develop

* Run `make workflow`
* Open `./dist` folder

## Test

* Run `make test`

## Cover

* Run `make cover` or `make cover-html`
