# Babür

Babür is a discord bot created for personal purposes.

The bot's capabilities are rolling dice and converting some imperial units to metric units.

## Installing

Create an app in Discord Developer Portal. Set the token of that app as `TOKEN` environment variable.

Invite the bot into your channel and TADA! He's ready for your messages.

## Configiration

You can change maximum dice count and side number from `config/dice.json`

You can change or add units from `config/units.json`

The units with `_` prefix are being used for only convert between metric units (e.g. cm to m)

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