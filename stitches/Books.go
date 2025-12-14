package stitches

import (
	"fmt"
	"image"
	"path/filepath"
	"strconv"

	"codeberg.org/ostech/craft_to_clonia_textures/configure"
	"codeberg.org/ostech/craft_to_clonia_textures/data"
	imaging "github.com/disintegration/imaging"
)

func RWChiseledBooks(input_pack_path, output_pack_path string, _ *configure.Config) error {
	stitch := "Books"

	in_path := filepath.Join(input_pack_path, data.CraftPaths["block"])
	out_path := filepath.Join(output_pack_path, data.CloniaPaths["books"])
	var err error

	file_name := "chiseled_bookshelf_occupied.png"
	block, err := imaging.Open(filepath.Join(in_path, file_name))
	if err != nil {
		return openErrMsg(stitch, in_path, file_name)
	}
	if block.Bounds().Dx() < 16 {
		return fmt.Errorf(stitch, "Texture is too small. (less than 16px)")
	}
	scale := block.Bounds().Dx() / 16

	var chiseled_books = make([]*image.NRGBA, 6)
	chiseled_books[0] = CropToScale(block, 1, 1, 5, 7, scale)
	chiseled_books[1] = CropToScale(block, 6, 1, 10, 7, scale)
	chiseled_books[2] = CropToScale(block, 11, 1, 15, 7, scale)

	chiseled_books[3] = CropToScale(block, 1, 9, 5, 15, scale)
	chiseled_books[4] = CropToScale(block, 6, 9, 10, 15, scale)
	chiseled_books[5] = CropToScale(block, 11, 9, 15, 15, scale)

	for i, e := range chiseled_books {
		var file_name = "mcl_books_book_" + strconv.Itoa(i) + ".png"
		if err = imaging.Save(e, filepath.Join(out_path, file_name)); err != nil {
			return saveErrMsg(stitch, out_path, file_name)
		}
	}
	return nil
}
