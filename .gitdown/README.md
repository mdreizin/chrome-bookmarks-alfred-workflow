{"gitdown": "badge", "name": "travis"}
{"gitdown": "badge", "name": "coveralls"}

# {"gitdown": "gitinfo", "name": "name"}

> Chrome/Canary/Chromium bookmarks search workflow for Alfred

- [x] Google Chrome (`chrome`, `chrome-profiles`)
- [x] Google Chrome Canary (`canary`, `canary-profiles`)
- [x] Chromium (`chromium`, `chromium-profiles`)

![](https://raw.github.com/{"gitdown": "gitinfo", "name": "username"}/{"gitdown": "gitinfo", "name": "name"}/{"gitdown": "gitinfo", "name": "branch"}/.gitdown/bookmarks.gif)
![](https://raw.github.com/{"gitdown": "gitinfo", "name": "username"}/{"gitdown": "gitinfo", "name": "name"}/{"gitdown": "gitinfo", "name": "branch"}/.gitdown/profiles.gif)

## Download

All available releases you can find [here]({"gitdown": "gitinfo", "name": "url"}/releases).

## Options

`*-profiles`

`Auto` checks active browser profile and uses it to search bookmarks.

## Setup

* Run `brew install go glide`
* Run `make restore`

## Develop

* Run `make workflow`
* Open `./dist` folder

## Test

* Run `make test`

## Cover

* Run `make cover` or `make cover-html`
