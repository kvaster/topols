[![GitHub release](https://img.shields.io/github/v/release/kvaster/topols.svg?maxAge=60)][releases]
[![Main](https://github.com/kvaster/topols/workflows/Main/badge.svg)](https://github.com/kvaster/topols/actions)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/kvaster/topols?tab=overview)](https://pkg.go.dev/github.com/kvaster/topols?tab=overview)
[![Go Report Card](https://goreportcard.com/badge/github.com/kvaster/topols)](https://goreportcard.com/badge/github.com/kvaster/topols)

TopoLS
=======

TopoLS is for of [TopoLVM](https://github.com/topolvm/topolvm) project.
TopoLS goal is the same as TopoLVM - provisioning of local storage, but TopoLS uses btrfs qoutas instead of lvm2.
TopoLS runs in separate namespace and can be installed alongside with TopoLVM.

TopoLS code is ready for testing, but documentation is not yet finished (even touched), so please wait for corrected docs.

TopoLVM
=======

TopoLVM is a [CSI][] plugin using LVM for Kubernetes.
It can be considered as a specific implementation of [local persistent volumes](https://kubernetes.io/docs/concepts/storage/volumes/#local) using CSI and LVM.

The team presented the motivation and implementation of TopoLVM at KubeCon Europe 2020: https://kccnceu20.sched.com/event/ZerD

Join our community on Slack: [Invitation form](https://docs.google.com/forms/d/e/1FAIpQLSd2zZhqZUDTs8YUfhvKmSI_xb_iiPnz3-Hy6S7ehmHHmiifEg/viewform?embedded=true)

- **Project Status**: Testing for production
- **Conformed CSI version**: [1.1.0](https://github.com/container-storage-interface/spec/blob/v1.1.0/spec.md)

Supported environments
----------------------

- Kubernetes: 1.18, 1.17, 1.16
- Node OS: Linux with LVM2
- Filesystems: ext4, xfs, btrfs

Features
--------

- [Dynamic provisioning](https://kubernetes-csi.github.io/docs/external-provisioner.html): Volumes are created dynamically when `PersistentVolumeClaim` objects are created.
- [Raw block volume](https://kubernetes-csi.github.io/docs/raw-block.html): Volumes are available as block devices inside containers.
- [Ephemeral inline volume](https://kubernetes.io/docs/concepts/storage/volumes/#csi-ephemeral-volumes): Volumes can be directly embedded in the Pod specification.
- [Topology](https://kubernetes-csi.github.io/docs/topology.html): TopoLVM uses CSI topology feature to schedule Pod to Node where LVM volume exists.
- Extended scheduler: TopoLVM extends the general Pod scheduler to prioritize Nodes having larger storage capacity.
- Volume metrics: Usage stats are exported as Prometheus metrics from `kubelet`.
- [Volume Expansion](https://kubernetes-csi.github.io/docs/volume-expansion.html): Volumes can be expanded by editing `PersistentVolumeClaim` objects.

### Planned features

- [Snapshot](https://kubernetes-csi.github.io/docs/snapshot-restore-feature.html): When we want it.

Programs
--------

A diagram of components is available in [docs/design.md](docs/design.md#diagram).

This repository contains these programs:

- `topolvm-controller`: CSI controller service.
- `topolvm-scheduler`: A [scheduler extender](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/scheduling/scheduler_extender.md) for TopoLVM.
- `topolvm-node`: CSI node service.
- `lvmd`: gRPC service to manage LVM volumes.

`lvmd` is a standalone program that should be run on Node OS as a systemd service.
Other programs are packaged into [a container image](https://quay.io/repository/cybozu/topolvm).

Getting started
---------------

A demonstration of TopoLVM running on [kind (Kubernetes IN Docker)][kind] is available at [example](example/) directory.

For production deployments, see [deploy](deploy/) directory.

User manual is at [docs/user-manual.md](docs/user-manual.md).
If you want to use TopoLVM on [Rancher/RKE](https://rancher.com/docs/rke/latest/en/), see [docs/rancher.md](docs/rancher.md).

Documentation
-------------

[docs](docs/) directory contains documents about designs and specifications.

Docker images
-------------

Docker images are available on [Quay.io](https://quay.io/repository/cybozu/topolvm)

[releases]: https://github.com/topolvm/topolvm/releases
[CSI]: https://github.com/container-storage-interface/spec
[kind]: https://github.com/kubernetes-sigs/kind
