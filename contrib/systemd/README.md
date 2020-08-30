# Systemd timers

## Installation

- Disable `nabclockd`

```sh
sudo systemctl stop nabclockd.service
sudo systemctl disable nabclockd.service
```

- Copy files in `/etc/systemd/system`

- Activate timers

```sh
sudo systemctl enable nabsleep.timer nabwakeup.timer
```

## Configuration

In `*.timer` files, change `OnCalendar=` directive.
