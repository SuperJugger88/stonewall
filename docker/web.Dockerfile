FROM denoland/deno:alpine-1.46.3

ARG GIT_REVISION
ENV DENO_DEPLOYMENT_ID=${GIT_REVISION}

WORKDIR /srv

COPY stonewall-web/package.json .

RUN --mount=type=cache,target=/deno-dir/npm \
    --mount=type=cache,target=/deno-dir/deps \
    --mount=type=cache,target=/deno-dir/gen \
    --mount=type=cache,target=/deno-dir/registries \
    deno cache package.json

COPY stonewall-web/ .

RUN --mount=type=cache,target=/deno-dir/npm \
        --mount=type=cache,target=/deno-dir/deps \
        --mount=type=cache,target=/deno-dir/gen \
        --mount=type=cache,target=/deno-dir/registries \
        deno task build

EXPOSE 3000

CMD ["deno", "task", "start"]