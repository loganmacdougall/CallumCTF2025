package consts

import (
	"city_game/src/pb"
)

var layer_requirements [][]*pb.ItemRequirement

func GetLayerRequirements(layer int) []*pb.ItemRequirement {
	initialize_layers()
	return layer_requirements[layer]
}

func initialize_layers() {
	if layer_requirements != nil {
		return
	}

	layer_requirements = [][]*pb.ItemRequirement{
		{
			{
				ItemId: pb.Item_Stone,
				Count:  50,
			},
		},
		{
			{
				ItemId: pb.Item_Stone,
				Count:  50,
			},
			{
				ItemId: pb.Item_Plank,
				Count:  50,
			},
		},
		{
			{
				ItemId: pb.Item_Stone,
				Count:  50,
			},
			{
				ItemId: pb.Item_Plank,
				Count:  50,
			},
			{
				ItemId: pb.Item_Door,
				Count:  5,
			},
			{
				ItemId: pb.Item_Metal,
				Count:  5,
			},
		},
		{
			{
				ItemId: pb.Item_Stone,
				Count:  100,
			},
			{
				ItemId: pb.Item_Plank,
				Count:  100,
			},
			{
				ItemId: pb.Item_Door,
				Count:  10,
			},
			{
				ItemId: pb.Item_Metal,
				Count:  10,
			},
			{
				ItemId: pb.Item_Window,
				Count:  4,
			},
		},
		{
			{
				ItemId: pb.Item_Stone,
				Count:  150,
			},
			{
				ItemId: pb.Item_Plank,
				Count:  100,
			},
			{
				ItemId: pb.Item_Door,
				Count:  10,
			},
			{
				ItemId: pb.Item_Metal,
				Count:  20,
			},
			{
				ItemId: pb.Item_Window,
				Count:  4,
			},
		},
		{
			{
				ItemId: pb.Item_Stone,
				Count:  250,
			},
			{
				ItemId: pb.Item_Plank,
				Count:  150,
			},
			{
				ItemId: pb.Item_Metal,
				Count:  50,
			},
		},
	}
}
