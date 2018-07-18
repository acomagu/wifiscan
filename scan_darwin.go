package wifiscan

import (
	"bufio"
	"context"
	"regexp"
	"strconv"

	"github.com/pkg/errors"
)

var ssidLineRegexp = regexp.MustCompile(`^.*\s+((\w\w:){5}\w\w)\s+(-\d+).*$`)

func scan(ctx context.Context, ifname string) ([]*AP, error) {
	var aps []*AP

	cmd := commandContext(ctx, "/System/Library/PrivateFrameworks/Apple80211.framework/Versions/Current/Resources/airport", ifname, "--scan")
	opipe, err := cmd.StdoutPipe()
	if err != nil {
		return nil, errors.Wrap(err, "could not get the FD of stdout of the command to run")
	}
	defer opipe.Close()

	if err := cmd.Start(); err != nil {
		return nil, errors.Wrap(err, "could not start the command")
	}

	scn := bufio.NewScanner(opipe)
	for scn.Scan() {
		matches := ssidLineRegexp.FindStringSubmatch(scn.Text())
		if matches == nil {
			continue
		}

		bssid := matches[1]
		signalstr := matches[3]
		signal, err := strconv.Atoi(signalstr)
		if err != nil {
			continue
		}

		aps = append(aps, &AP{
			BSSID:  bssid,
			Signal: signal,
		})
	}

	return aps, cmd.Wait()
}
