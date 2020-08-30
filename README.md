# nabcli
CLI tool to manage my Nabaztag

## Overview

This tool is used to control a [Nabaztag](https://en.wikipedia.org/wiki/Nabaztag) with [new board](https://www.tagtagtag.fr/).

[This project](https://github.com/nabaztag2018/pynab) use [a daemon](https://github.com/nabaztag2018/pynab/tree/master/nabd) that communicate through a [TCP json protocol](https://github.com/nabaztag2018/pynab/blob/master/PROTOCOL.md) that this CLI use.

## Usage

### Wakeup

```sh
nabcli wakeup
```

### Sleep

```sh
nabcli sleep
```

### [With _systemd_](contrib/systemd)

systemd can be used to program sleep and wakeup with `timer` units.
