using CityGame;
using Godot;
using System;
using System.Data;

namespace CityGameConsts
{
	public class BuildingSprites
	{
		public static int SOURCE_ID = 0;
		public static Vector2I FURNACE_0 = new(0, 0);
		public static Vector2I FURNACE_1 = new(1, 0);
		public static Vector2I CRATE = new(2, 0);
		public static Vector2I MINE = new(0, 1);
		public static Vector2I WORKBENCH = new(1, 2);
		public static Vector2I LUMBER = new(0, 3);
		public static Vector2I STONEPILE = new(0, 4);
		public static Vector2I SANDPIT = new(0, 5);
		public static Vector2I LAYER_0 = new(2, 4);
		public static Vector2I LAYER_1 = new(4, 4);
		public static Vector2I LAYER_2 = new(2, 1);
		public static Vector2I LAYER_3 = new(4, 0);
		public static Vector2I LAYER_4 = new(6, 2);
		public static Vector2I LAYER_5 = new(8, 1);

		public static bool BuildingIs2x1(BuildingState building)
		{
			switch (building.BuildingType)
			{
				case Building.Lumber:
					return true;
				case Building.Stonepile:
					return true;
				case Building.Mine:
					return true;
				case Building.Sandpit:
					return true;
				default:
					return false;
			}
		}
		public static Vector2I GetAtlas(BuildingState building)
		{
			switch (building.BuildingType)
			{
				case Building.Workbench:
					return WORKBENCH;
				case Building.Furnace:
					return building.State > 0 ? FURNACE_1 : FURNACE_0;
				case Building.Lumber:
					return LUMBER;
				case Building.Stonepile:
					return STONEPILE;
				case Building.Mine:
					return MINE;
				case Building.Sandpit:
					return SANDPIT;
				case Building.Crate:
					return CRATE;
				default:
					return new(3, 0);
			}
		}

		public static Vector2I GetHouseAtlas(int layer)
		{
			switch (layer)
			{
				case 0:
					return LAYER_0;
				case 1:
					return LAYER_1;
				case 2:
					return LAYER_2;
				case 3:
					return LAYER_3;
				case 4:
					return LAYER_4;
				case 5:
					return LAYER_5;
				default:
					return LAYER_0;
			}
		}
	}

	public class ItemSprites
	{
		public static int SOURCE_ID = 1;
		public static Vector2I HAMMER = new(0, 0);
		public static Vector2I BUCKET = new(1, 0);
		public static Vector2I PLANK = new(2, 0);
		public static Vector2I STONE = new(3, 0);
		public static Vector2I PICKAXE = new(0, 1);
		public static Vector2I ORE = new(1, 1);
		public static Vector2I METAL = new(2, 1);
		public static Vector2I GLASS = new(3, 1);
		public static Vector2I DOOR = new(0, 2);
		public static Vector2I WORKBENCH = new(1, 2);
		public static Vector2I FURNACE = new(2, 2);
		public static Vector2I CRATE = new(3, 2);
	}

	public class Consts
	{
		public static Vector2I HOUSE_COORDINATE = new(19, 6);
	}
}
