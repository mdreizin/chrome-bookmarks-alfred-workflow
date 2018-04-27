[![Travis build status](https://img.shields.io/travis/mdreizin/chrome-bookmarks-alfred-workflow/master.svg)](https://travis-ci.org/mdreizin/chrome-bookmarks-alfred-workflow)
[![Code Climate GPA](https://img.shields.io/codeclimate/maintainability/mdreizin/chrome-bookmarks-alfred-workflow.svg)](https://codeclimate.com/github/mdreizin/chrome-bookmarks-alfred-workflow)
[![Code Climate Coverage](https://img.shields.io/codeclimate/coverage/mdreizin/chrome-bookmarks-alfred-workflow.svg)](https://codeclimate.com/github/mdreizin/chrome-bookmarks-alfred-workflow)

# chrome-bookmarks-alfred-workflow

> Chrome/Canary/Chromium bookmarks search workflow for Alfred

- [x] Google Chrome (`chrome`, `chrome-profiles`)
- [x] Google Chrome Canary (`canary`, `canary-profiles`)
- [x] Chromium (`chromium`, `chromium-profiles`)

![Search Bookmarks](screenshot-bookmarks.gif)
![Select Profiles](screenshot-profiles.gif)

## Download

All available releases you can find [here](https://github.com/mdreizin/chrome-bookmarks-alfred-workflow/releases).

## Options

`*-profiles`

`Auto` checks active browser profile and uses it to search bookmarks.

## Setup

- Run `brew install go glide`
- Run `make deps`

## Develop

- Run `make workflow`
- Open `./build` folder

## Test

- Run `make test`

## Cover

- Run `make cover` or `make cover-html`
