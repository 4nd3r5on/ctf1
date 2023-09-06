FROM node:latest

COPY . .
RUN npm i
RUN npm run build

ENTRYPOINT [ "npm", "run", "preview" ]
