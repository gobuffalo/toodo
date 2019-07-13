package events

import (
	"fmt"
)

// LoadPlugins will add listeners for any plugins that support "events"
func LoadPlugins() error {
	const msg = "events.LoadPlugins has been removed from this package. Use buffalo.LoadPlugins instead (requires Buffalo v0.13.14+ or v0.14.2+)"
	fmt.Println(msg)
	return nil
}
