# frontend/Dockerfile

FROM node:latest

# Set the working directory inside the container
WORKDIR /app

# Copy package files separately to leverage Docker cache
COPY package.json .
COPY package-lock.json .

# Install dependencies using npm as non-root user
RUN npm install --unsafe-perm=true --allow-root

# Copy the rest of the application code
COPY . .

# Build the Next.js application
RUN npm run build

EXPOSE 3000

# Start the application
CMD ["npm", "start"]
