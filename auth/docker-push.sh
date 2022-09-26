VERSION=1.0.2
docker build --platform linux/arm64 -t vijeth/goms:$VERSION .
docker push vijeth/goms:$VERSION