package main

import "gorm.io/gorm"

type GameName struct {
	gorm.Model
	Title       string  `valid:"required~Title is required"`
	Genre       string  `valid:"required~Genre is required"`
	Platform    string  `valid:"required~Platform is required"`
	ReleaseYear int     `valid:"range(1950|2050)~Release Year must be between 1950 and 2050"`
	Rating      float32 `valid:"range(0|10)~Rating must be between 0 and 10"`
}
