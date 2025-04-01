default_location=$(pwd)

cd ./../

# Define paths
root_path=$(pwd)
packages_path="$root_path/packages"
frontend_path="$packages_path/frontend"

# Go to frontend package
cd $frontend_path

# Install yarn
npm i -g yarn

# Install deps and build
yarn
yarn build

# Run Docker Compose
cd $root_path
sudo docker compose up --build -d

# Return to default folder
cd $default_location

