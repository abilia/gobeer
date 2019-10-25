const graphql = require('graphql')
const { GraphQLSchema, GraphQLObjectType, GraphQLString } = graphql

const fakeDatabase = {
    'a': {
        id: 'a',
        name: 'alice'
    },
    'b': {
        id: 'b',
        name: 'bob'
    },
}

const userType = new GraphQLObjectType({
    name: 'User',
    fields: {
        id: { type: GraphQLString },
        name: { type: GraphQLString }
    }
})

const query = new GraphQLObjectType({
    name: 'Query',
    fields: {
        user: {
            type: userType,
            args: {
                id: { type: GraphQLString }
            },
            resolve: function (_, {id}) {
                return fakeDatabase[id];
            }
        }
    }
})

const schema = new GraphQLSchema({
    query: query
})

module.exports = schema