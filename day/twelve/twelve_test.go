package twelve

import (
	"fmt"
	"reflect"
	"testing"
)

func TestExecuteInstruction(t *testing.T) {
	tests := []struct {
		instruction string
		p           position
		want        position
	}{
		{"N1", position{x: 0, y: 0, facing: 'E'}, position{x: 0, y: 1, facing: 'E'}},
		{"N3", position{x: 0, y: 0, facing: 'E'}, position{x: 0, y: 3, facing: 'E'}},

		{"F3", position{x: 0, y: 0, facing: 'N'}, position{x: 0, y: 3, facing: 'N'}},
		{"F5", position{x: 0, y: 0, facing: 'S'}, position{x: 0, y: -5, facing: 'S'}},
		{"F4", position{x: 3, y: 6, facing: 'W'}, position{x: -1, y: 6, facing: 'W'}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s to %v°", tt.instruction, tt.p), func(t *testing.T) {
			got := executeInstruction(tt.instruction, tt.p)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("expected result of \"%s\", from %v to be %v, but was %v", tt.instruction, tt.p, tt.want, got)
			}
		})
	}
}
func TestTurnRight(t *testing.T) {
	tests := []struct {
		facing  rune
		degrees int
		want    rune
	}{
		{'N', 90, 'E'},
		{'N', 180, 'S'},
		{'N', 270, 'W'},
		{'N', 360, 'N'},

		{'E', 90, 'S'},
		{'E', 180, 'W'},
		{'E', 270, 'N'},
		{'E', 360, 'E'},

		{'S', 450, 'W'},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%c -> %d°", tt.facing, tt.degrees), func(t *testing.T) {
			got := turnRight(tt.facing, tt.degrees)
			if got != tt.want {
				t.Errorf("expected turn right of %d° from %c to be %c, not %c", tt.degrees, tt.facing, tt.want, got)
			}
		})
	}
}

func TestRotateRight(t *testing.T) {
	tests := []struct {
		wp      waypoint
		degrees int
		want    waypoint
	}{
		{waypoint{deltaX: 3, deltaY: 1}, 90, waypoint{deltaX: 1, deltaY: -3}},
		{waypoint{deltaX: 1, deltaY: -3}, 90, waypoint{deltaX: -3, deltaY: -1}},
		{waypoint{deltaX: -3, deltaY: -1}, 90, waypoint{deltaX: -1, deltaY: 3}},
		{waypoint{deltaX: -1, deltaY: 3}, 90, waypoint{deltaX: 3, deltaY: 1}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v rotate by %d°", tt.wp, tt.degrees), func(t *testing.T) {
			got := rotateRight(tt.wp, tt.degrees)
			if got != tt.want {
				t.Errorf("expected rotate right of %d° from %v to be %v, not %v", tt.degrees, tt.wp, tt.want, got)
			}
		})
	}
}

func TestTurnLeft(t *testing.T) {
	tests := []struct {
		facing  rune
		degrees int
		want    rune
	}{
		{'N', 90, 'W'},
		{'N', 180, 'S'},
		{'N', 270, 'E'},
		{'N', 360, 'N'},

		{'E', 90, 'N'},
		{'E', 180, 'W'},
		{'E', 270, 'S'},
		{'E', 360, 'E'},

		{'S', 450, 'E'},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%c -> %d°", tt.facing, tt.degrees), func(t *testing.T) {
			got := turnLeft(tt.facing, tt.degrees)
			if got != tt.want {
				t.Errorf("expected turn right of %d° from %c to be %c, not %c", tt.degrees, tt.facing, tt.want, got)
			}
		})
	}
}
