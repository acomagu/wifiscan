// Wi-Fi access points scanner.
//
// Supporting macOS and Linux. On each OS, depends on:
//
// - Linux: `iw`
// - macOS: `airport`
//
// A majority of these OS environments are supporting it.
//
// Be careful to the scanning needs `sudo` on Linux.
package wifiscan

import "context"

// Scan scans access points and returns it.
func Scan(ctx context.Context, ifname string) ([]*AP, error) {
	return scan(ctx, ifname)
}
