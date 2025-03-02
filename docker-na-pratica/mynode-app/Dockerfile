ARG NODE_VERSION=21.1.0

FROM node:${NODE_VERSION}
#FROM node:${NODE_VERSION}-alpine # cerca de 260MB
#FROM node:${NODE_VERSION}-slim # cerca de 380MB

ENV PORT=3001
ENV MESSAGE="Hello Docker"

RUN apt-get update && \
    apt-get install -y curl && \
    apt-get install -y vim && \
    rm -rf /var/lib/apt/lists/*


HEALTHCHECK --interval=5s --retries=3 --start-interval=5s --timeout=10s \
CMD [ "curl", "-f", "http://localhost:3001" ] || exit 1

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .

RUN useradd -m mynode
USER mynode

# VOLUME [ "/data" ]

EXPOSE 3001

CMD ["node", "index.js"]



