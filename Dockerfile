FROM alpine:3.19 as base
RUN apk add --update npm

FROM base as deps
RUN apk add --no-cache libc6-compat
WORKDIR /app

COPY ./front/package.json ./front/package-lock.json ./
RUN  npm install --production

FROM base as builder
WORKDIR /app
COPY --from=deps /app/node_modules ./node_modules
COPY ./front/ .

ENV NEXT_TELEMETRY_DISABLED 1

RUN npm run build

FROM base as runner
WORKDIR /app

COPY --from=builder /app/.next ./.next
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/package.json ./package.json

EXPOSE 3000

ENV PORT 3000

CMD ["npm", "start"]