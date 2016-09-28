package domain_test

import (
	"time"

	"github.com/Shop2market/goping/domain"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/mgo.v2"
)

var _ = Describe("Ping", func() {
	var Session *mgo.Session
	Session, _ = mgo.Dial("127.0.0.1")
	frozenTime := time.Date(2010, time.September, 26, 23, 0, 0, 0, time.Local)
	dbSessions := map[string]interface{}{
		"mongo": Session,
	}

	BeforeEach(func() {
		domain.Now = func() time.Time {
			return frozenTime
		}
	})

	It("Should return service restart time even if db connection is inactive", func() {
		response := domain.Ping(map[string]interface{}{})
		Expect(response.Err).NotTo(HaveOccurred())
		Expect(response.PongAt).To(Equal(frozenTime))
		Expect(response.RestartAt).To(Equal(domain.StartTime))
	})

	It("Should return service restart time after checking db connections for request", func() {
		response := domain.Ping(dbSessions)
		Expect(response.Err).NotTo(HaveOccurred())
		Expect(response.PongAt).To(Equal(frozenTime))
		Expect(response.RestartAt).To(Equal(domain.StartTime))
		Session.Close()
	})

	// It("Should return mongo connection error after checking db connections for request", func() {
	// 	Session.Close()
	// 	startTime, err := domain.Ping(dbSessions)
	// 	Expect(err).To(HaveOccurred())
	// 	Expect(err).To(Equal("Mongodb connection failed to establish"))
	// 	Expect(startTime).To(Equal(nil))
	// })

	// It("Should return mysql connection error after checking db connections for request", func() {
	// 	startTime, err := domain.Ping(dbSessions)
	// 	Expect(err).To(HaveOccurred())
	// 	Expect(err).To(Equal("Mysql connection failed to establish"))
	// 	Expect(startTime).To(Equal(nil))
	// })
})
