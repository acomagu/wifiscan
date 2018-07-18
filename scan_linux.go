package wifiscan

import (
	"bufio"
	"context"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

const minimumThreshold = -100

func scan(ctx context.Context, ifname string) ([]*AP, error) {
	done := make(chan error)
	c := make(chan *AP)
	go func() {
		done <- iw(ctx, c, ifname)
	}()

	var aps []*AP
	for ap := range c {
		aps = append(aps, ap)
	}

	return aps, <-done
}

func iw(ctx context.Context, out chan *AP, ifname string) error {
	defer close(out)
	cmd := commandContext(ctx, "/sbin/iw", "dev", ifname, "scan", "-u")

	opp, err := cmd.StdoutPipe()
	if err != nil {
		return errors.Wrap(err, "could not get the stdout fd of the command to run")
	}
	defer opp.Close()

	if err := cmd.Start(); err != nil {
		return errors.Wrap(err, "could not start the command")
	}

	var name string
	var signal int
	scn := bufio.NewScanner(opp)
	for scn.Scan() {
		line := scn.Text()
		if strings.Contains(line, "(on") {
			name = strings.Split(strings.Split(line, "(")[0], "BSS")[1]
			name = strings.ToLower(name)
			name = strings.TrimSpace(name)
		} else if strings.Contains(line, "signal:") {
			foo := strings.Split(line, "signal:")[1]
			foo = strings.Split(foo, ".")[0]
			foo = strings.TrimSpace(foo)
			var err error
			signal, err = strconv.Atoi(foo)
			if err != nil {
				panic(err)
			}
		}
		if name != "" && signal != 0 {
			if signal < minimumThreshold {
				continue
			}

			out <- &AP{
				BSSID:  name,
				Signal: signal,
			}

			name = ""
			signal = 0
		}
	}

	return cmd.Wait()
}
