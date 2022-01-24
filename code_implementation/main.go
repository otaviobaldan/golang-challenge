package main

import (
	"errors"
	"fmt"
	"time"
)

var storedUser = make(map[string]User)

type UserRepository interface {
	GetUser(id string) (User, error)
}

type User struct {
	Id      string
	Friends []string
}

func main() {
	storedUser["A"] = User{Id: "A", Friends: []string{"B", "H"}}
	storedUser["B"] = User{Id: "B", Friends: []string{"A", "H"}}
	storedUser["C"] = User{Id: "C", Friends: []string{"D", "F"}}
	storedUser["D"] = User{Id: "D", Friends: []string{"C", "H"}}
	storedUser["E"] = User{Id: "E", Friends: []string{"H"}}
	storedUser["F"] = User{Id: "F", Friends: []string{"C", "G", "H"}}
	storedUser["G"] = User{Id: "G", Friends: []string{"F", "H"}}
	storedUser["H"] = User{Id: "H", Friends: []string{"A", "B", "D", "G", "F", "E"}}

	fmt.Println(FindAllSocialCircles([]string{"H"}))
}

func FindAllSocialCircles(userIds []string) map[string][]string {
	var result = make(map[string][]string, len(userIds))
	//friendsRelations := make(chan string)
	for i := range userIds {
		var relations = make([]string, 0)
		benchTime := time.Now()
		//wg := &sync.WaitGroup{}
		relations = FindUsersRelations(userIds[i], userIds[i], relations)
		//go func() {
		//	defer wg.Done()

		//wg.Wait()
		//}()
		//result[userIds[i]] = strings.Split(<-friendsRelations, ",")
		result[userIds[i]] = relations
		fmt.Println("Time spent", time.Now().Sub(benchTime))
	}
	return result
}

func FindUsersRelations(userId, mainUser string, relations []string) []string {
	user, err := GetUser(userId)
	if err != nil {
		fmt.Println("[ERROR] ", err.Error())
		panic(err)
	}

	for i := range user.Friends {
		var exists bool
		relations, exists = sliceContains(user.Friends[i], mainUser, relations)
		if !exists {
			relations = FindUsersRelations(user.Friends[i], mainUser, relations)
		}
	}
	return relations
}

func sliceContains(userToAdd, mainUser string, relations []string) ([]string, bool) {
	if userToAdd == mainUser {
		return relations, true
	}
	for _, currentId := range relations {
		if currentId == userToAdd {
			return relations, true
		}
	}
	relations = append(relations, userToAdd)
	return relations, false
}

func GetUser(id string) (User, error) {
	var value User
	var exists bool

	value, exists = storedUser[id]
	if !exists {
		return User{}, errors.New("user not found")
	}

	return value, nil
}


