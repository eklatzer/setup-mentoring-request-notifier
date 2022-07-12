# setup-mentoring-request-notifier

![lint](https://github.com/eklatzer/setup-mentoring-request-notifier/actions/workflows/lint.yml/badge.svg)
![build](https://github.com/eklatzer/setup-mentoring-request-notifier/actions/workflows/build.yml/badge.svg)

Used for setting up the threads for
the [exercism-mentoring-request-notifier](https://github.com/eklatzer/exercism-mentoring-request-notifier)

# Usage

1. Get the default config:

```
./setup-mentoring-request-notifier generate-config
```

2. Get the available tracks

```
./setup-mentoring-request-notifier get-tracks 
```

3. Setup Slack:

```
./setup-mentoring-request-notifier setup-slack --slack-token="<Slack token>" --slack-delay=<Slack message delay> --default-channel-id="<Channel ID>"
```

# Commands

Flags for all commands:

* `config`: Path and file containing the config (default: `config.json`)

## generate-config

`generate-config` is used to generate the default config.

## get-tracks

`get-tracks` is used to collect all currently available tracks on Exercism.

Possible flags:

* `url`: HTTP url to request the track info at (default: `https://exercism.org/api/v2/tracks`)

## setup-slack

`setup-slack` is used to start the thread on Slack for each track with the message "Thread for `<Trackname>`" and to
print the needed track configuration for the mentoring request notifier.

Possible flags:

* `slack-token`: OAuth-Token for Slack
* `slack-delay`: Delay after sending a message to Slack in ms (rate limiting of Slack, default: `500`)
* `default-channel-id`: Default channel id that is set for each track that has no channel id.

