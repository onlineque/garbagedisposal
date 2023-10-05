# garbagedisposal


This small project terminates Kubernetes pods in Succeeded and Failed status. It inspects all cluster pods each minute
and terminates found pods immediately.

## Installation

    helm install -n garbagedisposal --create-namespace garbagedisposal oci://ghcr.io/onlineque/garbagedisposal --version 0.1.0

### TODO:

- possibility to specify the pod age for each status, so only pods older than specified age are killed

