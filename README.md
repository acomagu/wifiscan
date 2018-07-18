# wifiscan

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=for-the-badge)](https://godoc.org/github.com/acomagu/wifiscan) [![CircleCI](https://img.shields.io/circleci/project/github/RedSparr0w/node-csgo-parser.svg?style=for-the-badge)](https://circleci.com/gh/acomagu/wifiscan)

Go library to scan Wi-Fi access points.

```Go
aps, _ := wifiscan.Scan(ctx, "wlan0")
for _, ap := range aps {
	fmt.Printf("BSSID: %s, Signal: %d\n", ap.BSSID, ap.Signal)
}
```

Supporting macOS and Linux. On Linux, this may need `sudo`.
