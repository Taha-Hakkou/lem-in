package main

import "math"

// ---------- DATA TYPES ----------
type Room struct {
	x, y int
}

// number of ants
// var n int

// Rooms: id -> (x,y)
var rooms = map[string]Room{}

// Links
var links = [][2]string{}

// Steps
var steps = [][][2]string{}

// ---------- NORMALIZE ----------
var (
	width, height int
	scale         = 4
)

var (
	minX, maxX = math.MaxInt, math.MinInt
	minY, maxY = math.MaxInt, math.MinInt
)

var action bool
