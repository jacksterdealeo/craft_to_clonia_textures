package stitches

import (
	"image"
	"image/color"

	"codeberg.org/ostech/craft_to_clonia_textures/data"
	imaging "github.com/disintegration/imaging"
)

type Travelnet_t struct {
	Front_Color *image.NRGBA
	Front       *image.NRGBA
	Side_Color  *image.NRGBA
	Side        *image.NRGBA
	Bottom      *image.NRGBA
	Top         *image.NRGBA
	Back_Color  *image.NRGBA
	Back        *image.NRGBA

	// 64x64 are the OG item textures
	Inv_Base  *image.NRGBA
	Inv_Color *image.NRGBA
}

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

	Travelnet := Travelnet_t{}

	Travelnet.Front_Color = TravelnetFrontColor(lodestone_side, lodestone_top)
	Travelnet.Front = TravelnetFront()
	Travelnet.Side_Color = TravelnetSideColor(lodestone_side, lodestone_top)
	Travelnet.Side = TravelnetSide(edgeless_glass, edgeless_glass)

	Travelnet.Bottom = TravelnetBottom(carpet)
	Travelnet.Top = TravelnetTop(lodestone_top)
	Travelnet.Back_Color = TravelnetBackColor(lodestone_side, lodestone_top)
	Travelnet.Back = TravelnetBack()

	block = "travelnet_travelnet_front_color.png"
	err = imaging.Save(Travelnet.Front_Color, out_path+block)
	if err != nil {
		return saveErrMsg(stitch, clonia_path, block)
	}

	block = "travelnet_travelnet_front.png"
	err = imaging.Save(Travelnet.Front, out_path+block)
	if err != nil {
		return saveErrMsg(stitch, clonia_path, block)
	}

	block = "travelnet_travelnet_side_color.png"
	err = imaging.Save(Travelnet.Side_Color, out_path+block)
	if err != nil {
		return saveErrMsg(stitch, clonia_path, block)
	}

	block = "travelnet_travelnet_side.png"
	err = imaging.Save(Travelnet.Side, out_path+block)
	if err != nil {
		return saveErrMsg(stitch, clonia_path, block)
	}

	block = "travelnet_bottom.png"
	err = imaging.Save(Travelnet.Bottom, out_path+block)
	if err != nil {
		return saveErrMsg(stitch, clonia_path, block)
	}

	block = "travelnet_top.png"
	err = imaging.Save(Travelnet.Top, out_path+block)
	if err != nil {
		return saveErrMsg(stitch, clonia_path, block)
	}

	block = "travelnet_travelnet_back_color.png"
	err = imaging.Save(Travelnet.Back_Color, out_path+block)
	if err != nil {
		return saveErrMsg(stitch, clonia_path, block)
	}

	block = "travelnet_travelnet_back.png"
	err = imaging.Save(Travelnet.Back, out_path+block)
	if err != nil {
		return saveErrMsg(stitch, clonia_path, block)
	}

	// Door
	block = "travelnet_elevator_door_glass.png"
	err = imaging.Save(TravelnetGlassDoor(glass), out_path+block)
	if err != nil {
		return saveErrMsg(stitch, clonia_path, block)
	}

	elevator := Travelnet_t{}

	elevator.Front = ElevatorFront(lodestone_top, lodestone_top)
	elevator.Bottom = TravelnetBottom(elevator_carpet)
	elevator.Back = ElevatorFront(lodestone_side, lodestone_top) // Controls
	elevator.Side = ElevatorFront(lodestone_top, lodestone_top)  // Blank
	elevator.Top = TravelnetBottom(lodestone_top)

	Travelnet.Inv_Base, Travelnet.Inv_Color, elevator.Inv_Base = TravelnetInv(&Travelnet, &elevator)

	// Elevator
	block = "travelnet_elevator_front.png"
	err = imaging.Save(elevator.Front, out_path+block)
	if err != nil {
		return saveErrMsg(stitch, clonia_path, block)
	}

	block = "travelnet_elevator_inside_floor.png"
	err = imaging.Save(elevator.Bottom, out_path+block)
	if err != nil {
		return saveErrMsg(stitch, clonia_path, block)
	}

	block = "travelnet_elevator_inside_controls.png"
	err = imaging.Save(elevator.Back, out_path+block)
	if err != nil {
		return saveErrMsg(stitch, clonia_path, block)
	}

	block = "travelnet_elevator_sides_outside.png"
	err = imaging.Save(elevator.Side, out_path+block)
	if err != nil {
		return saveErrMsg(stitch, clonia_path, block)
	}

	block = "travelnet_elevator_inside_ceiling.png"
	err = imaging.Save(elevator.Top, out_path+block)
	if err != nil {
		return saveErrMsg(stitch, clonia_path, block)
	}

	// Item icons

	block = "travelnet_inv_base.png"
	err = imaging.Save(Travelnet.Inv_Base, out_path+block)
	if err != nil {
		return saveErrMsg(stitch, clonia_path, block)
	}

	block = "travelnet_inv_colorable.png"
	err = imaging.Save(Travelnet.Inv_Color, out_path+block)
	if err != nil {
		return saveErrMsg(stitch, clonia_path, block)
	}

	block = "travelnet_elevator_inv.png"
	err = imaging.Save(elevator.Inv_Base, out_path+block)
	if err != nil {
		return saveErrMsg(stitch, clonia_path, block)
	}

	return nil
}

func TravelnetFrontColor(top, bot image.Image) *image.NRGBA {
	scale := top.Bounds().Dx() / 16
	dst := imaging.New(16*scale, 32*scale, color.NRGBA{0, 0, 0, 0})

	top = imaging.AdjustBrightness(top, 22)
	bot = imaging.AdjustBrightness(bot, 22)
	dst = imaging.Paste(dst, top, image.Pt(0, 0))
	dst = imaging.Paste(dst, bot, image.Pt(0, 16*scale))

	return dst
}

func TravelnetFront() *image.NRGBA { // it's blank, lol
	dst := imaging.New(16, 32, color.NRGBA{0, 0, 0, 0})
	return dst
}

func TravelnetSide(top, bot image.Image) *image.NRGBA {
	scale := top.Bounds().Dx() / 16
	dst := imaging.New(16*scale, 32*scale, color.NRGBA{0, 0, 0, 0})

	dst = imaging.Paste(dst, top, image.Pt(0, 0))
	dst = imaging.Paste(dst, bot, image.Pt(0, 16*scale))

	return dst
}

func TravelnetSideColor(top, bot image.Image) *image.NRGBA {
	scale := top.Bounds().Dx() / 16
	dst := imaging.New(16*scale, 32*scale, color.NRGBA{0, 0, 0, 0})

	dst = imaging.Paste(dst, top, image.Pt(0, 0))
	dst = imaging.Paste(dst, bot, image.Pt(0, 16*scale))
	dst = imaging.AdjustBrightness(dst, 22)
	dst = imaging.PasteCenter(dst, imaging.New(12*scale, 28*scale, color.Transparent))

	return dst
}

func TravelnetBottom(bot image.Image) *image.NRGBA {
	scale := bot.Bounds().Dx() / 16
	dst := imaging.New(16*scale, 16*scale, color.Transparent)

	dst = imaging.Paste(dst, bot, image.Pt(0, 0))
	return dst
}

func TravelnetTop(top image.Image) *image.NRGBA {
	scale := top.Bounds().Dx() / 16
	dst := imaging.New(16*scale, 16*scale, color.Transparent)

	top = imaging.AdjustBrightness(top, 22)

	dst = imaging.Paste(dst, top, image.Pt(0, 0))
	return dst
}

func TravelnetBack() *image.NRGBA { // it's blank, lol
	dst := imaging.New(16, 32, color.NRGBA{0, 0, 0, 0})
	return dst
}

func TravelnetBackColor(top, bot image.Image) *image.NRGBA {
	scale := top.Bounds().Dx() / 16
	dst := imaging.New(16*scale, 32*scale, color.NRGBA{0, 0, 0, 0})

	top = imaging.AdjustBrightness(top, 22)
	bot = imaging.AdjustBrightness(bot, 22)
	dst = imaging.Paste(dst, top, image.Pt(0, 0))
	dst = imaging.Paste(dst, bot, image.Pt(0, 16*scale))

	return dst
}

func TravelnetGlassDoor(glass image.Image) *image.NRGBA {
	scale := glass.Bounds().Dx() / 16
	dst := imaging.New(16*scale, 16*scale, color.Transparent)

	glass = imaging.AdjustBrightness(glass, 22)
	blank_strip := imaging.New(14*scale, scale, color.Transparent)
	glass = imaging.Paste(glass, blank_strip, image.Pt(scale, 0))
	glass = imaging.Paste(glass, blank_strip, image.Pt(scale, 15*scale))

	dst = imaging.Paste(dst, glass, image.Pt(0, 0))
	return dst
}

func ElevatorFront(top, bot image.Image) *image.NRGBA {
	scale := top.Bounds().Dx() / 16
	dst := imaging.New(16*scale, 32*scale, color.NRGBA{0, 0, 0, 0})

	//top = imaging.AdjustBrightness(top, 22)
	//bot = imaging.AdjustBrightness(bot, 22)
	dst = imaging.Paste(dst, top, image.Pt(0, 0))
	dst = imaging.Paste(dst, bot, image.Pt(0, 16*scale))

	return dst
}

func TravelnetInv(travelnet *Travelnet_t, elevator *Travelnet_t) (*image.NRGBA, *image.NRGBA, *image.NRGBA) {
	dstSize := float64(64)
	dst := imaging.New(int(dstSize), int(dstSize), color.NRGBA{0, 0, 0, 0})

	// bot
	dst = PerspectiveOverlay(
		*dst,
		*travelnet.Bottom,
		dstSize*0.50, dstSize*0.70, dstSize*0.70, dstSize*0.80,
		dstSize*0.30, dstSize*0.80, dstSize*0.50, dstSize*0.90,
	)

	// behind side
	dst = PerspectiveOverlay(
		*dst,
		*travelnet.Side_Color,
		dstSize*0.20, dstSize*0.15, dstSize*0.50, dstSize*0.10,
		dstSize*0.30, dstSize*0.80, dstSize*0.50, dstSize*0.70,
	)

	// back
	dst = PerspectiveOverlay(
		*dst,
		*travelnet.Back_Color,
		dstSize*0.50, dstSize*0.10, dstSize*0.75, dstSize*0.15,
		dstSize*0.50, dstSize*0.70, dstSize*0.70, dstSize*0.80,
	)

	// front side
	dst = PerspectiveOverlay(
		*dst,
		*travelnet.Side_Color,
		dstSize*0.50, dstSize*0.25, dstSize*0.75, dstSize*0.15,
		dstSize*0.50, dstSize*0.90, dstSize*0.70, dstSize*0.80,
	)

	// top
	dst = PerspectiveOverlay(
		*dst,
		*travelnet.Top,
		dstSize*0.50, dstSize*0.10, dstSize*0.75, dstSize*0.15,
		dstSize*0.20, dstSize*0.15, dstSize*0.50, dstSize*0.25,
	)

	travelnet.Inv_Base = imaging.Clone(dst)
	base_net := dst

	// Now, we make the color version. //
	dst = imaging.New(int(dstSize), int(dstSize), color.NRGBA{0, 0, 0, 0})

	// behind side
	dst = PerspectiveOverlay(
		*dst,
		*travelnet.Back_Color,
		dstSize*0.50, dstSize*0.10, dstSize*0.75, dstSize*0.15,
		dstSize*0.50, dstSize*0.70, dstSize*0.70, dstSize*0.80,
	)

	// back
	dst = PerspectiveOverlay(
		*dst,
		*travelnet.Side_Color,
		dstSize*0.20, dstSize*0.15, dstSize*0.50, dstSize*0.10,
		dstSize*0.30, dstSize*0.80, dstSize*0.50, dstSize*0.70,
	)

	// front side
	dst = PerspectiveOverlay(
		*dst,
		*travelnet.Side_Color,
		dstSize*0.50, dstSize*0.25, dstSize*0.75, dstSize*0.15,
		dstSize*0.50, dstSize*0.90, dstSize*0.70, dstSize*0.80,
	)

	/*
		// top
		// this should remove color from the top
		// note: i did not. :(
		dst = imaging.OverlayCenter(dst, imaging.AdjustBrightness(PerspectiveOverlay(
			*imaging.New(dst.Rect.Dx(), dst.Rect.Dy(), color.Transparent),
			*travelnet.Top,
			dstSize*0.22, dstSize*0.10,   dstSize*0.46, dstSize*0.07,
			dstSize*0.58, dstSize*0.13,   dstSize*0.77, dstSize*0.09,
		), -100), 1.0)
	*/

	// top
	dst = PerspectiveOverlay(
		*dst,
		*travelnet.Top,
		dstSize*0.50, dstSize*0.10, dstSize*0.75, dstSize*0.15,
		dstSize*0.20, dstSize*0.15, dstSize*0.50, dstSize*0.25,
	)

	color_net := imaging.Clone(dst)

	// Now, we make the elevator.

	dst = imaging.New(int(dstSize), int(dstSize), color.NRGBA{0, 0, 0, 0})

	// bot
	dst = PerspectiveOverlay(
		*dst,
		*elevator.Bottom,
		dstSize*0.50, dstSize*0.70, dstSize*0.70, dstSize*0.80,
		dstSize*0.30, dstSize*0.80, dstSize*0.50, dstSize*0.90,
	)

	// behind side
	dst = PerspectiveOverlay(
		*dst,
		*elevator.Side,
		dstSize*0.50, dstSize*0.10, dstSize*0.75, dstSize*0.15,
		dstSize*0.50, dstSize*0.70, dstSize*0.70, dstSize*0.80,
	)

	// control side
	dst = PerspectiveOverlay(
		*dst,
		*elevator.Back,
		dstSize*0.20, dstSize*0.15, dstSize*0.50, dstSize*0.10,
		dstSize*0.30, dstSize*0.80, dstSize*0.50, dstSize*0.70,
	)

	// front side
	dst = PerspectiveOverlay(
		*dst,
		*elevator.Side,
		dstSize*0.50, dstSize*0.25, dstSize*0.75, dstSize*0.15,
		dstSize*0.50, dstSize*0.90, dstSize*0.70, dstSize*0.80,
	)

	// top
	dst = PerspectiveOverlay(
		*dst,
		*elevator.Top,
		dstSize*0.50, dstSize*0.10, dstSize*0.75, dstSize*0.15,
		dstSize*0.20, dstSize*0.15, dstSize*0.50, dstSize*0.25,
	)

	elevator.Inv_Base = imaging.Clone(dst)
	elevator_icon := elevator.Inv_Base

	return base_net, color_net, elevator_icon
}
