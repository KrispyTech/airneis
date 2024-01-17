FROM alpine:3.19 as base
RUN apk add --update npm

FROM base as deps
RUN apk add --no-cache libc6-compat
WORKDIR /app

COPY ./package.json ./package-lock.json ./
RUN  npm install

FROM base as builder
WORKDIR /app
COPY --from=deps /app/node_modules ./node_modules
COPY . .

ENV NEXT_TELEMETRY_DISABLED 1

RUN npm run build

FROM base as runner
WORKDIR /app

COPY --from=builder /app/.next ./.next
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/package.json ./package.json

CMD ["npm", "start"]