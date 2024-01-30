FROM oven/bun as base
WORKDIR /app

COPY . .
RUN  bun install

FROM base as dev
WORKDIR /app

CMD ["bun", "run", "dev"]

FROM alpine:3.19 as prod
WORKDIR /app

RUN apk add --update npm
RUN apk add --no-cache libc6-compat

COPY . .
RUN  npm install --omit=dev

ENV NEXT_TELEMETRY_DISABLED 1

RUN npm run build

CMD ["npm", "start"]