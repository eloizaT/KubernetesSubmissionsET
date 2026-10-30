# Exercise 1.2

Built and published the Docker image based on this version of Docker compose file. 

image name: **eloizat/webserver**


### Deployment command used:
``` bash
kubectl create deployment webserver-dep --image eloizat/webserver
```

### Set custom port:
``` bash
kubectl set env deployment/webserver-dep PORT=9090
```

### Verified with:
``` bash
kubectl logs -f webserver-dep-******
```
