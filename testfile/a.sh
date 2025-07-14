#!/usr/bin/env bash
set -ex
set -o pipefail

helm upgrade metric-exporter ${HELM_REPO}charts/metric-exporter-0.1.0.tgz -f ${ROOTDIR}/cloud-dev/common/metric-exporter-values.yaml --wait --install --namespace $NAMESPACE
