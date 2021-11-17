package utils

import (
	"github.com/open-traffic-generator/snappi/gosnappi"
)

func TxRxRoutesOk(tx, rx []int) bool {
	if len(tx) != len(rx) {
		return false
	}

	totalTx := 0
	for _, t := range tx {
		// not ok if not routes sent for any of the peer
		if t == 0 {
			return false
		}
		totalTx += t
	}

	for i := range rx {
		// not ok if expected rx doesn't match sum of all tx minus self tx
		if rx[i] != totalTx-tx[i] {
			return false
		}
	}

	return true
}

func (client *GnmiClient) PortAndFlowMetricsOk(config gosnappi.Config) (bool, error) {
	expected := 0
	for _, f := range config.Flows().Items() {
		expected += int(f.Duration().FixedPackets().Packets())
	}

	fMetrics, err := client.GetFlowMetrics(nil)
	if err != nil {
		return false, err
	}
	pMetrics, err := client.GetPortMetrics(nil)
	if err != nil {
		return false, err
	}
	PrintMetricsTable(&MetricsTableOpts{
		ClearPrevious: false,
		FlowMetrics:   fMetrics,
		PortMetrics:   pMetrics,
	})

	actual := 0
	for _, m := range fMetrics.Items() {
		actual += int(m.FramesRx())
	}

	return expected == actual, nil
}

func (client *GnmiClient) AllBgp4SessionUp(config gosnappi.Config) (bool, error) {
	dNames := []string{}
	for _, d := range config.Devices().Items() {
		bgp := d.Bgp()
		for _, ip := range bgp.Ipv4Interfaces().Items() {
			for _, peer := range ip.Peers().Items() {
				dNames = append(dNames, peer.Name())
			}
		}
	}

	dMetrics, err := client.GetBgpv4Metrics(dNames)
	if err != nil {
		return false, err
	}

	PrintMetricsTable(&MetricsTableOpts{
		ClearPrevious: false,
		Bgpv4Metrics:  dMetrics,
	})

	routesTx := []int{}
	routesRx := []int{}
	actualSessionUp := 0
	for _, d := range dMetrics.Items() {
		if d.SessionState() == gosnappi.Bgpv4MetricSessionState.UP {
			actualSessionUp += 1
		}
		routesTx = append(routesTx, int(d.RoutesAdvertised()))
		routesRx = append(routesRx, int(d.RoutesReceived()))
	}

	return len(dNames) == actualSessionUp && TxRxRoutesOk(routesTx, routesRx), nil
}

func (client *GnmiClient) AllBgp6SessionUp(config gosnappi.Config) (bool, error) {
	dNames := []string{}
	for _, d := range config.Devices().Items() {
		bgp := d.Bgp()
		for _, ipv6 := range bgp.Ipv6Interfaces().Items() {
			for _, peer := range ipv6.Peers().Items() {
				dNames = append(dNames, peer.Name())
			}
		}
	}

	dMetrics, err := client.GetBgpv6Metrics(dNames)
	if err != nil {
		return false, err
	}

	PrintMetricsTable(&MetricsTableOpts{
		ClearPrevious: false,
		Bgpv6Metrics:  dMetrics,
	})

	routesTx := []int{}
	routesRx := []int{}
	actualSessionUp := 0
	for _, d := range dMetrics.Items() {
		if d.SessionState() == gosnappi.Bgpv6MetricSessionState.UP {
			actualSessionUp += 1
		}
		routesTx = append(routesTx, int(d.RoutesAdvertised()))
		routesRx = append(routesRx, int(d.RoutesReceived()))
	}

	return len(dNames) == actualSessionUp && TxRxRoutesOk(routesTx, routesRx), nil
}

type PortMetric struct {
	Name     string
	FramesTx int32
}

type FlowMetric struct {
	Name     string
	FramesRx int32
}
