# cpu-pushbullet

![img](https://www.pushbullet.com/img/header/logo.png)

This little CLI app written in Go will allow you to monitor your average computer's CPU usage and enables you to get notifications onto all your devices with the help of Pushbullet.
[Pushbullet](https://www.pushbullet.com/) is available for [Android](https://play.google.com/store/apps/details?id=com.pushbullet.android&pcampaignid=web_share), [Windows](https://update.pushbullet.com/pushbullet_installer.exe) aswell as [Chromium](https://addons.mozilla.org/en-US/firefox/addon/pushbullet/versions/) and [Firefox](https://addons.mozilla.org/en-US/firefox/addon/pushbullet/versions/) based Browsers.

All you have to do is visit https://www.pushbullet.com/#settings/account and [Create Access Token].

Paste that into the correspondig slot in your .env file and everything is set up!

Change the filename .env.changeme to .env

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
# Settings
## Needs to be changed
PUSHBULLET_API_KEY = "put your key here and rename this file from .env.changeme to .env "

## Can be configured to your liking or stay default

`CPU_AVERAGE_MAX_THRESHOLD = "80.0"`

`CHECK_INTERVAL_SECONDS = "1"`

`TIMESPAN_AVERAGE_MINUTES = "1"`

`THRESHOLD_DURATION_ALARM_MINUTES = "5"`

`ENABLE_CONSOLE_OUTPUT = "true"`

`SEND_TEST_NOTIFICATION_ON_LAUNCH = "true"`

## Keep default values
`PUSHBULLET_ENDPOINT_URL = "https://api.pushbullet.com/v2/pushes"`

