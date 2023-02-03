package main

import (
	model "GLIM_Hacktiv8/golang-intermediate/session-3/introduction/model"
	"log"
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
		// Original
		fmt.Println("# ===== Original\n			%#v \n ", user1)

		// As String
		fmt.Println("# ===== As String\n 		%v	\n", user1.String())
	*/
}
