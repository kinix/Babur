# Babür

Babür is a discord bot created for personal purposes.

The bot's capabilities are rolling dice, converting some imperial units to metric units and, finding an URL to learn it.

In addition, Babür can answer some questions with random responses. (WIP - Sample json is for Turkish)

## Installing

Create an app in Discord Developer Portal. Set the token of that app as `BABUR_TOKEN` environment variable.

Invite the bot into your channel and TADA! He's ready for your messages.

## Configiration

You can change maximum dice count and side number in `config/dice.json`

You can change or add units in `config/units.json`. The units with `_` prefix are being used for only convert between metric units (e.g. cm to m)

### Chat Config

You can change or add regex for categories in `chat_regex.json`

You can change or add responses in `chat.json`. You have to have same groups with `chat_regex.json` in this file.

In addition you have to the special groups groups which are `_?` and `_` to respond unmatched messages. `_?` is for the messages with question mark. `_` is for the others.

### Image Search

If you want to use Google Image search in your chat functions, you have to define `GOOGLE_TOKEN` and `GOOGLE_CX` environment variables.

For more information: https://developers.google.com/custom-search/v1/introduction

## Usage

### Roll Dice

If your messages started with one of below format, Babür will roll dice for you.

Example formats:
- 1d20
- d6
- 2d10
- 2d20 +5

### Converting Units

If your messages contains any imperial units in `config/units.json`, Babür will convert the measurements to metric units.

### DnD Search

If your messages starts with !dnd, Babür will search the message and find a link contain the info you searched.
