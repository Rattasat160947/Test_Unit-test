package entity

import (
	"gorm.io/gorm"
)

type ExampleSpec struct {
	gorm.Model

	// 1. range(min|max): ตรวจสอบค่าตัวเลข (Int/Float)
	// ต้องอยู่ระหว่าง 18 ถึง 99
	Age int `valid:"range(18|99)~Age must be between 18 and 99"`

	// 2. stringlength(min|max): ตรวจสอบความยาวตัวอักษร (นับตามตัวอักษร)
	// ยาว 4-10 ตัวอักษร (ภาษาไทยนับเป็น 1 ตัวปกติ)
	Username string `valid:"stringlength(4|10)~Username length must be 4-10"`

	// 3. length(min|max): ตรวจสอบ Byte (ระวัง! ภาษาไทย 1 ตัว = 3 bytes)
	// จำกัด 10-20 bytes (ไม่ค่อยแนะนำให้ใช้กับ text ที่คนพิมพ์ เว้นแต่จะจำกัดขนาด data)
	RawData string `valid:"length(10|20)~Data byte size must be 10-20"`

	// 4. runelength(min|max): ตรวจสอบความยาวตัวอักษร (เหมือน stringlength)
	// เหมาะกับภาษาไทยมากที่สุด (ก = 1 ตัว)
	Nickname string `valid:"runelength(2|10)~Nickname must be 2-10 chars"`

	// 5. matches(pattern): ตรวจสอบ Regex
	// ต้องขึ้นต้นด้วย S ตามด้วยเลข 8 หลัก (เช่น S12345678)
	StudentID string `valid:"matches(^S\\d{8}$)~Invalid Student ID format"`

	// 6. in(...): ตรวจสอบค่าที่ยอมรับ (Enum)
	// ต้องเป็น "admin", "user", หรือ "guest" เท่านั้น
	Role string `valid:"in(admin|user|guest)~Role must be admin user or guest"`

	// 7. minstringlength(int): ความยาวขั้นต่ำ
	// อย่างน้อย 5 ตัว
	Password string `valid:"minstringlength(5)~Password too short"`

	// 8. maxstringlength(int): ความยาวสูงสุด
	// ห้ามเกิน 100 ตัว
	Description string `valid:"maxstringlength(100)~Description too long"`

}