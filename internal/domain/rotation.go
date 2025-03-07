package domain

type Rotation struct {
	X      int
	Y      int
	UVLock bool
}

// Rotation lookup table for stairs.
var stairRot = map[string]map[string]map[string]Rotation{
	"east": {
		"bottom": {
			"straight":    {},
			"inner_left":  {Y: 270, UVLock: true},
			"inner_right": {},
			"outer_left":  {Y: 270, UVLock: true},
			"outer_right": {},
		},
		"top": {
			"straight":    {X: 180, UVLock: true},
			"inner_left":  {X: 180, UVLock: true},
			"inner_right": {X: 180, Y: 90, UVLock: true},
			"outer_left":  {X: 180, UVLock: true},
			"outer_right": {X: 180, Y: 90, UVLock: true},
		},
	},
	"north": {
		"bottom": {
			"straight":    {Y: 270, UVLock: true},
			"inner_left":  {Y: 180, UVLock: true},
			"inner_right": {Y: 270, UVLock: true},
			"outer_left":  {Y: 180, UVLock: true},
			"outer_right": {Y: 270, UVLock: true},
		},
		"top": {
			"straight":    {X: 180, Y: 270, UVLock: true},
			"inner_left":  {X: 180, Y: 270, UVLock: true},
			"inner_right": {X: 180, UVLock: true},
			"outer_left":  {X: 180, Y: 270, UVLock: true},
			"outer_right": {X: 180, UVLock: true},
		},
	},
	"south": {
		"bottom": {
			"straight":    {Y: 90, UVLock: true},
			"inner_left":  {},
			"inner_right": {Y: 90, UVLock: true},
			"outer_left":  {},
			"outer_right": {Y: 90, UVLock: true},
		},
		"top": {
			"straight":    {X: 180, Y: 90, UVLock: true},
			"inner_left":  {X: 180, Y: 90, UVLock: true},
			"inner_right": {X: 180, Y: 180, UVLock: true},
			"outer_left":  {X: 180, Y: 90, UVLock: true},
			"outer_right": {X: 180, Y: 180, UVLock: true},
		},
	},
	"west": {
		"bottom": {
			"straight":    {Y: 180, UVLock: true},
			"inner_left":  {Y: 90, UVLock: true},
			"inner_right": {Y: 180, UVLock: true},
			"outer_left":  {Y: 90, UVLock: true},
			"outer_right": {Y: 180, UVLock: true},
		},
		"top": {
			"straight":    {X: 180, Y: 180, UVLock: true},
			"inner_left":  {X: 180, Y: 180, UVLock: true},
			"inner_right": {X: 180, Y: 270, UVLock: true},
			"outer_left":  {X: 180, Y: 180, UVLock: true},
			"outer_right": {X: 180, Y: 270, UVLock: true},
		},
	},
}

// Rotation lookup table for walls.
var wallRot = map[string]Rotation{
	"north": {UVLock: true},
	"east":  {Y: 90, UVLock: true},
	"south": {Y: 180, UVLock: true},
	"west":  {Y: 270, UVLock: true},
}

func GetStairRotation(facing, half, shape string) Rotation {
	return stairRot[facing][half][shape]
}

func GetWallRotation(facing string) Rotation {
	return wallRot[facing]
}
