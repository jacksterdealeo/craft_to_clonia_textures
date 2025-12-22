<h1 align="center">Craft to Clonia Textures</h1>

## Build

- Use the Go compiler at [go.dev](https://go.dev/). This project requires at least version 1.24.2.

- Build for yourself by using `go build .`.

- Build for every supported platform by using `go run build.go`. This is not recommended for personal use.

## Disclaimers

- Custom models are not supported by this project. Luanti's texture packs only allow for custom textures. If you convert a pack that makes use of custom models, there will be garbled results. I have not found a way around this limitation.

- “Minetest Game” is supported, but not the default target. You just need to enable it in the program's config file.

- VoxeLibre is supported. Its textures are exported with Mineclonia exports. This may change in the future.

- Some Minecraft texture packs have a license that prohibits using their packs outside of Minecraft. This software does not care what packs you convert with it, so use responsibly.

- All mobs, some UI, and other things will not transfer over. Some things are not reasonably possible to transfer (like the brewing stand), and many things are just not added in yet.

- The compatibility rating is loose. Minecraft itself wouldn't report perfect compatibility.

## Usage

- Run the project once, and it'll create an input folder, an output folder, and a config file. The config file has options you may want to check before continuing.

- Place texture packs from Minecraft into the input folder.

- Run the project again, and it'll place the converted packs into the output folder.

- The config file is explained at [`CONFIG.md`](https://codeberg.org/ostech/craft_to_clonia_textures/src/branch/main/CONFIG.md).

- You can find information on how to install texture packs on the Luanti website, at https://docs.luanti.org/for-players/texture-packs/

## License

This project is distributed under the terms of [the MIT License](LICENSE).

## Thank You

This project is mainly for my own use, but I am happy other people have found it helpful.

Additional help was given by:

- [lemonzest](https://codeberg.org/lemonzest)

- [RokeJulianLockhart](https://codeberg.org/RokeJulianLockhart)

- [arcensyl](https://github.com/arcensyl)
