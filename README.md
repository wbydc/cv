# CV generator

Generates and serves CV from provided data.json and photo.jpg. Auto updates generated html on file change.

## Usage

Put your [experience](./types.go) and photo into data folder, then run 

### Docker

```
docker compose -f ops/docker-compose.yaml up --build -d
```

### Kubenetes

For helm chart configuration refer to [values.yaml](./ops/helm/values.yaml)

```
helm upgrade --install \
  -f values.yaml \
  cv ./ops/helm \
  -n cv

kubectl cp ./data/data.json <pod-name>:/data/data.json  
kubectl cp ./data/photo.jpg <pod-name>:/data/photo.jpg
```
