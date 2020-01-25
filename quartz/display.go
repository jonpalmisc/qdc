package quartz

// #cgo LDFLAGS: -framework Foundation -framework CoreGraphics
// #include <CoreGraphics/CoreGraphics.h>
import "C"
import (
	"errors"
	"fmt"
	"sort"
)

// Display represents a CoreGraphics display.
type Display struct {
	id C.uint
}

// Resolution gets the current resolution of a display.
func (d *Display) Resolution() string {
	width := C.CGDisplayPixelsWide(d.id)
	height := C.CGDisplayPixelsHigh(d.id)

	return fmt.Sprintf("%dx%d", uint(width), uint(height))
}

// ID returns the display's CoreGraphics ID.
func (d *Display) ID() uint {
	return uint(d.id)
}

// Modes gets all of the display modes for a display.
func (d *Display) Modes() []DisplayMode {
	rawModes := C.CGDisplayCopyAllDisplayModes(d.id, 0)
	count := int(C.CFArrayGetCount(rawModes))

	modes := make([]DisplayMode, count)
	for i := 0; i < count; i++ {
		ref := C.CGDisplayModeRef(C.CFArrayGetValueAtIndex(rawModes, C.long(i)))
		modes[i] = DisplayMode{ref: ref}
	}

	C.CFRelease(C.CFTypeRef(rawModes))

	sort.Slice(modes, func(i, j int) bool {
		if modes[i].Magnitude() > modes[j].Magnitude() {
			return true
		} else if modes[i].Magnitude() == modes[j].Magnitude() {
			return modes[i].RefreshRate() > modes[j].RefreshRate()
		}

		return false
	})

	return modes
}

// FindMode finds a display mode with the specified parameters.
func (d *Display) FindMode(res string) (*DisplayMode, error) {
	modes := d.Modes()

	for _, m := range modes {
		if m.Resolution() == res {
			return &m, nil
		}
	}

	return &DisplayMode{}, errors.New("couldn't find matching display mode")
}

// ApplyMode configures the display to use a display mode.
func (d *Display) ApplyMode(m *DisplayMode) {
	var config C.CGDisplayConfigRef

	C.CGBeginDisplayConfiguration(&config)
	C.CGConfigureDisplayWithDisplayMode(config, d.id, m.ref, 0)
	C.CGCompleteDisplayConfiguration(config, 1)
}

// MirrorTo configures mirroring from the display to another.
func (d *Display) MirrorTo(td *Display) {
	var config C.CGDisplayConfigRef

	C.CGBeginDisplayConfiguration(&config)
	C.CGConfigureDisplayMirrorOfDisplay(config, td.id, d.id)
	C.CGCompleteDisplayConfiguration(config, 1)
}
