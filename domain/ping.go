package domain

import (
	"database/sql"
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
)

var startTime time.Time

func init() {
	startTime = time.Now()
}

// Ping - reports service start time and checks db connections
func Ping(dbSessions map[string]interface{}) (string, error) {
	sessions := getConnections(dbSessions)
	err := pingConnections(sessions)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Pong at %s. Service restarted at %s", time.Now(), startTime), nil
}

func getConnections(dbSessions map[string]interface{}) []dbSession {
	var sessions []dbSession

	for k, session := range dbSessions {
		switch k {
		case "mongo":
			sessions = append(sessions, session.(*mgo.Session))
		case "mysql":
			sessions = append(sessions, session.(*sql.DB))
		}
	}
	return sessions
}

func pingConnections(dbSessions []dbSession) error {
	for _, session := range dbSessions {
		err := session.Ping()
		if err != nil {
			return err
		}
	}
	return nil
}

type dbSession interface {
	Ping() error
}

// type mongoSession struct {
// 	*mgo.Session
// }
//
// type mysqlSession struct {
// 	*sql.DB
// }

// func (mgo *mongoSession) ping() error {
// 	err := mgo.Session.Ping()
// 	return err
// }
//
// func (msql *mysqlSession) ping() error {
// 	err := msql.DB.Ping()
// 	return err
// }
