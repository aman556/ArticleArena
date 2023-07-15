#!/usr/bin/env bash

set -o nounset
set -o pipefail
# set -e

RELEASE_TAG="$1"
make -C ../ build-server-image release-tag="${RELEASE_TAG}"

sed -i '' "s/articlearena:.*/articlearena:${RELEASE_TAG}/g" backend-server/backend-server-deployment.yaml


kubectl delete ServiceAccount mongo-account && kubectl delete ClusterRole mongo-role && kubectl delete ClusterRoleBinding mongo_role_binding || true
    kubectl create -f mongo-db-database/rbac.yaml
    if [ $? -eq 0 ]; then
      echo "Database rbac created"
    else
      echo "Failed to create rbac."
      exit 1
    fi

printf "\n"

kubectl delete statefulset mongodb-database || true
    kubectl create -f mongo-db-database/statefulset.yaml
    if [ $? -eq 0 ]; then
      echo "Database statefulset created"
    else
      echo "Failed to create Database statefulset."
      exit 1
    fi

printf "\n"

kubectl delete svc mongo || true
    kubectl create -f mongo-db-database/service.yaml
    if [[ $? -eq 0 ]]; then
      echo "Database Service created"
    else
      echo "Failed to create Database Service."
      exit 1
    fi

printf "\n"

sleep 30s

kubectl delete deployment backend-server-deployment-article-arena || true
    kubectl create -f backend-server/backend-server-deployment.yaml
    if [[ $? -eq 0 ]]; then
      echo "Backend server deployment created"
    else
      echo "Failed to create Backend server deployment."
      exit 1
    fi

printf "\n"

kubectl delete svc backend-server-article-arena-svc || true
kubectl create -f backend-server/backend-server-service.yaml
    if [[ $? -eq 0 ]]; then
      echo "Backend server service created"
    else
      echo "Failed to create Backend server service."
      exit 1
    fi

BACKEND_SERVER_POD="$(kubectl get pod | grep backend-server-deployment-article-arena | awk '{print $1}')"
kubectl port-forward "${BACKEND_SERVER_POD}" 8081:8081 

