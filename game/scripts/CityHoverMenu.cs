using CityGame;
using Godot;
using System.Collections.Generic;
using System.Runtime.CompilerServices;

public partial class CityHoverMenu : VBoxContainer
{
  private Vector2I currentCoord = new Vector2I(-1, -1);
  private CityStateManager manager;
  private string helperDisplayScenePath = "res://components/helper_display_ui.tscn";
  private string buildingDisplayScenePath = "res://components/building_display_ui.tscn";
  private PackedScene helperDisplayScene;
  private PackedScene buildingDisplayScene;

  public override void _Ready()
  {
	manager = GetNode<CityStateManager>("%CityStateManager");
	helperDisplayScene = GD.Load<PackedScene>(helperDisplayScenePath);
	buildingDisplayScene = GD.Load<PackedScene>(buildingDisplayScenePath);
	HideWithoutClear();
  }


  public override void _Process(double _delta)
  {
	if (!Input.IsActionJustPressed("Click"))
	{
	  return;
	}

	GD.Print("Mouse click detected");

	if (!manager.IsValidTileFromMouseHover())
	{
	  return;
	}

	GD.Print("Valid tile detected");

	Vector2I tile = manager.GetTileFromMouseHover(true);

	if (tile == currentCoord)
	{
	  HideWithoutClear();
	  return;
	}

	Coordinate coord = new Coordinate();
	coord.X = (uint)tile.X;
	coord.Y = (uint)tile.Y;

	TickData tickData = manager.GetCurrentTickData();

	UpdateInformation(tickData, coord);
  }


  public void UpdateInformation(TickData tickData, Coordinate coord)
  {
	List<HelperState> helpers = new List<HelperState>();
	BuildingState buildingState = null;
	Dictionary<uint, HelperInput> helperInputs = new Dictionary<uint, HelperInput>();

	foreach (BuildingState building in tickData.State.BuildingStates)
	{
	  if (building.Coordinate.X == coord.X && building.Coordinate.Y == coord.Y)
	  {
		buildingState = building;
		break;
	  }
	}

	if (buildingState == null)
	{

	  foreach (HelperState helper in tickData.State.HelperStates)
	  {
		if (helper.Coordinate.X == coord.X && helper.Coordinate.Y == coord.Y)
		{
		  helpers.Add(helper);
		}
	  }

	  foreach (HelperInput helperInput in tickData.Input.HelperInput)
	  {
		helperInputs.Add(helperInput.HelperId, helperInput);
	  }

	  if (helpers.Count == 0)
	  {
		HideWithoutClear();
		return;
	  }
	}

	ClearAll();
	Visible = true;

	if (buildingState != null)
	{
	  BuildingDisplay building_display = buildingDisplayScene.Instantiate<BuildingDisplay>();
	  building_display.buildingState = buildingState;
	  AddChild(building_display);
	}
	else
	{
	  foreach (HelperState helper in helpers)
	  {
		HelperDisplay helper_display = helperDisplayScene.Instantiate<HelperDisplay>();
		helper_display.helperState = helper;
		if (helperInputs.ContainsKey(helper.HelperId))
		{
		  helper_display.helperInput = helperInputs[helper.HelperId];
		}
		AddChild(helper_display);
	  }
	}
  }

  public void HideWithoutClear()
  {
	Visible = false;
	currentCoord = new Vector2I(-1, -1);
  }

  void ClearAll()
  {
	foreach (Node node in GetChildren())
	{
	  node.QueueFree();
	}
  }

}
