package entity

type Cal struct{
	Number2 string `valid:"numeric~Number2 must be negative"`
	Number3 string `valid:"alpha~Number3 is not zero"`
	Number4 string `valid:"lowercase~Number4 must be non-positive"`
}
