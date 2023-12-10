# frontend/Dockerfile
FROM node:latest

# Set the working directory inside the container
WORKDIR /app

COPY package.json ./
COPY yarn.lock ./  

RUN yarn install     

COPY . .

# Build the Next.js application
RUN yarn build       

EXPOSE 3000

# Start the application
CMD ["yarn", "start"] 
