# GawkBox Mobile

> A React-Native app written with GO to search and watch your favorite Twitch Streams!

## Table of Contents

1.  [User-Story](#user-story)
1.  [Requirements](#requirements)
1.  [Installing Dependencies](#installing-dependencies)
1.  [Usage](#Usage)
    1.  [Running-with-Docker](#running-with-docker)
    1.  [Running-without-Docker](#running-without-docker)

## User Story

Here's a walkthrough of implemented user stories:

<img src='https://i.imgur.com/8VIoc4l.gif' title='Video Walkthrough' width='350' alt='Video Walkthrough' />

[walkthrough](https://i.imgur.com/8VIoc4l.gif)

## Requirements

* [GO](https://golang.org/dl/) (if not running with Docker)
* [Docker](https://docs.docker.com/install/)
* [NPM](https://www.npmjs.com/get-npm)
* [XCode for iOS Simulator](https://developer.apple.com/download/)

## Installing Dependencies

From within the root directory:

```sh
npm install
```

## Usage

### Get repo on local machine

If you have GO configured correctly you can run and the project will be located on your computer ${GOPATH}/github.com/chasekaylee/gawkbox-mobile

```
go get -u github.com/chasekaylee/gawkbox-mobile
```

OR

You can get this project by using git clone. Navigate to the where you want your project to be located ${HOME}/Desktop

```
git clone https://github.com/chasekaylee/gawkbox-mobile.git

OR

download zip file from top of repo 'CLONE OR DOWNLOAD'
```

## Running with Docker

I recommend this so you do not have to wrestle with setting up GO (if not installed)

Run these commands from the root directory of the project folder
${HOME}/${Wherever you cloned file}/gawkbox-mobile

If you haven't installed Docker yet see [REQUIREMENTS](#requirements)

WITH DOCKER CLOUD

```sh
docker pull chasekaylee/gawkbox-mobile
docker run -p 8080:8080 chasekaylee/gawkbox-mobile
npm run ios
```

WITHOUT DOCKER CLOUD

```sh
docker build -t gawkbox-mobile .
docker run -p 8080:8080 gawkbox-mobile
npm run ios
```

## Running without Docker

If you haven't installed GO yet see [REQUIREMENTS](#requirements)

```
go build
go run main.go OR ./gawkbox-mobile
npm run ios
```
