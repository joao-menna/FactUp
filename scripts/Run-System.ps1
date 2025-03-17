$defaultLocation = Get-Location

Set-Location ./../

# Define paths
$rootPath = Get-Location
$packagesPath = Join-Path $rootPath "packages"
$frontendPath = Join-Path $packagesPath "frontend"

# Go to frontend package
Set-Location $frontendPath

# Install yarn
npm i -g yarn

# Install deps and build
yarn
yarn build

# Run Docker Compose
Set-Location $rootPath
docker compose up --build -d

# Return to default folder
Set-Location $defaultLocation
