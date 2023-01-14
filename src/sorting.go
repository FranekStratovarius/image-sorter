package main

import (
	"image"
	"image/color"
	"sort"
)

type sortColor []color.NRGBA

func get_column(img *image.NRGBA, column_index int, height int) sortColor {
	column := make(sortColor, height)
	for i := 0; i < height; i++ {
		column[i] = img.NRGBAAt(column_index, i)
	}
	return column
}

func (column sortColor) Len() int {
	return len(column)
}

func (column sortColor) Swap(i, j int) {
	column[i], column[j] = column[j], column[i]
}

func get_luminance(color color.NRGBA) float64 {
	return 0.2126*float64(color.R) + 0.7152*float64(color.G) + 0.0722*float64(color.B)
}

func (column sortColor) Less(i, j int) bool {
	return get_luminance(column[i]) > get_luminance(column[j])
}

func sort_column(column sortColor) sortColor {
	sort.Sort(column)
	return column
}

func set_column(img *image.NRGBA, column_index int, height int, column sortColor) {
	for i := 0; i < height; i++ {
		img.SetNRGBA(column_index, i, column[i])
	}
}

func process_column(img *image.NRGBA, column_index int, height int) {
	column := get_column(img, column_index, height)
	column = sort_column(column)
	set_column(img, column_index, height, column)
}
