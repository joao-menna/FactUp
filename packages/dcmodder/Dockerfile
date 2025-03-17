FROM node:22-alpine AS build

WORKDIR /source

COPY . .

RUN yarn && \
    yarn build

FROM node:22-alpine AS run

WORKDIR /app

COPY --from=build /source/dist/*.js /source/dist/**/*.js ./

COPY --from=build /source/dist/commands ./commands

COPY --from=build /source/node_modules ./node_modules

CMD node registerCommands.js && node main.js
