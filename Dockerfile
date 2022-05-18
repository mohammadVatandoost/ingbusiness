FROM BASE_IMAGE_TAG as base_image

FROM debian:stable-20220509 as final

ENV TZ 'Asia/Tehran'

WORKDIR /app

COPY --from=base_image /build-app/service .

ENTRYPOINT ["./service", "serve"]

