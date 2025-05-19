using Godot;
using System;

public partial class CoordinateLabel : Label
{
  private CityStateManager manager;

  public override void _Ready()
  {
	manager = GetNode<CityStateManager>("%CityStateManager");
  }

  public override void _Process(double _delta)
  {
	Vector2I coord = manager.GetTileFromMouseHover(false);
	if (coord.X < 0 || coord.Y < 0 || coord.X >= CityGameConsts.Consts.WORLD_SIZE.X || coord.Y >= CityGameConsts.Consts.WORLD_SIZE.Y)
	{
	  Visible = false;
	}
	else
	{
	  Visible = true;
	  Text = "(" + coord.X + ", " + coord.Y + ")";
	}
  }


}
