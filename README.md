# Sample App Kubernetes

Sample App Deployment using Docker and Kubernetes

### BackendApp

The backend app using golang. Need to build docker image


### Docker Image

#### Backend APP

Build `backend-app` image using this commandline:

```
cd backend-app
docker build --tag backend-app:v0.1.1 .
```

#### Backend User

Build `backend-user` image using this commandline:

```
cd backend-user
docker build --tag backend-user:v0.1.1 .
```


### Test the docker image with run to the container

`docker run -it -p 3000:3000 backend-app:v0.1.1`

`docker run -it -p 3001:3001 backend-app:v0.1.1`

Open your browser and visit `http://localhost:3000/` and it should show Hello World

Open your browser and visit `http://localhost:3001/` and it should show Hello Backend User


### Deploy to Kubernetes

Create on `/etc/hosts`:

`127.0.0.1 yourdomain.local`

Run deployment:

`kubectl apply -f deployment`


### Setup /etc/hosts

```
127.0.0.1 yourdomain.local user.yourdomain.local
```