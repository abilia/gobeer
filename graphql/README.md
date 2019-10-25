# beer graphql

GraphQL for beer database

## Run locally
Install dependencies
```
npm install
```

Start server on port `8001`
```
node app.js
```

At `http://localhost:8001/graphql`, query for user with id `a`
```
{
    user(id: "a") {
        name
    }
}
```
