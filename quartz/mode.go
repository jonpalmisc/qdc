package quartz

// #cgo LDFLAGS: -framework Foundation -framework CoreGraphics
// #include <CoreGraphics/CoreGraphics.h>
import "C"
import "fmt"

// DisplayMode represents a CoreGraphics display mode.
type DisplayMode struct {
	ref C.CGDisplayModeRef
}

// Width gets the display mode's width.
func (dm *DisplayMode) Width() uint {
	return uint(C.CGDisplayModeGetWidth(dm.ref))
}

// Height gets the display mode's height.
func (dm *DisplayMode) Height() uint {
	return uint(C.CGDisplayModeGetHeight(dm.ref))
}

// Resolution returns the display mode's size formatted as a string.
func (dm *DisplayMode) Resolution() string {
	return fmt.Sprintf("%dx%d", dm.Width(), dm.Height())
}

// RefreshRate gets the display mode's refresh rate.
func (dm *DisplayMode) RefreshRate() float64 {
	return float64(C.CGDisplayModeGetRefreshRate(dm.ref))
}

// Magnitude returns the product of the display mode's width and height.
func (dm *DisplayMode) Magnitude() uint {
	return dm.Width() * dm.Height()
}
