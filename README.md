# cpu-pushbullet

![img](https://www.pushbullet.com/img/header/logo.png)

This little CLI app written in Go will allow you to monitor your average computer's CPU usage and enables you to get notifications onto all your devices with the help of Pushbullet.
[Pushbullet](https://www.pushbullet.com/) is available for [Android](https://play.google.com/store/apps/details?id=com.pushbullet.android&pcampaignid=web_share), [Windows](https://update.pushbullet.com/pushbullet_installer.exe) aswell as [Chromium](https://addons.mozilla.org/en-US/firefox/addon/pushbullet/versions/) and [Firefox](https://addons.mozilla.org/en-US/firefox/addon/pushbullet/versions/) based Browsers.

All you have to do is visit https://www.pushbullet.com/#settings/account and [Create Access Token].

Paste that into the correspondig slot in your .env file and everything is set up!




# Settings
## Needs to be changed
- Change the filename `.env.changeme` to `.env`

- `PUSHBULLET_API_KEY = "put your key here and rename this file from .env.changeme to .env "`

## Can be configured to your liking or stay default

- `CPU_AVERAGE_MAX_THRESHOLD = "80.0"` - Averaged CPU load in percent that will trigger a push notification and a log entry

- `CHECK_INTERVAL_SECONDS = "1"` - Check current CPU load every XX seconds

- `TIMESPAN_AVERAGE_MINUTES = "1"` - Length of time window in minutes that will be used for calculating average CPU load

- `THRESHOLD_DURATION_ALARM_MINUTES = "5"` - Currently not used

- `ENABLE_CONSOLE_OUTPUT = "true"` - Can be disabled if app is run in background

- `SEND_TEST_NOTIFICATION_ON_LAUNCH = "true"` - To test if your API Key is set up correctly this will send you a notification via Pushbullet once on startup

## Keep default values
- `PUSHBULLET_ENDPOINT_URL = "https://api.pushbullet.com/v2/pushes"`



Example Console output:
```bash
WARNING average CPU load: 45.36% - momentary 42.69%
WARNING average CPU load: 45.67% - momentary 71.09%
WARNING average CPU load: 46.04% - momentary 55.00%
WARNING average CPU load: 46.41% - momentary 58.98%
WARNING average CPU load: 46.68% - momentary 51.54%
WARNING average CPU load: 47.09% - momentary 56.92%
WARNING average CPU load: 47.36% - momentary 56.54%
```


Example `application.log` content:
```
CPU: 2023/10/27 15:22:59 CPU watcher launched.
CPU: 2023/10/27 15:23:00 Environment variables loaded.
CPU: 2023/10/27 15:23:22 Closing Application
CPU: 2023/10/27 15:23:25 CPU watcher launched.
CPU: 2023/10/27 15:23:26 Environment variables loaded.
```

# Docker
The simplest way to use the Docker image is by using 
```
docker run -e PUSHBULLET_API_KEY="YourAPIKeyHere" ghcr.io/stefandanzl/cpu-pushbullet:latest
```
