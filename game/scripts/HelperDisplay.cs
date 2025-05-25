using CityGame;
using Godot;
using Google.Protobuf;
using System;

public partial class HelperDisplay : BoxContainer
{
  public HelperState helperState;
  public HelperInput helperInput;

  private Label helper_id;
  private Label coord;
  private Label items;
  private Label action;
  private Label input_action;
  private Label action_dir;
  private Label input_action_dir;
  private Label action_coord;
  private Label input_action_coord;
  private Label action_item;
  private Label input_action_item;
  private PanelContainer line;
  private VBoxContainer input_info;

  public override void _Ready()
  {
	helper_id = GetNode<Label>("%helper_id");
	coord = GetNode<Label>("%coord");
	items = GetNode<Label>("%items");
	action = GetNode<Label>("%action");
	input_action = GetNode<Label>("%input_action");
	action_dir = GetNode<Label>("%action_dir");
	input_action_dir = GetNode<Label>("%input_action_dir");
	action_coord = GetNode<Label>("%action_coord");
	input_action_coord = GetNode<Label>("%input_action_coord");
	action_item = GetNode<Label>("%action_item");
	input_action_item = GetNode<Label>("%input_action_item");
	line = GetNode<PanelContainer>("%Line");
	input_info = GetNode<VBoxContainer>("%InputInfo");


	if (helperInput == null)
	{
	  UpdateInformation(helperState, false);
	}
	else
	{
	  UpdateInformation(helperState, helperInput, false);
	}
  }

  public void UpdateInformation(HelperState helperState, bool only_if_visible)
  {
	if (!Visible && only_if_visible)
	{
	  return;
	}

	helper_id.Text = "Helper ID: " + helperState.HelperId;
	coord.Text = "Coordinate: (" + helperState.Coordinate.X + ", " + helperState.Coordinate.Y + ")";
	items.Text = "Items: ";
	if (helperState.Items.Count == 0)
	{
	  items.Text += "None";
	}
	else
	{
	  for (int i = 0; i < helperState.Items.Count; i++)
	  {
		Stack stack = helperState.Items[i];
		if (i > 0)
		{
		  items.Text += ", ";
		}
		items.Text += stack.Count + " " + stack.ItemId.ToString();
	  }
	}

	action.Text = "Action: " + helperState.Action.ActionType.ToString();
	bool found_attr;

	found_attr = helperState.Action.HasDirection;
	action_dir.Visible = found_attr;
	if (found_attr)
	{
	  action_dir.Text = "Action Dir: " + helperState.Action.Direction.ToString();
	}

	found_attr = helperState.Action.Coordinate != null;
	action_coord.Visible = found_attr;
	if (found_attr)
	{
	  action_coord.Text = "Action Coord: (" + helperState.Action.Coordinate.X + ", " + helperState.Action.Coordinate.Y + ")";
	}
	found_attr = helperState.Action.HasItemId;
	action_item.Visible = found_attr;
	if (found_attr)
	{
	  action_item.Text = "Action Item: " + helperState.Action.ItemId.ToString();
	}

	line.Visible = false;
	input_info.Visible = false;
  }
  public void UpdateInformation(HelperState helperState, HelperInput helperInput, bool only_if_visible)
  {
	if (!Visible && only_if_visible)
	{
	  return;
	}

	UpdateInformation(helperState, only_if_visible);
	line.Visible = true;
	input_info.Visible = true;

	GD.Print(helperInput.Action);

	input_action.Text = "Action: " + helperInput.Action.ActionType.ToString();
	bool found_attr;

	found_attr = helperInput.Action.HasDirection;
	input_action_dir.Visible = found_attr;
	if (found_attr)
	{
	  input_action_dir.Text = "Action Dir: " + helperInput.Action.Direction.ToString();
	}
	found_attr = helperInput.Action.Coordinate != null;
	input_action_coord.Visible = found_attr;
	if (found_attr)
	{
	  input_action_coord.Text = "Action Coord: (" + helperInput.Action.Coordinate.X + ", " + helperInput.Action.Coordinate.Y + ")";
	}
	found_attr = helperInput.Action.HasItemId;
	input_action_item.Visible = found_attr;
	if (found_attr)
	{
	  input_action_item.Text = "Action Item: " + helperInput.Action.ItemId.ToString();
	}
  }
}
