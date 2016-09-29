package goping_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	. "github.com/Shop2market/goping"
	"github.com/gorilla/handlers"
	"github.com/julienschmidt/httprouter"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Goping", func() {
	BeforeEach(func() {
		router := httprouter.New()
		router.GET("/pigable-service/ping", Ping())
		go http.ListenAndServe(fmt.Sprintf(":8889"), handlers.LoggingHandler(os.Stdout, router))
	})
	Context("Ping server", func() {
		It("renders ping information", func() {
			resp, err := http.Get("http://localhost:8889/pigable-service/ping")
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			decoder := json.NewDecoder(resp.Body)
			pingData := map[string]interface{}{}
			Expect(decoder.Decode(&pingData)).NotTo(HaveOccurred())
			Expect(pingData["Pong At"]).NotTo(BeNil())
		})
	})
})
