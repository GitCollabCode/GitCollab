FROM node:alpine
RUN mkdir /app
WORKDIR /app
RUN apk --no-cache add curl
COPY package.json /app
RUN npm install
COPY . /app
CMD ["npm", "start"]
