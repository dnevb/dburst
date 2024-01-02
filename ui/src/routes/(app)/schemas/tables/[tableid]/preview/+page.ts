import { QueryRequest } from 'pb/dburst/sqlclient/v1/service';
import type { PageLoad } from './$types';
import knex from 'knex';

export const load: PageLoad = async (e) => {
  const { api, currentConn, table } = await e.parent();

  const b = knex({ client: 'pg' });
  const sql = b
    .withSchema(table!.schema)
    .from(table!.name)
    .limit(50)
    .toQuery();

  const req = QueryRequest.create({
    connection: currentConn!,
    sql,
  });
  return api?.sqlclient.query(req).response;
};