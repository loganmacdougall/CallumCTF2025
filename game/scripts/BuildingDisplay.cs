using CityGame;
using Godot;
using System;

public partial class BuildingDisplay : BoxContainer
{
  private Label building_type;
  private Label coord;
  private Label items;
  private Label state;

  public BuildingState buildingState;

  public override void _Ready()
  {
	building_type = GetNode<Label>("%building_type");
	coord = GetNode<Label>("%coord");
	items = GetNode<Label>("%items");
	state = GetNode<Label>("%state");

	UpdateInformation(buildingState);
  }

  void UpdateInformation(BuildingState buidingState)
  {
	building_type.Text = "Building: " + buildingState.BuildingType.ToString();
	coord.Text = "Coordinate: (" + buidingState.Coordinate.X + ", " + buidingState.Coordinate.Y + ")";
	items.Text = "Items: ";
	if (buidingState.Items.Count == 0)
	{
	  items.Text += "None";
	}
	else
	{
	  for (int i = 0; i < buidingState.Items.Count; i++)
	  {
		Stack stack = buidingState.Items[i];
		if (i > 0)
		{
		  items.Text += ", ";
		}
		items.Text += stack.Count + " " + stack.ItemId.ToString();
	  }
	}
	state.Text = "State: " + buidingState.State;
  }



}
