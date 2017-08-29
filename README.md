## Description

  * Simple golang app to test web containers
  * Uses an environment variable to display where container is hosted
  
## How to run
 
  ```
  # Kubernetes
  kubectl apply -f go-web.yaml

  # ACI Connector
  kubectl apply -f go-web-aci.yaml

  # Docker
  docker run -d --name go-web -e "HOST_PLATFORM=localhost" -e "BACK_COLOR=blue" -p 8080:8080 chzbrgr71/go-web
  ```