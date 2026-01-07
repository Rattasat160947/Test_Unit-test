package entity

type Library struct{
	Title string `valid:"required~Title is required"`
	Author string `valid:"required~Author is required"`
	Page int `valid:"required~Page is required, range(1|1000)~Page must be between 1 and 1000"`
	ISBN string `valid:"matches(^ISBN-\\d{4}$)~ISBN invalid format"`
}
