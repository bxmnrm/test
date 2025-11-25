package main
import (
	"testing"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

// 1️⃣ กรณีข้อมูลครบทุก field (valid)
func TestStudentValid(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run("StudentValid", func(t *testing.T) {
		student := Student{
			FullName: "John Doe",
			Age:      20,
			Email:    "john.doe@example.com",
			GPA:      3.5,
		}
		ok, err := govalidator.ValidateStruct(student)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}

// 2️⃣ กรณี FullName เป็นค่าว่าง
func TestStudentNoFullName(t *testing.T) {
	g := NewGomegaWithT(t)
	student := Student{
		FullName: "",
		Age:      20,
		Email:    "john.doe@example.com",
		GPA:      3.5,
	}
	ok, err := govalidator.ValidateStruct(student)
	g.Expect(ok).To(BeFalse())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("FullName is required"))
}

// 3️⃣ กรณีอายุน้อยกว่า 18 ปี
func TestStudentUnderAge(t *testing.T) {
	g := NewGomegaWithT(t)
	student := &Student{
		FullName: "John Doe",
		Age:      16,
		Email:    "john.doe@example.com",
		GPA:      3.5,
	}
	ok, err := govalidator.ValidateStruct(student)
	g.Expect(ok).To(BeFalse())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("Age must be 18 or older"))
}

// 4️⃣ กรณีอีเมลไม่ถูกต้อง
func TestStudentInvalidEmail(t *testing.T) {
	g := NewGomegaWithT(t)
	student := &Student{
		FullName: "John Doe",
		Age:      20,
		Email:    "invalid-email",
		GPA:      3.5,
	}
	ok, err := govalidator.ValidateStruct(student)
	g.Expect(ok).To(BeFalse())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("Email is invalid"))
}

// 5️⃣ กรณี GPA นอกช่วง 0.00 - 4.00
func TestStudentInvalidGPA(t *testing.T) {
	g := NewGomegaWithT(t)
	student := &Student{
		FullName: "John Doe",
		Age:      20,
		Email:    "john.doe@example.com",
		GPA:      4.5,
	}
	ok, err := govalidator.ValidateStruct(student)
	g.Expect(ok).To(BeFalse())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("GPA must be between 0.00 and 4.00"))
}