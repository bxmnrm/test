package main
import(
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

// 1️⃣ กรณีข้อมูลครบทุก field (valid)
func TestGamnenameValid(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run("GamenameValid", func(t *testing.T){
		gamename := GameName{
			Title:       "The Legend of Code",
			Genre:       "Adventure",
			Platform:    "PC",
			ReleaseYear: 2022,
			Rating:      9.5,
		}
		ok, err := govalidator.ValidateStruct(gamename)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}

// 2️⃣ กรณี Title เป็นค่าว่าง (invalid)
func TestGamenameTitleNotBlank(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run("GamenameTitleNotBlank", func(t *testing.T){
		gamename := GameName{
			Title :	   "",
			Genre:       "Adventure",
			Platform:    "PC",
			ReleaseYear: 2022,
			Rating:      9.5,
		}
		ok, err := govalidator.ValidateStruct(gamename)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Title is required"))
	})
}

// 3️⃣ กรณี ReleaseYear นอกช่วงที่กำหนด (invalid)
func TestGamenameReleaseYearRange(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run("GamenameReleaseYearRange", func (t *testing.T){
		gamename := GameName{
			Title:       "The Legend of Code",
			Genre:       "Adventure",
			Platform:    "PC",
			ReleaseYear: 1940, //1950-2050
			Rating:      9.5,
		}
		ok, err := govalidator.ValidateStruct(gamename)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Release Year must be between 1950 and 2050"))
	})
}