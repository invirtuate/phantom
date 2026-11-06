package ipban

import (
	"fmt"
	"slices"
)

// Authorized returns nil if authorized
func Authorized(ip string) error {
	if slices.Contains(BannedIPs, ip) {
		return fmt.Errorf("%s: %w", ip, ErrIPBanned)
	}

	return nil
}
