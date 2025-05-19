using Godot;
using Google.Protobuf.WellKnownTypes;
using System;

public partial class SaveTheCityPlayback : Node2D
{
	private Sprite2D playbackButtonSprite;
	private HSlider playbackSpeedSlider;
	private Label playbackSpeedLabel;

	private Texture2D playTexture;
	private Texture2D pauseTexture;

	private CityStateManager manager;

	private bool playing = false;
	private int tick
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
	private double time = 0;
	private int ticks_per_second
	{
		get
		{
			return (int)playbackSpeedSlider.Value;
		}
	}
	private double seconds_per_tick
	{
		get
		{
			return 1.0 / ticks_per_second;
		}
	}

	public override void _Ready()
	{
		playbackButtonSprite = GetNode<Sprite2D>("%PlaybackButtonSprite");
		playbackSpeedSlider = GetNode<HSlider>("%PlaybackSpeed");
		playbackSpeedLabel = GetNode<Label>("%PlaybackSpeedLabel");
		playbackSpeedLabel.Text = ticks_per_second + " t/s";

		playTexture = GD.Load<Texture2D>("res://assets/sprites/playback_play.png");
		pauseTexture = GD.Load<Texture2D>("res://assets/sprites/playback_pause.png");

		manager = GetNode<CityStateManager>("%CityStateManager");
	}

	public override void _Process(double delta)
	{
		if (playing)
		{
			time += delta;
			if (time > seconds_per_tick)
			{
				int ticks_past = (int)(time / seconds_per_tick);
				tick += ticks_past;
				time -= ticks_past * seconds_per_tick;
				manager.SetTick(tick);

				if (tick >= manager.Count() - 1)
				{
					playing = false;
				}
			}
		}
	}

	public void OnPlaybackButtonPressed()
	{
		TogglePlayback();
	}

	public void TogglePlayback()
	{
		playing = !playing;

		if (playing)
		{
			playbackButtonSprite.Texture = pauseTexture;
			time = 0;
		}
		else
		{
			playbackButtonSprite.Texture = playTexture;
		}
	}

	public void OnStepLeftButtonPressed()
	{
		if (playing) TogglePlayback();
		tick -= 1;
	}
	public void OnStepRightButtonPressed()
	{
		if (playing) TogglePlayback();
		tick += 1;
	}

	public void OnPlaybackSpeedValueChanged(float value)
	{
		playbackSpeedLabel.Text = value + " t/s";
	}

}
