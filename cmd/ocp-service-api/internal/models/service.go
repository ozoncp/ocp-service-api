package models

import "fmt"

type Service struct {
	Id     uint64
	UserId uint64
	Name   string
	Link   string
}

func (service Service) String() string {
	result := fmt.Sprintf("<Service id=%d, UserId=%d, Name=%s, Link=%s>", service.Id, service.UserId, service.Name, service.Link)
	return result
}
