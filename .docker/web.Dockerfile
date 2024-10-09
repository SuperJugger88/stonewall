FROM oven/bun:1.1.29-alpine AS deps

WORKDIR /var/www

COPY stonewall-web/package.json stonewall-web/bun.lockb ./

RUN bun install --frozen-lock

FROM oven/bun:1.1.29-alpine AS builder

WORKDIR /app

COPY --from=deps /var/www/node_modules node_modules
COPY stonewall-web .

RUN bun run build

FROM oven/bun:1.1.29-alpine

ENV HOST=0.0.0.0
ENV PORT=3000

WORKDIR /srv

COPY --from=builder --chown=1000 /app/.next/standalone .
COPY --from=builder --chown=1000 /app/.next/static ./.next/static

EXPOSE 3000

CMD ["bun", "run", "server.js"]