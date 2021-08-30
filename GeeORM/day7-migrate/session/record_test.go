package session

import (
	"testing"

	"github.com/qinchenfeng/HelloGoSevenDays/GeeORM/day7-migrate/log"
)

var (
	user1 = &User{"Tom", 18}
	user2 = &User{"Sam", 25}
	user3 = &User{"Jack", 25}
)

func testRecordInit(t *testing.T) *Session {
	t.Helper()
	s := NewSession().Model(&User{})
	err1 := s.DropTable()
	err2 := s.CreateTable()
	_, err3 := s.Insert(user1, user2)
	if err1 != nil || err2 != nil || err3 != nil {
		t.Fatal("failed init test records")
	}
	return s
}

func TestSession_Insert(t *testing.T) {
	s := testRecordInit(t)
	affected, err := s.Insert(user3)
	if err != nil || affected != 1 {
		t.Fatal("failed to create record")
	}
}

func TestSession_Find(t *testing.T) {
	s := testRecordInit(t)
	var users []User
	if err := s.Find(&users); err != nil || len(users) != 2 {
		t.Fatal("failed to query all")
	}
}

func TestSession_Limit(t *testing.T) {
	s := testRecordInit(t)
	var users []User
	err := s.Limit(1).Find(&users)
	if err != nil || len(users) != 1 {
		t.Fatal("failed to query with limit condition")
	}
}

func TestSession_Update(t *testing.T) {
	s := testRecordInit(t)
	affected, _ := s.Where("Name = ?", "Tom").Update("Age", 30)
	u := &User{}
	_ = s.OrderBy("Age DESC").First(u)

	if affected != 1 || u.Age != 30 {
		t.Fatal("failed to update")
	}
}

func TestSession_DeleteAndCount(t *testing.T) {
	s := testRecordInit(t)
	affected, _ := s.Where("Name = ?", "Tom").Delete()
	count, _ := s.Count()

	if affected != 1 || count != 1 {
		t.Fatal("failed to delete or count")
	}
}

type Account struct {
	ID       int `geeorm:"PRIMARY KEY"`
	Password string
}

func (a *Account) BeforeInsert(s *Session) error {
	log.Info("before inert", a)
	a.ID += 1000
	return nil
}
func (a *Account) AfterQuery(s *Session) error {
	log.Info("after query", a)
	a.Password = "******"
	return nil
}

func TestSession_CallMethod(t *testing.T) {
	s := NewSession().Model(&Account{})
	_ = s.DropTable()
	_ = s.CreateTable()
	_, _ = s.Insert(&Account{1, "123456"}, &Account{2, "qwerty"})

	u := &Account{}
	err := s.First(u)

	if err != nil || u.ID != 1001 || u.Password != "******" {
		t.Fatal("Failed to call hooks after query, got", u)
	}
}
