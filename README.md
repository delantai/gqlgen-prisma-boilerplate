# grVvy: Core

This is the Grvvy Core Repo, a serverless, yoga-graphql backend.

## Quickstart

#### 1. Install Prisma via Homebrew or Yarn
```
brew tap prisma/prisma
brew install prisma
```

<Details>
<Summary><b>Alternative</b>: Install with NPM or Yarn</Summary>
```
npm install -g prisma
# or
yarn global add prisma
```

#### 2. Connect Prisma to a database
Run the following command to setup Prisma
```
docker-compose up
```

#### 3. Deploy Prisma API if it is not running locally.
To deploy your Prisma API, run the following command:
```
prisma deploy
```

#### 4. Install GO dependencies.
Run the following command to ensure GO dependencies are installed:
```
dep ensure
```

#### 5. Generate gqlgen files.
Build any necessary graphql files using gqlgen (this isn't quite working):
```
go generate ./...
```

To rebuild stubs, remove `internal/build/gqlgen/.resolver_stubs.go` and run:
```
go run scripts/gqlgen.go -v
```

#### 6. Run the GO server.
Now execute the script with the following command:
```
go run cmd/gqlgen-prisma-boilerplate/main.go
```
