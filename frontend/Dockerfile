FROM node:16

ENV PORT 3000
RUN mkdir -p /app
WORKDIR /app
COPY *.json /app/
RUN yarn

COPY . /app

# Building app
RUN yarn build
EXPOSE 3000

# Running the app
CMD [ "yarn", "start" ]
