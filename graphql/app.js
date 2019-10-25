const express = require('express')
const graphqlHTTP = require('express-graphql')
const tutorialSchema = require('./tutorial-schema.js')

const app = express()

app.use('/graphql', graphqlHTTP({
	schema: tutorialSchema,
	graphiql: true
}))

app.listen(8001, () => {
	console.log('App listening on port 8001')
})
