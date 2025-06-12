// Package raindrops converts nunmbers into rain sounds
package raindrops

import "strconv"

// Sounds to make
var Sounds = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

// better choice would be ordered map
var keys = [3]int{3, 5, 7}

// Convert int into sound
func Convert(num int) string {
	var sound string
	for _, key := range keys {
		if num%key == 0 {
			sound += Sounds[key]
		}
	}
	//if no multipliers found
	if sound == "" {
		sound = strconv.Itoa(num)
	}
	return sound
}
