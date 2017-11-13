package storage

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/davecgh/go-spew/spew"
	"github.com/jehaby/webapp102/entity"
	"golang.org/x/crypto/bcrypt"
)

type Memory struct {
	l sync.RWMutex
	s map[string]entity.User
}

func NewMemory() *Memory {
	m := &Memory{
		s: make(map[string]entity.User, 0),
	}

	pass, err := bcrypt.GenerateFromPassword([]byte("111"), bcrypt.DefaultCost)
	if err != nil {
		log.Panic(err)
	}

	log.Println("encrypted pass: ", string(pass))

	m.s["urf"] = entity.User{
		Name:    "urf",
		Password: string(pass),
		Email:    "jjj@ya.ru",
	}

	return m
}

func (m *Memory) GetUser(creds entity.Credentials) (entity.User, error) {
	m.l.RLock()
	defer m.l.RUnlock()

	user, ok := m.s[creds.Name]
	if !ok {
		return entity.User{}, &ErrNotFound{"user", map[string]string{"login": creds.Name}}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		spew.Dump("bcrypt error", err)
		return entity.User{}, &ErrBadPassword{creds}
	}

	return user, nil
}

type ErrNotFound struct {
	entity string
	data   map[string]string
}

func (e *ErrNotFound) Error() string {
	var tmp []string
	for k, v := range e.data {
		tmp = append(tmp, fmt.Sprintf("%s: %s", k, v))
	}
	return fmt.Sprintf("entity '%s' not found by data: %s", e.entity, strings.Join(tmp, " | "))
}

type ErrBadPassword struct {
	entity.Credentials
}

func (e *ErrBadPassword) Error() string {
	return fmt.Sprintf("bad password for user: %s", e.Name)
}
