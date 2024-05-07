package controller_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"monster/db"
	"net/http"

	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Monster", func() {
	_ = godotenv.Load()
	db := db.Connect()

	var postJson []byte

	BeforeEach(func() {
		_, err := db.Exec("DELETE FROM monsters;")
		if err != nil {
			panic(fmt.Errorf("failed to delete monsters. %w", err))
		}

		postJson = []byte(`{
				"name": "Purple Horse",
				"attack": 40, 
				"defense": 50, 
				"hp": 14, 
				"speed": 16
			}`)
	})

	JustBeforeEach(func() {
		req, _ := http.NewRequest(http.MethodPost, "http://localhost:9000/monsters", bytes.NewBuffer(postJson))

		_, err := http.DefaultClient.Do(req)

		if err != nil {
			panic(fmt.Errorf("Unable to create monster. %w", err))
		}
	})

	AfterEach(func() {
		_, err := db.Exec("DELETE FROM monsters;")
		if err != nil {
			panic(fmt.Errorf("failed to delete monsters. %w", err))
		}
	})

	Describe("List", func() {
		var response *http.Response

		JustBeforeEach(func() {
			req, err := http.NewRequest(http.MethodGet, "http://localhost:9000/monsters/1", nil)
			response, err = http.DefaultClient.Do(req)
			if err != nil {
				panic(fmt.Errorf("failed to get monsters. %w", err))
			}
		})

		Context("should list all monsters", func() {

			It("status code should be 200", func() {
				Expect(response.StatusCode).To(Equal(200))
			})

			It("body should not be nil", func() {
				Expect(response.Body).ToNot(BeNil())
			})

			It("body should have equivalent values", func() {
				var m map[string]interface{}
				var l []map[string]interface{}
				bodyString, _ := io.ReadAll(response.Body)
				if err := json.Unmarshal([]byte(bodyString), &l); err != nil {
					panic(fmt.Errorf("failed to deserialize. %w", err))
				}
				Expect(len(l)).To(Equal(1))
				for _, m = range l {
					Expect(m["name"]).To(Equal("Dead Unicorn"))
					Expect(m["attack"]).To(BeEquivalentTo(60))
					Expect(m["defense"]).To(BeEquivalentTo(40))
					Expect(m["hp"]).To(BeEquivalentTo(10))
					Expect(m["speed"]).To(BeEquivalentTo(80))
				}
			})

		})

	})

})
