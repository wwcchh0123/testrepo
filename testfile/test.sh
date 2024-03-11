#!/usr/bin/env bash

err() {
  echo -e "[$(date +'%Y-%m-%dT%H:%M:%S.%N%z')] FAIL: $@" >&2
  exit 1
}


run-dynode-ecdn-integration() {
    info "Dynode cluster status"
    kubectl get pod -o wide -n $NS_NAME
    cd ${QBOXROOT}/qtest/ecdn/integration
    ginkgo -r -v -p --junit-report=integration.xml --output-dir="${ARTIFACTS}" --label-filter="integration" . -- -kube_namespace=$NS_NAME
}