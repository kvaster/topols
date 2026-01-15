# https://github.com/docker/buildx/releases
BUILDX_VERSION := 0.30.1
# https://github.com/cert-manager/cert-manager/releases
CERT_MANAGER_VERSION := v1.17.4
# https://github.com/helm/chart-testing/releases
CHART_TESTING_VERSION := 3.14.0
# https://github.com/containernetworking/plugins/releases
CNI_PLUGINS_VERSION := v1.9.0
# https://github.com/GoogleContainerTools/container-structure-test/releases
CONTAINER_STRUCTURE_TEST_VERSION := 1.22.0
# https://github.com/Mirantis/cri-dockerd/releases
CRI_DOCKERD_VERSION := v0.3.21
# https://github.com/kubernetes-sigs/cri-tools/releases
CRICTL_VERSION := v1.35.0
# https://github.com/golangci/golangci-lint/releases
GOLANGCI_LINT_VERSION := v2.7.2
# https://github.com/norwoodj/helm-docs/releases
HELM_DOCS_VERSION := 1.14.2
# https://github.com/helm/helm/releases
HELM_VERSION := 3.19.4
# It is set by CI using the environment variable, use conditional assignment.
KUBERNETES_VERSION ?= 1.34.3
# https://github.com/protocolbuffers/protobuf/releases
PROTOC_VERSION := 33.2
# https://github.com/mikefarah/yq/releases
YQ_VERSION := 4.50.1

# Tools versions which are defined in go.mod
SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
CONTROLLER_RUNTIME_VERSION := $(shell awk '/sigs\.k8s\.io\/controller-runtime/ {print substr($$2, 2)}' $(SELF_DIR)/go.mod)
CONTROLLER_TOOLS_VERSION := $(shell awk '/sigs\.k8s\.io\/controller-tools/ {print substr($$2, 2)}' $(SELF_DIR)/go.mod)
GINKGO_VERSION := $(shell awk '/github.com\/onsi\/ginkgo\/v2/ {print $$2}' $(SELF_DIR)/go.mod)
PROTOC_GEN_DOC_VERSION := $(shell awk '/github.com\/pseudomuto\/protoc-gen-doc/ {print substr($$2, 2)}' $(SELF_DIR)/go.mod)
PROTOC_GEN_GO_GRPC_VERSION := $(shell awk '/google.golang.org\/grpc\/cmd\/protoc-gen-go-grpc/ {print substr($$2, 2)}' $(SELF_DIR)/go.mod)
PROTOC_GEN_GO_VERSION := $(shell awk '/google.golang.org\/protobuf/ {print substr($$2, 2)}' $(SELF_DIR)/go.mod)

ENVTEST_KUBERNETES_VERSION := $(shell echo $(KUBERNETES_VERSION) | cut -d "." -f 1-2).0

# CSI sidecar versions
# https://github.com/kubernetes-csi/external-provisioner/releases
EXTERNAL_PROVISIONER_VERSION := 5.3.0
# https://github.com/kubernetes-csi/external-resizer/releases
EXTERNAL_RESIZER_VERSION := 2.0.0
# https://github.com/kubernetes-csi/external-snapshotter/releases
EXTERNAL_SNAPSHOTTER_VERSION := 8.4.0
# https://github.com/kubernetes-csi/livenessprobe/releases
LIVENESSPROBE_VERSION := 2.17.0
# https://github.com/kubernetes-csi/node-driver-registrar/releases
NODE_DRIVER_REGISTRAR_VERSION := 2.15.0
