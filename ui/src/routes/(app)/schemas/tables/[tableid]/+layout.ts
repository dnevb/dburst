import { redirect } from '@sveltejs/kit';
import { GetTableInfoRequest } from 'pb/dburst/sqlclient/v1/service';
import type { LayoutLoad } from './$types';

export const load: LayoutLoad = async (e) => {
  const { api } = await e.parent();

  const selectedConn = sessionStorage.getItem('conn');
  if (!selectedConn) throw redirect(307, '/schemas');

  const req = GetTableInfoRequest.create({
    connection: selectedConn,
    id: e.params.tableid,
  });

  return api?.sqlclient.getTableInfo(req).response;
};
