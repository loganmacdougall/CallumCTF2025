package consts

import (
	"city_game/src/pb"
	"sync"
)

var (
	layer_requirements           [][]*pb.ItemRequirement
	init_once_layer_requirements sync.Once
)

func GetLayerRequirements(layer int) []*pb.ItemRequirement {
	init_once_crafting_recipes.Do(initialize_layers)
	return layer_requirements[layer]
}

func initialize_layers() {
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
