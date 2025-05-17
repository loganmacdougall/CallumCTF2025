using Godot;
using System;

public partial class TickSlider : HSlider
{
	private Label tickLabel;
	private CityStateManager manager;

	private int current_tick
	{
		get
		{
			return manager.GetTick();
		}
		set
		{
			manager.SetTick(value);
		}
	}

	public override void _Ready()
	{
		tickLabel = GetNode<Label>("%TickLabel");
		manager = GetNode<CityStateManager>("%CityStateManager");

		MinValue = 0;
		MaxValue = manager.Count() - 1;
		Rounded = true;

		UpdateSlider();
	}

	public override void _Process(double delta)
	{
		if (Value != current_tick)
		{
			UpdateSlider();
		}
	}

	private void UpdateSlider()
	{
		Value = current_tick;
		tickLabel.Text = "Tick: " + current_tick;
	}

	private void OnValueChanged(float value)
	{
		if (Value != current_tick)
		{
			current_tick = (int)value;
			tickLabel.Text = "Tick: " + current_tick;
		}
	}

}
