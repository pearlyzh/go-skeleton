# Go Skeleton

Imagine this is a boostrap project for your backend app.
This reduces boilerplate codes and help you focus on your business logic.

### Install dependencies
```bash
make install
```

### Build image
```bash
docker build -t go-skeleton --progress=plain .
```

### Run backing services
```bash
docker-compose -f mysql-docker-compose.yml up -d 
```

### Run go-skeleton service
```bash
docker run -d -p 9090:9090 -e CONFIG_PATH=config/local.yaml go-skeleton
```