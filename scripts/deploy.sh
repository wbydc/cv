cp data/data.json ops/helm/data.json
cp data/photo.jpg ops/helm/photo.jpg

helm upgrade --install cv ops/helm -n cv
