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
  --set-file cvData.json=./data/data.json \
  --set cvData.photoBase64=$(base64 -w0 ./data/photo.jpg) \
  cv ./ops/helm \
  -n cv 
```
