package main

import (
	"errors"
	"fmt"
)

type setStruct struct {
	Elements map[string]struct{}
}

func Set()  *setStruct{
	var set setStruct
	set.Elements = make(map[string]struct{})
	return &set
}

func (s *setStruct) Add(elem string){
	s.Elements[elem] = struct{}{}
}

func (s *setStruct) Delete(elem string) error{
	if _, exists := s.Elements[elem]; !exists {
		return errors.New("element not present in set")
	}
	delete(s.Elements, elem)
	return nil
}

func (s *setStruct) Exists(elem string) bool{
	_, exists := s.Elements[elem]
	return exists
}

func (s *setStruct) List() {
	for k, _ := range s.Elements {
		fmt.Println(k)
	}
}

func main() {
	mySet := Set()
	mySet.Add("Real Madrid")
	mySet.Add("Abahoni")
	mySet.Add("Barcelona")
	mySet.Add("Chelsea")
	mySet.Add("Chelsea")

	mySet.List()
	mySet.Delete("Barcelona")
	fmt.Println(mySet.Exists("Barcelona"))

}
