const pgPromise = require('pg-promise');

const connStr = 'postgresql://gobeeruser:thisisthepassword@localhost:5454/beerdb';

const pgp = pgPromise({});
const psql = pgp(connStr);

exports.psql = psql;
