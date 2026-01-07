package entity

type Booking struct{
	UserEmail string `valid:"email~Invalid email format"`
	MeetingLink string `valid:"url~Invalid URL"`
	Topic string `valid:"required~Topic is required"`
	Participants int `valid:"range(2|50)~Participants must be between 2 and 50"`
}