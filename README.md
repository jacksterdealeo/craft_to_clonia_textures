<h1 align="center">Craft to Clonia Textures</h1>

## Build
- Build for yourself by using 'go build .'.
- Build for every supported platform by using 'go run build.go'.

## Disclaimers
- The "Minetest Game" is supported, but not the default target. You just need to enable it in the program's config file.
- VoxeLibre is supported. It's textures are exported with Mineclonia exports.
- Some Minecraft texture packs have a license that prohibits using their packs outside of Minecraft. This software does not care what packs you convert with it, so use responsibly.
- All mobs, some UI, and other things will not transfer over.
Some things are not reasonably possible to transfer (like the brewing stand),
and many things are just not added in yet.
- If a pack has custom models for blocks, that can cause problems with the transfer.
- The compatiblity rating is loose. Minecraft itself wouldn't report perfect compatiblity.

## Usage
- Run the project once, and it'll create an input folder, an output folder, and a config file. The config file has options you may want to check before continuing.
- Place texture packs from Minecraft into the input folder.
- Run the project again, and it'll place the converted packs into the output folder.
- The config file is explained here: [CONFIG.md](https://codeberg.org/ostech/craft_to_clonia_textures/src/branch/main/CONFIG.md)

## License
This project is distributed under the terms of [MIT License](LICENSE).

## Thank You
This project is over a year old, and it has been fun working on it. Thank you all for playing!
