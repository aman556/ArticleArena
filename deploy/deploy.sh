#!/usr/bin/env bash

set -o nounset
set -o pipefail

kubectl delete cm article-arena-cm || true
    kubectl create -f mysql-database/database-configmap.yaml
    if [ $? -eq 0 ]; then
      echo "Database configmap created"
    else
      echo "Failed to create Database configmap."
      exit 1
    fi

printf "\n"

kubectl delete statefulset articlearenadb-statefulset || true
    kubectl create -f mysql-database/database-statefulset.yaml
    if [ $? -eq 0 ]; then
      echo "Database statefulset created"
    else
      echo "Failed to create Database statefulset."
      exit 1
    fi

printf "\n"

kubectl delete svc articlearenadb-svc || true
    kubectl create -f mysql-database/database-service.yaml
    if [[ $? -eq 0 ]]; then
      echo "Database Service created"
    else
      echo "Failed to create Database Service."
      exit 1
    fi

printf "\n"

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

# EXIT_CODE="$(kubectl port-forward )"
# if [[ ${EXIT_CODE} != "0" ]]; then
#   echo "Failed to create backend-server-service."
#   exit "${EXIT_CODE}"
# fi
