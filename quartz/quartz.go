package quartz

// #cgo LDFLAGS: -framework Foundation -framework CoreGraphics
// #include <CoreGraphics/CoreGraphics.h>
import "C"

// OnlineDisplays retrieves all online displays.
func OnlineDisplays() []Display {
	var rawDisplays [16]C.uint
	var count C.uint

	C.CGGetOnlineDisplayList(16, &rawDisplays[0], &count)

	displays := make([]Display, count)
	for i := 0; i < int(count); i++ {
		displays[i] = Display{id: rawDisplays[i]}
	}

	return displays
}
