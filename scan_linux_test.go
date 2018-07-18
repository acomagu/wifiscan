package wifiscan

import (
	"context"
	"io"
	"io/ioutil"
	"strings"
	"testing"
)

type cmdStub struct {
	stdout string
}

func (c *cmdStub) StdoutPipe() (io.ReadCloser, error) {
	return ioutil.NopCloser(strings.NewReader(c.stdout)), nil
}

func (c *cmdStub) Wait() error {
	return nil
}

func TestListAP(t *testing.T) {
	airportout := `
BSS 00:23:b1:8c:02:a1(on wlp2s0) -- associated
	TSF: 11261290096 usec (0d, 03:07:41)
	freq: 2422
	beacon interval: 100 TUs
	capability: ESS Privacy ShortPreamble ShortSlotTime (0x0431)
	signal: -34.00 dBm
	last seen: 164 ms ago
	Information elements from Probe Response frame:
	SSID: acomagu1
	Supported rates: 1.0* 2.0* 5.5* 11.0* 
	DS Parameter set: channel 3
	Country: JP	Environment: bogus
		Channels [1 - 1] @ 20 dBm
		Channels [2 - 2] @ 20 dBm
		Channels [3 - 3] @ 20 dBm
		Channels [4 - 4] @ 20 dBm
		Channels [5 - 5] @ 20 dBm
		Channels [6 - 6] @ 20 dBm
		Channels [7 - 7] @ 20 dBm
		Channels [8 - 8] @ 20 dBm
		Channels [9 - 9] @ 20 dBm
		Channels [10 - 10] @ 20 dBm
		Channels [11 - 11] @ 20 dBm
		Channels [12 - 12] @ 20 dBm
		Channels [13 - 13] @ 20 dBm
		Channels [14 - 14] @ 20 dBm
	ERP: Barker_Preamble_Mode
	Extended supported rates: 6.0 9.0 12.0 18.0 24.0 36.0 48.0 54.0 
	RSN:	 * Version: 1
		 * Group cipher: CCMP
		 * Pairwise ciphers: CCMP
		 * Authentication suites: PSK
		 * Capabilities: 16-PTKSA-RC 1-GTKSA-RC MFP-capable (0x008c)
		 * 0 PMKIDs
		 * Group mgmt cipher suite: AES-128-CMAC
	HT capabilities:
		Capabilities: 0x12d
			RX LDPC
			HT20
			SM Power Save disabled
			RX HT20 SGI
			RX STBC 1-stream
			Max AMSDU length: 3839 bytes
			No DSSS/CCK HT40
		Maximum RX AMPDU length 65535 bytes (exponent: 0x003)
		Minimum RX AMPDU time spacing: 16 usec (0x07)
		HT RX MCS rate indexes supported: 0-7
		HT TX MCS rate indexes are undefined
	HT operation:
		 * primary channel: 3
		 * secondary channel offset: no secondary
		 * STA channel width: 20 MHz
		 * RIFS: 0
		 * HT protection: no
		 * non-GF present: 1
		 * OBSS non-GF present: 0
		 * dual beacon: 0
		 * dual CTS protection: 0
		 * STBC beacon: 0
		 * L-SIG TXOP Prot: 0
		 * PCO active: 0
		 * PCO phase: 0
	WMM:	 * Parameter version 1
		 * u-APSD
		 * BE: CW 15-1023, AIFSN 3
		 * BK: CW 15-1023, AIFSN 7
		 * VI: CW 7-15, AIFSN 2, TXOP 3008 usec
		 * VO: CW 3-7, AIFSN 2, TXOP 1504 usec
	Extended capabilities:
		 * Extended Channel Switching
	WPS:	 * Version: 1.0
		 * Wi-Fi Protected Setup State: 2 (Configured)
		 * Response Type: 3 (AP)
		 * UUID: 73c0a46d-68ed-5771-87fe-4c5ed9b712b8
		 * Manufacturer: Fujisoft.inc
		 * Model: FS030W
		 * Model Number: FS030W
		 * Serial Number: 123456
		 * Primary Device Type: 6-0050f204-1
		 * Device name: FS030W
		 * Config methods: PBC
		 * Unknown TLV (0x1049, 6 bytes): 00 37 2a 00 01 20
BSS f0:9c:e9:44:34:56(on wlp2s0)
	TSF: 36310136631315 usec (420d, 06:08:56)
	freq: 2412
	beacon interval: 100 TUs
	capability: ESS Privacy ShortPreamble ShortSlotTime (0x0431)
	signal: -66.00 dBm
	last seen: 22096 ms ago
	Information elements from Probe Response frame:
	SSID: 
	Supported rates: 1.0* 2.0* 5.5* 11.0* 6.0 9.0 12.0 18.0 
	DS Parameter set: channel 1
	TIM: DTIM Count 0 DTIM Period 1 Bitmap Control 0x0 Bitmap[0] 0x0
	Country: JP	Environment: Indoor/Outdoor
		Channels [1 - 13] @ 20 dBm
	ERP: <no flags>
	RSN:	 * Version: 1
		 * Group cipher: CCMP
		 * Pairwise ciphers: CCMP
		 * Authentication suites: IEEE 802.1X
		 * Capabilities: 1-PTKSA-RC 1-GTKSA-RC (0x0000)
		 * 0 PMKIDs
	Extended supported rates: 24.0 36.0 48.0 54.0 
	HT capabilities:
		Capabilities: 0x18d
			RX LDPC
			HT20
			SM Power Save disabled
			TX STBC
			RX STBC 1-stream
			Max AMSDU length: 3839 bytes
			No DSSS/CCK HT40
		Maximum RX AMPDU length 65535 bytes (exponent: 0x003)
		Minimum RX AMPDU time spacing: 8 usec (0x06)
		HT RX MCS rate indexes supported: 0-15
		HT TX MCS rate indexes are undefined
	HT operation:
		 * primary channel: 1
		 * secondary channel offset: no secondary
		 * STA channel width: 20 MHz
		 * RIFS: 0
		 * HT protection: nonmember
		 * non-GF present: 0
		 * OBSS non-GF present: 0
		 * dual beacon: 0
		 * dual CTS protection: 0
		 * STBC beacon: 0
		 * L-SIG TXOP Prot: 0
		 * PCO active: 0
		 * PCO phase: 0
	Overlapping BSS scan params:
		 * passive dwell: 20 TUs
		 * active dwell: 10 TUs
		 * channel width trigger scan interval: 300 s
		 * scan passive total per channel: 200 TUs
		 * scan active total per channel: 20 TUs
		 * BSS width channel transition delay factor: 5
		 * OBSS Scan Activity Threshold: 0.25 %
	Extended capabilities:
		 * HT Information Exchange Supported
BSS 00:01:8e:a4:58:93(on wlp2s0)
	TSF: 6353689497974 usec (73d, 12:54:49)
	freq: 2412
	beacon interval: 100 TUs
	capability: ESS Privacy ShortSlotTime (0x0411)
	signal: -88.00 dBm
	last seen: 15440 ms ago
	Information elements from Probe Response frame:
	SSID: logitec28
	Supported rates: 1.0* 2.0* 5.5* 11.0* 6.0 9.0 12.0 18.0 
	DS Parameter set: channel 1
	TIM: DTIM Count 0 DTIM Period 3 Bitmap Control 0x0 Bitmap[0] 0x0
	ERP: Barker_Preamble_Mode
	Extended supported rates: 24.0 36.0 48.0 54.0 
	HT capabilities:
		Capabilities: 0x186e
			HT20/HT40
			SM Power Save disabled
			RX HT20 SGI
			RX HT40 SGI
			No RX STBC
			Max AMSDU length: 7935 bytes
			DSSS/CCK HT40
		Maximum RX AMPDU length 32767 bytes (exponent: 0x002)
		Minimum RX AMPDU time spacing: 16 usec (0x07)
		HT RX MCS rate indexes supported: 0-15
		HT TX MCS rate indexes are undefined
	HT operation:
		 * primary channel: 1
		 * secondary channel offset: above
		 * STA channel width: any
		 * RIFS: 0
		 * HT protection: no
		 * non-GF present: 0
		 * OBSS non-GF present: 0
		 * dual beacon: 0
		 * dual CTS protection: 0
		 * STBC beacon: 0
		 * L-SIG TXOP Prot: 0
		 * PCO active: 0
		 * PCO phase: 0
	RSN:	 * Version: 1
		 * Group cipher: CCMP
		 * Pairwise ciphers: CCMP
		 * Authentication suites: PSK
		 * Capabilities: 1-PTKSA-RC 1-GTKSA-RC (0x0000)
	WMM:	 * Parameter version 1
		 * BE: CW 15-1023, AIFSN 3
		 * BK: CW 15-1023, AIFSN 7
		 * VI: CW 7-15, AIFSN 2, TXOP 3008 usec
		 * VO: CW 3-7, AIFSN 2, TXOP 1504 usec
	WPS:	 * Version: 1.0
		 * Wi-Fi Protected Setup State: 2 (Configured)
BSS 10:66:82:f8:62:03(on wlp2s0)
	TSF: 680149360514 usec (7d, 20:55:49)
	freq: 2412
	beacon interval: 100 TUs
	capability: ESS Privacy ShortPreamble ShortSlotTime (0x0431)
	signal: -78.00 dBm
	last seen: 2356 ms ago
	SSID: pr500m-097499-1
	Supported rates: 1.0* 2.0* 5.5* 11.0* 6.0 9.0 12.0 18.0 
	DS Parameter set: channel 1
	Country: JP	Environment: Indoor/Outdoor
		Channels [1 - 13] @ 20 dBm
	ERP: <no flags>
	Extended supported rates: 24.0 36.0 48.0 54.0 
	HT capabilities:
		Capabilities: 0x11ef
			RX LDPC
			HT20/HT40
			SM Power Save disabled
			RX HT20 SGI
			RX HT40 SGI
			TX STBC
			RX STBC 1-stream
			Max AMSDU length: 3839 bytes
			DSSS/CCK HT40
		Maximum RX AMPDU length 65535 bytes (exponent: 0x003)
		Minimum RX AMPDU time spacing: 8 usec (0x06)
		HT RX MCS rate indexes supported: 0-15
		HT TX MCS rate indexes are undefined
	HT operation:
		 * primary channel: 1
		 * secondary channel offset: no secondary
		 * STA channel width: 20 MHz
		 * RIFS: 1
		 * HT protection: no
		 * non-GF present: 0
		 * OBSS non-GF present: 0
		 * dual beacon: 0
		 * dual CTS protection: 0
		 * STBC beacon: 0
		 * L-SIG TXOP Prot: 0
		 * PCO active: 0
		 * PCO phase: 0
	Overlapping BSS scan params:
		 * passive dwell: 20 TUs
		 * active dwell: 10 TUs
		 * channel width trigger scan interval: 300 s
		 * scan passive total per channel: 200 TUs
		 * scan active total per channel: 20 TUs
		 * BSS width channel transition delay factor: 5
		 * OBSS Scan Activity Threshold: 0.25 %
	Extended capabilities:
		 * HT Information Exchange Supported
		 * Operating Mode Notification
	WMM:	 * Parameter version 1
		 * u-APSD
		 * BE: CW 15-1023, AIFSN 3
		 * BK: CW 15-1023, AIFSN 7
		 * VI: CW 7-15, AIFSN 2, TXOP 3008 usec
		 * VO: CW 3-7, AIFSN 2, TXOP 1504 usec
	RSN:	 * Version: 1
		 * Group cipher: CCMP
		 * Pairwise ciphers: CCMP
		 * Authentication suites: PSK
		 * Capabilities: 1-PTKSA-RC 1-GTKSA-RC (0x0000)
BSS 88:57:ee:2a:c7:ad(on wlp2s0)
	TSF: 375299473793 usec (4d, 08:14:59)
	freq: 2412
	beacon interval: 100 TUs
	capability: ESS Privacy ShortSlotTime (0x0411)
	signal: -71.00 dBm
	last seen: 2276 ms ago
	Information elements from Probe Response frame:
	SSID: Buffalo-G-C7AC
	Supported rates: 1.0* 2.0* 5.5* 11.0* 9.0 18.0 36.0 54.0 
	DS Parameter set: channel 1
	ERP: Barker_Preamble_Mode
	Extended supported rates: 6.0 12.0 24.0 48.0 
	HT capabilities:
		Capabilities: 0x1ac
			HT20
			SM Power Save disabled
			RX HT20 SGI
			TX STBC
			RX STBC 1-stream
			Max AMSDU length: 3839 bytes
			No DSSS/CCK HT40
		Maximum RX AMPDU length 65535 bytes (exponent: 0x003)
		Minimum RX AMPDU time spacing: 4 usec (0x05)
		HT RX MCS rate indexes supported: 0-15
		HT TX MCS rate indexes are undefined
	HT operation:
		 * primary channel: 1
		 * secondary channel offset: no secondary
		 * STA channel width: 20 MHz
		 * RIFS: 0
		 * HT protection: nonmember
		 * non-GF present: 1
		 * OBSS non-GF present: 0
		 * dual beacon: 0
		 * dual CTS protection: 0
		 * STBC beacon: 0
		 * L-SIG TXOP Prot: 0
		 * PCO active: 0
		 * PCO phase: 0
	WPA:	 * Version: 1
		 * Group cipher: TKIP
		 * Pairwise ciphers: TKIP CCMP
		 * Authentication suites: PSK
	RSN:	 * Version: 1
		 * Group cipher: TKIP
		 * Pairwise ciphers: TKIP CCMP
		 * Authentication suites: PSK
		 * Capabilities: 1-PTKSA-RC 1-GTKSA-RC (0x0000)
	Extended capabilities:
	BSS Load:
		 * station count: 1
		 * channel utilisation: 51/255
		 * available admission capacity: 31250 [*32us]
	WMM:	 * Parameter version 1
		 * BE: CW 15-1023, AIFSN 3
		 * BK: CW 15-1023, AIFSN 7
		 * VI: CW 7-15, AIFSN 2, TXOP 3008 usec
		 * VO: CW 3-7, AIFSN 2, TXOP 1504 usec
	Country: JP	Environment: Indoor/Outdoor
		Channels [1 - 13] @ 16 dBm
	WPS:	 * Version: 1.0
		 * Wi-Fi Protected Setup State: 2 (Configured)
		 * Response Type: 3 (AP)
		 * UUID: 28802880-2880-1880-a880-8857ee2ac7ad
		 * Manufacturer: Buffalo Inc.
		 * Model: WSR-1166DHP2
		 * Model Number: RT2860
		 * Serial Number: 12345678
		 * Primary Device Type: 6-0050f204-1
		 * Device name: WPS
		 * Config methods: Label, Display, PBC
		 * RF Bands: 0x1
		 * Unknown TLV (0x1049, 6 bytes): 00 37 2a 00 01 20
`

	commandContext = func(ctx context.Context, name string, args ...string) cmd {
		return &cmdStub{
			stdout: airportout,
		}
	}

	ctx := context.Background()
	aps, err := Scan(ctx, "abc")
	if err != nil {
		t.Fatalf("error is not expected but %v", err)
	}

	expected := []*AP{
		{
			BSSID:  "00:23:b1:8c:02:a1",
			Signal: -34,
		},
		{
			BSSID:  "f0:9c:e9:44:34:56",
			Signal: -66,
		},
		{
			BSSID:  "00:01:8e:a4:58:93",
			Signal: -88,
		},
		{
			BSSID:  "10:66:82:f8:62:03",
			Signal: -78,
		},
		{
			BSSID:  "88:57:ee:2a:c7:ad",
			Signal: -71,
		},
	}

	if len(aps) != len(expected) {
		t.Errorf("expected len(aps) is %v but %v", len(expected), len(aps))
	}

	for _, ap := range expected {
		if !containsAP(aps, ap) {
			t.Errorf("expected the %#+v is contained in the return value but don't", *ap)
		}
	}
}

func containsAP(aps []*AP, ap *AP) bool {
	for _, pap := range aps {
		if ap.BSSID == pap.BSSID && ap.Signal == pap.Signal {
			return true
		}
	}
	return false
}
