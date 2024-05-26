# FROM node:latest as build-deps
# WORKDIR /usr/src/app
# COPY package.json yarn.lock ./
# RUN yarn
# COPY . ./
# CMD yarn dev

# # FROM nginx:1.12-alpine
# # COPY --from=build-deps /usr/src/app/dist /usr/share/nginx/html
# # EXPOSE 80
# # CMD ["nginx", "-g", "daemon off;"]

FROM node:21-alpine3.19

WORKDIR /app

COPY yarn.lock package.json ./

RUN yarn

COPY . .

EXPOSE 3000

CMD ["yarn", "dev", "--host", "0.0.0.0"]