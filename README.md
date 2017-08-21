## Description

  * Simple golang app to test web containers
  
## How to run
 
  ```
  docker run -d --name go-web -e "HOST_PLATFORM=localhost" -p 8080:8080 chzbrgr71/go-web
  ```