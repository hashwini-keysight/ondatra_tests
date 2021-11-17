# ondatra_tests
This repository contains ondatra test suite using gosnappi

```
To run the test go to tests/ directory and run like below:

go test -run=TestGoSnappiK8s_001 -testbed=../testbed/testbed-001.txt -config=../kne/kne-001.yaml
```

```
Sample test output:

/ondatra_tests/tests$ go test -run=TestGoSnappiK8s_001 -testbed=../testbed/testbed-001.txt -config=../kne/kne-001.yaml
*** Reserving the testbed
2021/11/17 12:04:10 Setting port name and location ...
2021/11/17 12:04:10 Setting port Name(Topology): ixia-c-port1, Location: service-ixia-c-port1.ixia-c.svc.cluster.local:5555+service-ixia-c-port1.ixia-c.svc.cluster.local:50071, Name(Config) ixia-c-port1
2021/11/17 12:04:10 Setting port Name(Topology): ixia-c-port2, Location: service-ixia-c-port2.ixia-c.svc.cluster.local:5555+service-ixia-c-port2.ixia-c.svc.cluster.local:50071, Name(Config) ixia-c-port2
2021/11/17 12:04:10 Setting port Name(Topology): ixia-c-port3, Location: service-ixia-c-port3.ixia-c.svc.cluster.local:5555+service-ixia-c-port3.ixia-c.svc.cluster.local:50071, Name(Config) ixia-c-port3
2021/11/17 12:04:10 Pushing config ...
2021/11/17 12:04:10 Start protocols ...
2021/11/17 12:04:10 New GNMI Query @172.18.0.107:50051
2021/11/17 12:04:10 Creating gNMI client for server ...
2021/11/17 12:04:10 Successfully created gNMI client !
2021/11/17 12:04:10 Waiting for condition to be true ...
2021/11/17 12:04:10 GNMI Query: [[bgpv6_metrics[name=BGPv6 Peer 1]] [bgpv6_metrics[name=BGPv6 Peer 2]] [bgpv6_metrics[name=BGPv6 Peer 3]]]
2021/11/17 12:04:10 Getting bgpv6 metrics ...
2021/11/17 12:04:12 GetBgpv6Metrics GNMI took 2191 ms
2021/11/17 12:04:12 

Bgpv6 Metrics
-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
Name                Session State       Session Flaps       Routes Advertised   Routes Received     Route Withdraws Tx  Route Withdraws Rx  Keepalives Tx       Keepalives Rx       
BGPv6 Peer 1        down                0                   0                   0                   0                   0                   0                   0                   
BGPv6 Peer 2        down                0                   0                   0                   0                   0                   0                   0                   
BGPv6 Peer 3        up                  0                   2                   0                   0                   0                   2                   2                   
-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------


2021/11/17 12:04:12 GNMI Query: [[bgpv6_metrics[name=BGPv6 Peer 1]] [bgpv6_metrics[name=BGPv6 Peer 2]] [bgpv6_metrics[name=BGPv6 Peer 3]]]
2021/11/17 12:04:12 Getting bgpv6 metrics ...
2021/11/17 12:04:14 GetBgpv6Metrics GNMI took 1521 ms
2021/11/17 12:04:14 

Bgpv6 Metrics
-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
Name                Session State       Session Flaps       Routes Advertised   Routes Received     Route Withdraws Tx  Route Withdraws Rx  Keepalives Tx       Keepalives Rx       
BGPv6 Peer 1        up                  0                   2                   4                   0                   0                   2                   2                   
BGPv6 Peer 2        up                  0                   2                   4                   0                   0                   2                   2                   
BGPv6 Peer 3        up                  0                   2                   4                   0                   0                   2                   2                   
-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------


2021/11/17 12:04:14 Done waiting for condition to be true
2021/11/17 12:04:14 Waiting for condition to be true took 4213 ms
2021/11/17 12:04:14 Starting transmit ...
2021/11/17 12:04:14 Waiting for condition to be true ...
2021/11/17 12:04:14 GNMI Query: [[flow_metrics[name=ixia-c-port1 -> ixia-c-port2]] [flow_metrics[name=ixia-c-port1 -> ixia-c-port3]] [flow_metrics[name=ixia-c-port2 -> ixia-c-port1]] [flow_metrics[name=ixia-c-port3 -> ixia-c-port1]]]
2021/11/17 12:04:14 Getting flow metrics ...
2021/11/17 12:04:16 GetFlowMetrics GNMI took 1981 ms
2021/11/17 12:04:16 GNMI Query: [[port_metrics[name=ixia-c-port1]] [port_metrics[name=ixia-c-port2]] [port_metrics[name=ixia-c-port3]]]
2021/11/17 12:04:16 Getting port metrics ...
2021/11/17 12:04:18 GetPortMetrics GNMI took 1822 ms
2021/11/17 12:04:18 

Port Metrics
-----------------------------------------------------------------
Name           Frames Tx      Frames Rx      FPS Tx         
ixia-c-port1   2000           2003           0              
ixia-c-port2   1000           1003           0              
ixia-c-port3   1000           1003           0              
-----------------------------------------------------------------


Flow Metrics
--------------------------------------------------
Name           Frames Rx      FPS Rx         
ixia-c-port1 -> ixia-c-port2977            0              
ixia-c-port1 -> ixia-c-port3947            0              
ixia-c-port2 -> ixia-c-port1975            0              
ixia-c-port3 -> ixia-c-port1979            0              
--------------------------------------------------


2021/11/17 12:04:18 GNMI Query: [[flow_metrics[name=ixia-c-port1 -> ixia-c-port2]] [flow_metrics[name=ixia-c-port1 -> ixia-c-port3]] [flow_metrics[name=ixia-c-port2 -> ixia-c-port1]] [flow_metrics[name=ixia-c-port3 -> ixia-c-port1]]]
2021/11/17 12:04:18 Getting flow metrics ...
2021/11/17 12:04:20 GetFlowMetrics GNMI took 1684 ms
2021/11/17 12:04:20 GNMI Query: [[port_metrics[name=ixia-c-port1]] [port_metrics[name=ixia-c-port2]] [port_metrics[name=ixia-c-port3]]]
2021/11/17 12:04:20 Getting port metrics ...
2021/11/17 12:04:22 GetPortMetrics GNMI took 1832 ms
2021/11/17 12:04:22 

Port Metrics
-----------------------------------------------------------------
Name           Frames Tx      Frames Rx      FPS Tx         
ixia-c-port1   2000           2003           0              
ixia-c-port2   1000           1003           0              
ixia-c-port3   1000           1003           0              
-----------------------------------------------------------------


Flow Metrics
--------------------------------------------------
Name           Frames Rx      FPS Rx         
ixia-c-port1 -> ixia-c-port21000           383            
ixia-c-port1 -> ixia-c-port31000           441            
ixia-c-port2 -> ixia-c-port11000           416            
ixia-c-port3 -> ixia-c-port11000           437            
--------------------------------------------------


2021/11/17 12:04:22 Done waiting for condition to be true
2021/11/17 12:04:22 Waiting for condition to be true took 7821 ms
2021/11/17 12:04:22 Stop protocols ...
PASS
*** Releasing the testbed
ok      tests/tests     13.332s

```
