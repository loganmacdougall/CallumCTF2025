package main

import (
	"city_game/src/game/utility"
	"city_game/src/pb"
	"testing"
)

func TestPathFinding(t *testing.T) {
	grid := [][]int{
		{-1, -1, -1, -1, 0, 0, -1},
		{-1, -1, -1, 0, 0, -1, -1},
		{0, 0, -1, -1, 0, -1, -1},
		{0, 0, -1, 0, -1, -1, -1},
		{0, 0, -1, -1, -1, -1, 0},
	}
	start := &pb.Coordinate{X: 0, Y: 0}
	end := &pb.Coordinate{X: 6, Y: 0}

	steps := 15
	path := utility.FindShortestPath(grid, start, end)

	if len(path) != steps {
		t.Logf("number of steps should be %d, found %d\n", steps, len(path))
		t.FailNow()
	}
}

func TestPathFindingNoPath(t *testing.T) {
	grid := [][]int{
		{-1, -1, -1, -1, 0, 0, -1},
		{-1, -1, -1, 0, 0, -1, -1},
		{0, 0, -1, -1, 0, -1, -1},
		{0, 0, -1, 0, -1, -1, -1},
		{0, 0, -1, -1, 0, -1, 0},
	}
	start := &pb.Coordinate{X: 0, Y: 0}
	end := &pb.Coordinate{X: 6, Y: 0}

	path := utility.FindShortestPath(grid, start, end)

	if path != nil {
		t.Logf("path should be nil, found path in %d steps\n", len(path))
		t.FailNow()
	}
}
