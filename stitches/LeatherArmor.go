package stitches

import (
	"image"
	"image/color"

	"codeberg.org/ostech/craft_to_clonia_textures/data"
	imaging "github.com/disintegration/imaging"
)

func RWLeatherArmor(input_pack_path, output_pack_path string) error {
	var (
		stitch           = "Leather Armor"
		armor_path       = input_pack_path + data.CraftPaths["humanoid"]
		leggings_path    = input_pack_path + data.CraftPaths["humanoid_leggings"]
		item_path        = input_pack_path + data.CraftPaths["item"]
		out_path         = output_pack_path + data.CloniaPaths["armor"]
		filter_of_choice = imaging.Lanczos
	)

	in_texture_name := "leather.png"
	in_overlay_name := "leather_overlay.png"

	out_helmet_name := "mcl_armor_helmet_leather.png"
	out_chestplate_name := "mcl_armor_chestplate_leather.png"
	out_leggings_name := "mcl_armor_leggings_leather.png"
	out_boots_name := "mcl_armor_boots_leather.png"

	out_helmet_desat_name := "mcl_armor_helmet_leather_desat.png"
	out_chestplate_desat_name := "mcl_armor_chestplate_leather_desat.png"
	out_leggings_desat_name := "mcl_armor_leggings_leather_desat.png"
	out_boots_desat_name := "mcl_armor_boots_leather_desat.png"

	src, err := imaging.Open(armor_path + in_texture_name)
	if err != nil {
		return openErrMsg(stitch, "humanoid", in_texture_name)
	}
	leather_texture := ImageToNRGBA(src)

	src, err = imaging.Open(armor_path + in_overlay_name)
	if err != nil {
		return openErrMsg(stitch, "humanoid", in_overlay_name)
	}
	leather_overlay := ImageToNRGBA(src)

	src, err = imaging.Open(leggings_path + in_texture_name)
	if err != nil {
		return openErrMsg(stitch, "humanoid_leggings", in_texture_name)
	}
	leather_pants_texture := ImageToNRGBA(src)

	src, err = imaging.Open(leggings_path + in_overlay_name)
	if err != nil {
		return openErrMsg(stitch, "humanoid_leggings", in_overlay_name)
	}
	leather_pants_overlay := ImageToNRGBA(src)

	in_textures := data.MCArmorSet{
		MostParts: &leather_texture,
		Leggings:  &leather_pants_texture,
	}
	in_overlay := data.MCArmorSet{
		MostParts: &leather_overlay,
		Leggings:  &leather_pants_overlay,
	}

	normal, desat := LeatherArmorWorn(in_textures, in_overlay, filter_of_choice)

	if err := imaging.Save(normal.Helmet, out_path+out_helmet_name); err != nil {
		return saveErrMsg(stitch, "armor", out_helmet_name)
	}

	if err := imaging.Save(desat.Helmet, out_path+out_helmet_desat_name); err != nil {
		return saveErrMsg(stitch, "armor", out_helmet_desat_name)
	}

	if err := imaging.Save(normal.Chestplate, out_path+out_chestplate_name); err != nil {
		return saveErrMsg(stitch, "armor", out_chestplate_name)
	}

	if err := imaging.Save(desat.Chestplate, out_path+out_chestplate_desat_name); err != nil {
		return saveErrMsg(stitch, "armor", out_chestplate_desat_name)
	}

	if err := imaging.Save(normal.Leggings, out_path+out_leggings_name); err != nil {
		return saveErrMsg(stitch, "armor", out_chestplate_desat_name)
	}

	if err := imaging.Save(desat.Leggings, out_path+out_leggings_desat_name); err != nil {
		return saveErrMsg(stitch, "armor", out_chestplate_desat_name)
	}

	if err := imaging.Save(normal.Boots, out_path+out_boots_name); err != nil {
		return saveErrMsg(stitch, "armor", out_chestplate_desat_name)
	}

	if err := imaging.Save(desat.Boots, out_path+out_boots_desat_name); err != nil {
		return saveErrMsg(stitch, "armor", out_chestplate_desat_name)
	}

	// items

	i_helmet := "leather_helmet.png"
	i_helmet_overlay := "leather_helmet_overlay.png"
	i_leggings := "leather_leggings.png"
	i_leggings_overlay := "leather_leggings_overlay.png"
	i_boots := "leather_boots.png"
	i_boots_overlay := "leather_boots_overlay.png"
	i_chestplate := "leather_chestplate.png"

	o_helmet := "mcl_armor_inv_helmet_leather.png"
	o_helmet_desat := "mcl_armor_inv_helmet_leather_desat.png"
	o_leggings := "mcl_armor_inv_leggings_leather.png"
	o_leggings_desat := "mcl_armor_inv_leggings_leather_desat.png"
	o_boots := "mcl_armor_inv_boots_leather.png"
	o_boots_desat := "mcl_armor_inv_boots_leather_desat.png"
	o_chestplate := "mcl_armor_inv_chestplate_leather.png"
	o_chestplate_desat := "mcl_armor_inv_chestplate_leather_desat.png"

	i_helmet_img, err := imaging.Open(item_path + i_helmet)
	if err != nil {
		return openErrMsg(stitch, "item", i_helmet)
	}
	i_helmet_overlay_img, err := imaging.Open(item_path + i_helmet_overlay)
	if err != nil {
		return openErrMsg(stitch, "item", i_helmet_overlay)
	}

	out_helmet, out_helmet_desat := LeatherArmorItem(i_helmet_img, i_helmet_overlay_img)
	if err := imaging.Save(out_helmet, out_path+o_helmet); err != nil {
		return saveErrMsg(stitch, "armor", o_helmet)
	}
	if err := imaging.Save(out_helmet_desat, out_path+o_helmet_desat); err != nil {
		return saveErrMsg(stitch, "armor", o_helmet_desat)
	}

	i_leggings_img, err := imaging.Open(item_path + i_leggings)
	if err != nil {
		return openErrMsg(stitch, "item", i_leggings)
	}
	i_leggings_overlay_img, err := imaging.Open(item_path + i_leggings_overlay)
	if err != nil {
		return openErrMsg(stitch, "item", i_leggings_overlay)
	}

	out_leggings, out_leggings_desat := LeatherArmorItem(i_leggings_img, i_leggings_overlay_img)
	if err := imaging.Save(out_leggings, out_path+o_leggings); err != nil {
		return saveErrMsg(stitch, "armor", o_leggings)
	}
	if err := imaging.Save(out_leggings_desat, out_path+o_leggings_desat); err != nil {
		return saveErrMsg(stitch, "armor", o_leggings_desat)
	}

	i_boots_img, err := imaging.Open(item_path + i_boots)
	if err != nil {
		return openErrMsg(stitch, "item", i_boots)
	}
	i_boots_overlay_img, err := imaging.Open(item_path + i_boots_overlay)
	if err != nil {
		return openErrMsg(stitch, "item", i_boots_overlay)
	}

	out_boots, out_boots_desat := LeatherArmorItem(i_boots_img, i_boots_overlay_img)
	if err := imaging.Save(out_boots, out_path+o_boots); err != nil {
		return saveErrMsg(stitch, "armor", o_boots)
	}
	if err := imaging.Save(out_boots_desat, out_path+o_boots_desat); err != nil {
		return saveErrMsg(stitch, "armor", o_boots_desat)
	}

	// The chestplate item doesn't have an overlay.
	i_chestplate_img, err := imaging.Open(item_path + i_chestplate)
	if err != nil {
		return openErrMsg(stitch, "item", i_chestplate)
	}
	i_chestplate_overlay_img := imaging.New(i_chestplate_img.Bounds().Dx(), i_chestplate_img.Bounds().Dy(), color.Transparent)

	out_chestplate, out_chestplate_desat := LeatherArmorItem(i_chestplate_img, i_chestplate_overlay_img)
	if err := imaging.Save(out_chestplate, out_path+o_chestplate); err != nil {
		return saveErrMsg(stitch, "armor", o_chestplate)
	}
	if err := imaging.Save(out_chestplate_desat, out_path+o_chestplate_desat); err != nil {
		return saveErrMsg(stitch, "armor", o_chestplate_desat)
	}

	return nil
}

func LeatherArmorItem(mc_texture, mc_overlay image.Image) (normal, desat image.Image) {
	dst := imaging.New(mc_texture.Bounds().Dx(), mc_texture.Bounds().Dy(), color.Transparent)
	dst = imaging.Paste(dst, mc_texture, image.Pt(0, 0))
	dst = imaging.Overlay(dst, mc_overlay, image.Pt(0, 0), 1.0)

	desat = dst
	dst = MakeBrown(dst)
	normal = dst

	return
}

func LeatherArmorWorn(mc_set, mc_overlay data.MCArmorSet, filter imaging.ResampleFilter) (normal, desat data.ArmorSetTextures) {
	scale := mc_set.MostParts.Bounds().Dx() / 64

	makeScaled := func(texture *image.NRGBA) *image.NRGBA {
		if scale != 1 {
			texture = imaging.Fit(texture, texture.Bounds().Dx()/scale, texture.Bounds().Dy()/scale, filter)
		}
		return texture
	}

	helmet := CropToScale(mc_set.MostParts, 0, 0, 32, 16, scale)
	helmet = makeScaled(helmet)
	helmet_out := imaging.New(64, 32, color.NRGBA{0, 0, 0, 0})
	helmet_out = imaging.Paste(helmet_out, helmet, image.Pt(32, 0))

	helmet_overlay := CropToScale(mc_overlay.MostParts, 0, 0, 32, 16, scale)
	helmet_overlay = makeScaled(helmet_overlay)
	helmet_overlay = imaging.Grayscale(helmet_overlay)
	helmet_overlay_out := imaging.New(64, 32, color.NRGBA{0, 0, 0, 0})
	helmet_overlay_out = imaging.Paste(helmet_overlay_out, helmet_overlay, image.Pt(32, 0))

	helmet_out = imaging.Overlay(helmet_out, helmet_overlay_out, image.Pt(0, 0), 1.0)

	chestplate := CropToScale(mc_set.MostParts, 16, 16, 56, 32, scale)
	chestplate = makeScaled(chestplate)
	chestplate_out := imaging.New(64, 32, color.NRGBA{0, 0, 0, 0})
	chestplate_out = imaging.Paste(chestplate_out, chestplate, image.Pt(16, 16))

	chestplate_overlay := CropToScale(mc_overlay.MostParts, 16, 16, 56, 32, scale)
	chestplate_overlay = makeScaled(chestplate_overlay)
	chestplate_overlay = imaging.Grayscale(chestplate_overlay)
	chestplate_overlay_out := imaging.New(64, 32, color.NRGBA{0, 0, 0, 0})
	chestplate_overlay_out = imaging.Paste(chestplate_overlay_out, chestplate_overlay, image.Pt(16, 16))

	chestplate_out = imaging.Overlay(chestplate_out, chestplate_overlay, image.Pt(0, 0), 1.0)

	boots := CropToScale(mc_set.MostParts, 0, 16, 16, 32, scale)
	boots = makeScaled(boots)
	boots_out := imaging.New(64, 32, color.NRGBA{0, 0, 0, 0})
	boots_out = imaging.Paste(boots_out, boots, image.Pt(0, 16))

	boots_overlay := CropToScale(mc_overlay.MostParts, 0, 16, 16, 32, scale)
	boots_overlay = makeScaled(boots_overlay)
	boots_overlay = imaging.Grayscale(boots_overlay)
	boots_overlay_out := imaging.New(64, 32, color.NRGBA{0, 0, 0, 0})
	boots_overlay_out = imaging.Paste(boots_overlay_out, boots_overlay, image.Pt(0, 16))

	boots_out = imaging.Overlay(boots_out, boots_overlay_out, image.Pt(0, 0), 1.0)

	//

	leggings := mc_set.Leggings
	leggings = makeScaled(leggings)
	leggings_out := imaging.New(64, 32, color.NRGBA{0, 0, 0, 0})
	leggings_out = imaging.Paste(leggings_out, leggings, image.Pt(0, 0))

	leggings_overlay := mc_overlay.Leggings
	leggings_overlay = makeScaled(leggings_overlay)
	leggings_overlay = imaging.Grayscale(leggings_overlay)
	leggings_overlay_out := imaging.New(64, 32, color.NRGBA{0, 0, 0, 0})
	leggings_overlay_out = imaging.Paste(leggings_overlay_out, leggings_overlay, image.Pt(0, 0))

	leggings_out = imaging.Overlay(leggings_out, leggings_overlay_out, image.Pt(0, 0), 1.0)

	//

	desat.Helmet = helmet_out
	desat.Chestplate = chestplate_out
	desat.Boots = boots_out
	desat.Leggings = leggings_out

	helmet_out = MakeBrown(helmet_out)
	chestplate_out = MakeBrown(chestplate_out)
	boots_out = MakeBrown(boots_out)
	leggings_out = MakeBrown(leggings_out)

	normal.Helmet = helmet_out
	normal.Chestplate = chestplate_out
	normal.Boots = boots_out
	normal.Leggings = leggings_out

	return
}
