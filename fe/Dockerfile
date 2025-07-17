FROM node:bookworm AS development

WORKDIR /ctx_fndr_fe/

COPY package*.json ./

RUN npm install 

COPY . .

RUN npm run build

FROM node:bookworm AS production

ARG NODE_ENV=production
ENV NODE_ENV=${NODE_ENV}}

WORKDIR /ctx_fndr_fe/

COPY package*.json ./

RUN npm install 

COPY . .

COPY /ctx_fndr_fe/dist ./dist

CMD ["node", "dist/main"]