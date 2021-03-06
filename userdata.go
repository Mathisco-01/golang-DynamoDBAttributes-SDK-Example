package main

import (
	"encoding/json"
	"log"
)

// All names are generated by a random name generator!
// I'm not doxxing anybody here!
const data = `[{"userId": 1, "name": "Kenneth G Thomas", "age": 38, "location": "Silver Spring", "friendId":[3,4,5,8], "likes":23},
      {"userId": 2, "name": "Marianne C Reagan", "age": 24, "location": "Gulfport", "friendId":[1,5,9], "likes":12},
      {"userId": 3, "name": "Joy R Crawford", "age": 17, "location": "Jacksonville", "friendId":[1,2,4,5], "likes":120},
      {"userId": 4, "name": "Billy J Shelton", "age": 36, "location": "Poughkeepsie", "friendId":[7,8], "likes":81},
      {"userId": 5, "name": "Tyler J Define", "age": 97, "location": "Los Gatos", "friendId":[1,4,6], "likes":44},
      {"userId": 6, "name": "Mary L Maddox", "age": 15, "location": "Las Vegas", "friendId":[2,8,9], "likes":31},
      {"userId": 7, "name": "Stephanie M Pauls", "age": 67, "location": "Knoxville", "friendId":[], "likes":0},
      {"userId": 8, "name": "Hector R Hall", "age": 41, "location": "New Bern", "friendId":[3,9], "likes":17},
      {"userId": 9, "name": "Bessie J Graded", "age": 28, "location": "Nashville", "friendId":[1,7], "likes":30}]`

// Returns an unmarshalled version of the data constant
func getData() (members []Member) {
	err := json.Unmarshal([]byte(data), &members)
	if err != nil {
		log.Panic(err)
	}
	return members
}
