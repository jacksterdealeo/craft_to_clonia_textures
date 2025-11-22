# Config

## DefinedInput
> Decides whether to use InputDir config option.
- Default: false
- Accepts: true/false

## DefinedOutput
> Decides whether to use OutputDir config option.
- Default: false
- Accepts: true/false

## ExportMinetest_Game
> Decides whether textures are converted to Minetest Game's format.
- Default: false
- Accepts: true/false

## ExportMineclonia
> Decides whether textures are converted to Mineclonia's format.
- Default: true
- Accepts: true/false

## InputDir
> Changes where Minecraft texture packs are read.
- Default: guessed vanilla Minecraft's texture pack path
- Accepts: path

## OutputDir
> Changes where Luanti texture packs are placed.
- Default: guessed Luanti's texture pack path
- Accepts: path

## HUDOnFireAnimationFrames
> Changes how many frames the fire overlay has when you are on fire. Does not affect fire blocks.
- Default: 8
- Accepts: positive integers
- Minecraft uses 32, but you need to make a change to your server's configuration for it can look right.

## SpearVersion
> Changes which texture to use for spear items.
- Default: "short"
- Accepts: "long", "old", "short"
- Minecraft uses "short" textures for spears as an item and "long" for spears thrown or in hand. "old" will switch to textures I made from before Minecraft got spears.
