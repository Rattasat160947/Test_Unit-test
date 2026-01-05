# Unit Test

# สร้าง Repo 
git clone to Desktop
cd to folder(name repo) ใช้คำสั่ง go mod init foldername

พิมพ์ คำสั่ง go get github.com/asaskevich/govalidator , go get github.com/onsi/gomega
สร้าง folder entity , folder test
>> สร้างไฟล์ user.go in entity > 
> package entity
import "gorm.io/gorm"  < อย่าลืม go mod tidy หรือ go get gorm.io/gorm
type User (ตัวหน้าต้องพิมใหญ่ไม่งั้นเรียกใช้ไม่ได้) struct {
 gorm.Model
 Name string `valid:"required~Name is required"`
}

>> สร้างไฟล์ user_test.go < (ต้องมี _test ต่อท้าย) in test >
> package test
import (
	"testing" //จำเป็น
	"github.com/asaskevich/govalidator" //จำเป็น
	. "github.com/onsi/gomega" //จำเป็น
	"Test_Unit-test/entity" //จำเป็น ต้องตรงกับชื่อ module ที่สั่ง go mod init
)

// เเบบเขียน Test
func TestUser(t *testing.T){
	g := NewGomegaWithT(t)
	t.Run(`Name is required (สำคัญต้องตรง)`, func(t *testing.T){
	user := entity.User{
		Name:"" // ผิดตรงนี้
	}
	ok, err := govalidator.ValidateStruct(user)
	g.Expect(ok).NotTo(BeTrue()) 
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("Name is required")) < ต้องตรง
})

// เเบบถูกทั้งหมดเผื่อพี่ให้เขียน
func TestUser(t *testing.T){
	g := NewGomegaWithT(t)
	t.Run(`Name is valid`, func(t *testing.T){
	user := entity.User{
		Name:"Test"	
	}
	ok, err := govalidator.ValidateStruct(user)
	g.Expect(ok).To(BeTrue()) 
	g.Expect(err).To(BeNil())
})

# run test
> go test -v ./...

# รันแล้วผลไม่อัปเดตล้าง cache ก่อน
> go clean -testcache

# คาดหวังให้ผลลัพธ์เป็นยังไง
g.Expect(ok).To(BeTrue()) > ต้องเป็นจริง
g.Expect(err).To(BeNil()) > ต้องว่างเปล่า no error
g.Expect(ok).NotTo(BeTrue()) > ต้องไม่ผ่าน
g.Expect(err).NotTo(BeNil()) > ต้องมี Error
g.Expect(err.Error()).To(Equal("Name is required")) การเช็คว่า Error Message ตรงกับที่เราเขียนไว้ใน Struct Tag ไหม

# `cheat_sheet.md` สำหรับ GoValidator Cheat Sheet function เรียกใช้เทส