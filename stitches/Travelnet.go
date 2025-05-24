package stitches

import (
	"image"
	"image/color"

	"codeberg.org/ostech/craft_to_clonia_textures/data"
	imaging "github.com/disintegration/imaging"
)

func RWTravelnet(input_pack_path, output_pack_path string) error {
	stitch := "Travelnet"
	craft_path := "block"
	clonia_path := "travelnet"

	in_path := input_pack_path + data.CraftPaths["block"]
	out_path := output_pack_path + "/" + data.CloniaPaths["travelnet"]
	var block string

	block = "lodestone_side.png"
	lodestone_side, err := imaging.Open(in_path + block)
	if err != nil {
		return openErrMsg(stitch, craft_path, block)
	}

	block = "lodestone_top.png"
	lodestone_top, err := imaging.Open(in_path + block)
	if err != nil {
		return openErrMsg(stitch, craft_path, block)
	}

	block = "glass.png"
	glass, err := imaging.Open(in_path + block)
	if err != nil {
		return openErrMsg(stitch, craft_path, block)
	}
	glass_scale := glass.Bounds().Dx() / 16
	edgeless_glass_crop := imaging.CropCenter(glass, glass_scale*12, glass_scale*12)
	edgeless_glass := imaging.New(16*glass_scale, 16*glass_scale, color.NRGBA{0, 0, 0, 0})
	edgeless_glass = imaging.Paste(edgeless_glass, edgeless_glass_crop, image.Pt(4*glass_scale, 4*glass_scale))
	edgeless_glass_crop = imaging.CropCenter(edgeless_glass, glass_scale*12, glass_scale*12)
	edgeless_glass = imaging.New(16*glass_scale, 16*glass_scale, color.NRGBA{0, 0, 0, 0})
	edgeless_glass = imaging.PasteCenter(edgeless_glass, edgeless_glass_crop)

	block = "white_wool.png"
	carpet, err := imaging.Open(in_path + block)
	if err != nil {
		return openErrMsg(stitch, craft_path, block)
	}

	block = "orange_wool.png"
	elevator_carpet, err := imaging.Open(in_path + block)
	if err != nil {
		return openErrMsg(stitch, craft_path, block)
	}

	//

	block = "travelnet_travelnet_front_color.png"
	if err := imaging.Save(TravelnetFrontColor(lodestone_side, lodestone_top), out_path+block); err != nil {
		return saveErrMsg(stitch, clonia_path, block)
	}

	block = "travelnet_travelnet_front.png"
	if err := imaging.Save(TravelnetFront(), out_path+block); err != nil {
		return saveErrMsg(stitch, clonia_path, block)
	}

	block = "travelnet_travelnet_side.png"
	if err := imaging.Save(TravelnetSide(edgeless_glass, edgeless_glass), out_path+block); err != nil {
		return saveErrMsg(stitch, clonia_path, block)
	}

	block = "travelnet_travelnet_side_color.png"
	if err := imaging.Save(TravelnetSideColor(lodestone_side, lodestone_top), out_path+block); err != nil {
		return saveErrMsg(stitch, clonia_path, block)
	}

	block = "travelnet_bottom.png"
	if err := imaging.Save(TravelnetBottom(carpet), out_path+block); err != nil {
		return saveErrMsg(stitch, clonia_path, block)
	}

	block = "travelnet_top.png"
	if err := imaging.Save(TravelnetTop(lodestone_top), out_path+block); err != nil {
		return saveErrMsg(stitch, clonia_path, block)
	}

	block = "travelnet_travelnet_back.png"
	if err := imaging.Save(TravelnetBack(), out_path+block); err != nil {
		return saveErrMsg(stitch, clonia_path, block)
	}

	block = "travelnet_travelnet_back_color.png"
	if err := imaging.Save(TravelnetBackColor(lodestone_side, lodestone_top), out_path+block); err != nil {
		return saveErrMsg(stitch, clonia_path, block)
	}

	block = "travelnet_elevator_door_glass.png"
	if err := imaging.Save(TravelnetGlassDoor(glass), out_path+block); err != nil {
		return saveErrMsg(stitch, clonia_path, block)
	}

	// Elevator Time
	block = "travelnet_elevator_front.png"
	if err := imaging.Save(ElevatorFront(lodestone_top, lodestone_top), out_path+block); err != nil {
		return saveErrMsg(stitch, clonia_path, block)
	}

	block = "travelnet_elevator_inside_floor.png"
	if err := imaging.Save(TravelnetBottom(elevator_carpet), out_path+block); err != nil {
		return saveErrMsg(stitch, clonia_path, block)
	}

	block = "travelnet_elevator_inside_controls.png"
	if err := imaging.Save(ElevatorFront(lodestone_side, lodestone_top), out_path+block); err != nil {
		return saveErrMsg(stitch, clonia_path, block)
	}

	block = "travelnet_elevator_sides_outside.png"
	if err := imaging.Save(ElevatorFront(lodestone_top, lodestone_top), out_path+block); err != nil {
		return saveErrMsg(stitch, clonia_path, block)
	}

	block = "travelnet_elevator_inside_ceiling.png"
	if err := imaging.Save(TravelnetBottom(lodestone_top), out_path+block); err != nil {
		return saveErrMsg(stitch, clonia_path, block)
	}

	return nil
}

func TravelnetFrontColor(top, bot image.Image) image.Image {
	scale := top.Bounds().Dx() / 16
	dst := imaging.New(16*scale, 32*scale, color.NRGBA{0, 0, 0, 0})

	top = imaging.AdjustBrightness(top, 22)
	bot = imaging.AdjustBrightness(bot, 22)
	dst = imaging.Paste(dst, top, image.Pt(0, 0))
	dst = imaging.Paste(dst, bot, image.Pt(0, 16*scale))

	return dst
}

func TravelnetFront() image.Image { // it's blank, lol
	dst := imaging.New(16, 32, color.NRGBA{0, 0, 0, 0})
	return dst
}

func TravelnetSide(top, bot image.Image) image.Image {
	scale := top.Bounds().Dx() / 16
	dst := imaging.New(16*scale, 32*scale, color.NRGBA{0, 0, 0, 0})

	dst = imaging.Paste(dst, top, image.Pt(0, 0))
	dst = imaging.Paste(dst, bot, image.Pt(0, 16*scale))

	return dst
}

func TravelnetSideColor(top, bot image.Image) image.Image {
	scale := top.Bounds().Dx() / 16
	dst := imaging.New(16*scale, 32*scale, color.NRGBA{0, 0, 0, 0})

	dst = imaging.Paste(dst, top, image.Pt(0, 0))
	dst = imaging.Paste(dst, bot, image.Pt(0, 16*scale))
	dst = imaging.AdjustBrightness(dst, 22)
	dst = imaging.PasteCenter(dst, imaging.New(12*scale, 28*scale, color.Transparent))

	return dst
}

func TravelnetBottom(bot image.Image) image.Image {
	scale := bot.Bounds().Dx() / 16
	dst := imaging.New(16*scale, 16*scale, color.Transparent)

	dst = imaging.Paste(dst, bot, image.Pt(0, 0))
	return dst
}

func TravelnetTop(top image.Image) image.Image {
	scale := top.Bounds().Dx() / 16
	dst := imaging.New(16*scale, 16*scale, color.Transparent)

	top = imaging.AdjustBrightness(top, 22)

	dst = imaging.Paste(dst, top, image.Pt(0, 0))
	return dst
}

func TravelnetBack() image.Image { // it's blank, lol
	dst := imaging.New(16, 32, color.NRGBA{0, 0, 0, 0})
	return dst
}

func TravelnetBackColor(top, bot image.Image) image.Image {
	scale := top.Bounds().Dx() / 16
	dst := imaging.New(16*scale, 32*scale, color.NRGBA{0, 0, 0, 0})

	top = imaging.AdjustBrightness(top, 22)
	bot = imaging.AdjustBrightness(bot, 22)
	dst = imaging.Paste(dst, top, image.Pt(0, 0))
	dst = imaging.Paste(dst, bot, image.Pt(0, 16*scale))

	return dst
}

func TravelnetGlassDoor(glass image.Image) image.Image {
	scale := glass.Bounds().Dx() / 16
	dst := imaging.New(16*scale, 16*scale, color.Transparent)

	glass = imaging.AdjustBrightness(glass, 22)
	blank_strip := imaging.New(14*scale, scale, color.Transparent)
	glass = imaging.Paste(glass, blank_strip, image.Pt(scale, 0))
	glass = imaging.Paste(glass, blank_strip, image.Pt(scale, 15*scale))

	dst = imaging.Paste(dst, glass, image.Pt(0, 0))
	return dst
}

func ElevatorFront(top, bot image.Image) image.Image {
	scale := top.Bounds().Dx() / 16
	dst := imaging.New(16*scale, 32*scale, color.NRGBA{0, 0, 0, 0})

	//top = imaging.AdjustBrightness(top, 22)
	//bot = imaging.AdjustBrightness(bot, 22)
	dst = imaging.Paste(dst, top, image.Pt(0, 0))
	dst = imaging.Paste(dst, bot, image.Pt(0, 16*scale))

	return dst
}
