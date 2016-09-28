package domain_test

import (
	"time"

	"github.com/Shop2market/goping/domain"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type testType struct{}

func (test *testType) Ping() error {
	return nil
}

var _ = Describe("Ping", func() {
	frozenTime := time.Date(2010, time.September, 26, 23, 0, 0, 0, time.Local)
	frozenTimePtr := &frozenTime
	dbSessions := map[string]interface{}{
		"mongo": &testType{},
	}

	It("Should return service restart time after checking db connections for request", func() {
		startTime, err := domain.Ping(dbSessions)
		Expect(err).NotTo(HaveOccurred())
		Expect(startTime).To(Equal(frozenTime))
	})

	It("Should return mysql connection error after checking db connections for request", func() {
		startTime, err := domain.Ping(dbSessions)
		Expect(err).To(HaveOccurred())
		Expect(err).To(Equal("Mysql connection failed to establish"))
		Expect(startTime).To(Equal(nil))
	})

	It("Should return mongo connection error after checking db connections for request", func() {
		startTime, err := domain.Ping(dbSessions)
		Expect(err).To(HaveOccurred())
		Expect(err).To(Equal("Mongodb connection failed to establish"))
		Expect(startTime).To(Equal(nil))
	})

	It("Should return service restart time even if db connection is inactive", func() {
		startTime, err := domain.Ping(dbSessions)
		Expect(err).NotTo(HaveOccurred())
		Expect(startTime).To(Equal(frozenTime))
	})
})
