# Example service

For demo and testing purposes.

![](https://i.giphy.com/media/joV1k1sNOT5xC/giphy.webp)

Build native binary:

```sh
go build .
./example-service
```

Build container image with [Buildpacks](https://buildpacks.io):

```sh
pack build example-service --builder paketobuildpacks/builder:full
# Or without Pack CLI:
# docker run \
#     --rm \
#     --env CNB_PLATFORM_API=0.9 \
#     --volume $PWD:/workspace \
#     paketobuildpacks/builder:full \
#         /cnb/lifecycle/creator registry.khuedoan.com/khuedoan/example-service
docker run --rm --publish 8080:8080 example-service
```
