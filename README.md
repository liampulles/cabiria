<div align="center"><img src="cabiria-1914-poster.jpg" alt="Film poster for Cabiria (1914)" width="650px"></div>
<div align="center"><small><sup>Film poster for <i>Cabiria (1914)</i>. Original artwork by Leopoldo Metlicovitz.</sup></small></div>
<h1 align="center">
  Cabiria
</h1>

<h4 align="center"> An ASS intertitle generator, for silent films.</a></h4>

<p align="center">
  <a href="#-status">Status</a> •
  <a href="#-key-objectives">Key Objectives</a> •
  <a href="#-install">Install</a> •
  <a href="#-basic-usage">Basic Usage</a> •
  <a href="#-planned-usage">Planned Usage</a> •
  <a href="#-contributing">Contributing</a> •
  <a href="#-license">License</a>
</p>

<p align="center">
  <a href="https://travis-ci.com/liampulles/cabiria">
    <img src="https://travis-ci.com/liampulles/cabiria.svg?branch=master" alt="[Build Status]">
  </a>
    <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/liampulles/cabiria">
  <a href="https://goreportcard.com/report/github.com/liampulles/cabiria">
    <img src="https://goreportcard.com/badge/github.com/liampulles/cabiria" alt="[Go Report Card]">
  </a>
  <a href="https://codecov.io/gh/liampulles/cabiria">
    <img src="https://codecov.io/gh/liampulles/cabiria/branch/master/graph/badge.svg" />
  </a>
</p>

## ⚔️ Status

Cabiria is currently in pre-alpha. Stay tuned for an upcoming release!

## 🛡️ Key Objectives

* Generate pretty ASS intertitles, in a style that is not jarring.

## 🗡️ Install

As Cabiria is currently in heavy development, no installation candidate is available at this time. If you're *really* eager, you will need to set up a development environment as per the Wiki to use the application.

## 🤺 Basic Usage


To generate appropriate styled intertitles for existing (`LesVampires1915.srt`) subtitles:

```bash
    cabiria generate LesVampires1915.mkv LesVampires1915.srt LesVampires1915.ass
```

## 🎭 Planned Usage

* `cabiria resync`: Sync external subtitles to intertitles

## 🐉 Contributing

If you wish to make a change, then I suggest making an issue for your proposal.
If you're interested in helping out more generally, then <a href="mailto:me@liampulles.com">drop me a mail</a>.

## 🦄 License

See [LICENSE](LICENSE)