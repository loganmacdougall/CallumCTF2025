using CityGame;
using Godot;
using Google.Protobuf;
using System;
using System.Collections.Generic;
using System.Linq;

public partial class CityStateManager : Node
{
  private TileMapLayer buildingLayer;
  private TileMapLayer helperLayer;
  private TileMapLayer itemLayer;

  private GameData data;

  private int tick = 1;

  [Export(PropertyHint.File, "*.gamedata")]
  public string gameDataFile { get; set; } = "";

  public override void _Ready()
  {
    buildingLayer = GetNode<TileMapLayer>("%buildings");
    helperLayer = GetNode<TileMapLayer>("%helpers");
    itemLayer = GetNode<TileMapLayer>("%items");

    using var file = FileAccess.Open(gameDataFile, FileAccess.ModeFlags.Read);
    byte[] byte_data = file.GetBuffer((long)file.GetLength());

    data = GameData.Parser.ParseFrom(byte_data);

    SetTick(0);
  }

  public void SetTick(int next_tick)
  {
    int old_tick = tick;
    tick = Math.Clamp(next_tick, 0, data.Data.Count - 1);

    if (tick != old_tick)
    {
      UpdateBuildings();
      UpdateHelpers();
    }
  }

  public int GetTick()
  {
    return tick;
  }

  public int Count()
  {
    return data.Data.Count();
  }

  public void UpdateBuildings()
  {
    TickData tick_data = data.Data[tick];
    GameState state = tick_data.State;
    int source_id = CityGameConsts.BuildingSprites.SOURCE_ID;

    buildingLayer.Clear();

    System.Collections.ArrayList coords2x1Tile = new System.Collections.ArrayList();
    foreach (BuildingState buildingState in state.BuildingStates)
    {
      if (buildingState.BuildingType == Building.House) continue;
      PlaceBuilding(buildingState, coords2x1Tile);
    }

    foreach (Vector2I coord in coords2x1Tile)
    {
      if (coord.X == 0) continue;

      Vector2I atlas_coord = buildingLayer.GetCellAtlasCoords(coord);
      Vector2I left_atlas_coord = buildingLayer.GetCellAtlasCoords(coord + Vector2I.Left);

      if (atlas_coord == left_atlas_coord)
      {
        buildingLayer.SetCell(coord, source_id, atlas_coord + Vector2I.Right);
      }

    }

    Vector2I house_coord = CityGameConsts.Consts.HOUSE_COORDINATE;
    Vector2I house_atlas = CityGameConsts.BuildingSprites.GetHouseAtlas((int)state.Layer);
    buildingLayer.SetCell(house_coord, CityGameConsts.BuildingSprites.SOURCE_ID, house_atlas);
  }

  public void PlaceBuilding(BuildingState building, System.Collections.ArrayList coords2x1Tile)
  {
    Vector2I coord = new((int)building.Coordinate.X, (int)building.Coordinate.Y);

    if (CityGameConsts.BuildingSprites.BuildingIs2x1(building))
    {
      coords2x1Tile.Add(new Vector2I((int)building.Coordinate.X, (int)building.Coordinate.Y));
    }

    int source_id = CityGameConsts.BuildingSprites.SOURCE_ID;
    Vector2I atlas = CityGameConsts.BuildingSprites.GetAtlas(building);
    buildingLayer.SetCell(coord, source_id, atlas);
  }

  public void UpdateHelpers()
  {
    TickData tick_data = data.Data[tick];
    GameState state = tick_data.State;

    helperLayer.Clear();

    foreach (HelperState helperState in state.HelperStates)
    {
      PlaceHelper(helperState);
    }
  }

  public void PlaceHelper(HelperState helper)
  {
    Vector2I coord = new((int)helper.Coordinate.X, (int)helper.Coordinate.Y);
    helperLayer.SetCell(coord, 2, Vector2I.Zero);
  }
}
