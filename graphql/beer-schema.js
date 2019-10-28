const graphql = require('graphql')
const psql = require('./database').psql
const { 
    GraphQLSchema,
    GraphQLObjectType,
    GraphQLString,
    GraphQLInt,
    GraphQLList,
} = graphql

const UserType = new GraphQLObjectType({
    name: 'User',
    fields: {
        id: { type: GraphQLInt },
        username: { type: GraphQLString }
    }
})

const query = new GraphQLObjectType({
    name: 'Query',
    fields: {
        hello: {
            type: GraphQLString,
            resolve: function() {
                return "hello"
            }
        },
        users: {
            type: new GraphQLList(UserType),
            resolve: function() {
                return psql.manyOrNone('select id, username from users')
            }
        },
        user: {
            type: UserType,
            args: {
                id: { type: GraphQLInt }
            },
            resolve: function(_, args) {
                return psql.oneOrNone('select id, username from users where id = $1', args.id)
            }
        }
    }
})

const schema = new GraphQLSchema({
    query: query
})

module.exports = schema