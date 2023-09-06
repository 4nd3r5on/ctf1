FROM node:latest

WORKDIR /app

# Installing npm libs
COPY package.json .
COPY package-lock.json .
RUN npm i

ENTRYPOINT [ "npm", "run", "dev" ]