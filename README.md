# Configuration Agent Template

This project is meant to be a template where you can start developing a new configuration Agent

## Functionalities

This template provides interfaces to receive different configuration requests such as

* __AssignIPToInterface__
* __ConfigureNat__
* __AddRoute__

### AssignIPToInterface

It accepts an object of type `IPAssignment`, which looks like

```go
type IPAssignment struct {
  Ip        string
  Interface string
}
```

### ConfigureNat

It accepts an object of type `NatConfiguration`, which looks like

```go
type NatConfiguration struct {
  LocalNetwork   string
  RemoteNetwork  string
  LocalInterface string
  TableNumber    string
  PacketMark     string
  NextHop        string
}
```

### AddRoute

It accepts an object of type `Route` which looks like

```go
type Route struct {
  DestinationNetwork string
  NextHop            string
}
```

## How to use it

You can clone this repository and insert the implementation required to configure your network devices.

## Contributions

If you want to make public your configurator (if you think that it could be useful for other people) contact us at `cloud@top-ix.org`, and we will create a repository in the organization.
