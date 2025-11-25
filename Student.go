package main

import (
	"gorm.io/gorm"
)

type Student struct {
    gorm.Model
    FullName string  `valid:"required~FullName is required"`
    Age      uint    `valid:"required~Age is required,range(18|100)~Age must be at least 18"`
    Email    string  `valid:"required~Email is required,email~Email is invalid"`
    GPA      float32 `valid:"range(0|4)~GPA must be between 0.00 and 4.00"`
}
