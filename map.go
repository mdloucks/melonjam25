package main

import (
	"encoding/json"
	"os"
)

// data we want for one layer in our list of layers
type TilemapLayerJSON struct {
	Data    []int    `json:"data"`
	Width   int      `json:"width"`
	Height  int      `json:"height"`
	Objects []Object `json:"objects,omitempty"`
}

type Object struct {
	Height   int          `json:"height"`
	Id       int          `json:"id"`
	Name     string       `json:"name"`
	Rotation int          `json:"rotation"`
	Type     string       `json:"type"`
	Visible  bool         `json:"visible"`
	Width    int          `json:"width"`
	X        int          `json:"x"`
	Y        int          `json:"y"`
	Polygon  []Coordinate `json:"polygon,omitempty"`
}

type Coordinate struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// all layers in a tilemap
type TilemapJSON struct {
	Layers []TilemapLayerJSON `json:"layers"`
}

// opens the file, parses it, and returns the json object + potential error
func NewTilemapJSON(filepath string) (*TilemapJSON, error) {
	contents, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var tilemapJSON TilemapJSON
	err = json.Unmarshal(contents, &tilemapJSON)
	if err != nil {
		return nil, err
	}

	return &tilemapJSON, nil
}
