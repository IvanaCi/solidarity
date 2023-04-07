FROM node:current-alpine as build
WORKDIR /app
COPY package*.json /app/
RUN npm install -g @ionic/cli@latest
RUN npm install

RUN npm uninstall typescript @types/jest @typescript-eslint/eslint-plugin @typescript-eslint/parser @vue/cli-plugin-typescript @vue/eslint-config-typescript --save


COPY . /app/


RUN npm run-script build --prod

