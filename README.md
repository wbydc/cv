# CV generator

Generates and serves CV from provided data.json and photo.jpg

## Usage

Put your experience and photo into data folder, then run 

### Docker

```
docker compose -f ops/docker-compose.yaml up --build -d
```

### Kubenetes

```
bash scripts/deploy.sh $NAMESPACE 
```
