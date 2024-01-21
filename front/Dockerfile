FROM oven/bun as base
WORKDIR /app

COPY . .
RUN  bun install

FROM base as dev
WORKDIR /app

CMD ["bun", "run", "dev"]

FROM base as builder

WORKDIR /app
ENV NEXT_TELEMETRY_DISABLED 1
RUN bun run build

FROM oven/bun as prod

WORKDIR /app

COPY --from=builder /app/.next ./.next
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/package.json ./package.json

CMD ["bun", "run", "start"]