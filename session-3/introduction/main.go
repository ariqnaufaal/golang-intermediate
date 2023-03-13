package main

import (
	model "GLIM_Hacktiv8/golang-intermediate/session-3/introduction/model"
	"bytes"
	"log"

	"github.com/golang/protobuf/jsonpb"
)

func main() {
	// import dari model proto
	var user1 = &model.User{
		Id:       "u001",
		Name:     "Sylvana Windrunner",
		Password: "f0r Th3 H0rD3",
		Gender:   model.UserGender_FEMALE,
	}

	var userList = &model.UserList{
		List: []*model.User{
			user1,
		},
	}

	log.Println("userList", userList)

	log.Printf("user1 %#v\n", user1)
	log.Printf("user1.String() %#v\n", user1.String())

	// tes jsonpb marshal (marshalling proto to json)
	var (
		buf bytes.Buffer
	)
	_ = (&jsonpb.Marshaler{}).Marshal(&buf, user1)
	log.Printf("user1.jsonString %#v\n", buf.String())

	var garage1 = &model.Garage{
		Id:   "g001",
		Name: "Kalimdar",
		Coordinate: &model.GarageCoordinate{
			Latitude:  23.2212847,
			Longitude: 53.22033123,
		},
	}

	var garageList = &model.GarageList{
		List: []*model.Garage{
			garage1,
		},
	}

	var garageListByUser = &model.GarageListByUser{
		List: map[string]*model.GarageList{
			user1.Id: garageList,
		},
	}

	log.Println("garageListByUser", garageListByUser)

	/*
		// =========== Original
		fmt.Println("# ===== Original\n			%#v \n ", user1)

		// =========== As String
		fmt.Println("# ===== As String\n 		%v	\n", user1.String())

		// =========== as json string
	*/

	/*
		var buf bytes.Buffer
		err1 := (&jsonpb.Marshaler{}).Marshal(&buf, garageList)
		if err1 != nil {
			fmt.Println(err1.Error())
			os.Exit(0)
		}
		jsonString := buf.String()
		fmt.Printf("# === As JSON String \n		&v \n", jsonString)
	*/
}
