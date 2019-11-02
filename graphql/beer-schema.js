const graphql = require('graphql')
const psql = require('./database').psql
const joinMonster = require('join-monster')

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

const BeerType = new GraphQLObjectType({
    name: 'Beer',
    fields: {
        id: { type: GraphQLInt },
        name: { type: GraphQLString }
    }
})
BeerType._typeConfig = {
    sqlTable: 'beers',
    uniqueKey: 'id'
}

const TastingsType = new GraphQLObjectType({
    name: 'Tasting',
    fields: {
        id: { type: GraphQLInt },
        name: { type: GraphQLString },
        beers: {
            type: GraphQLList(BeerType),
            sqlJoin: (tastingsTable, beersTable) => {
                return `${tastingsTable}.id = ${beersTable}.tastingid`
            }
        }
    }
})
TastingsType._typeConfig = {
    sqlTable: 'tastings',
    uniqueKey: 'id'
}

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
        },
        tastings: {
            type: new GraphQLList(TastingsType),
            resolve: function(parent, args, context, resolveInfo) {
                return joinMonster.default(resolveInfo, {}, sql => {
                    return psql.query(sql)
                })
            }
        }
    }
})

const schema = new GraphQLSchema({
    query: query
})

module.exports = schema