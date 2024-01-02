import { ListConnectionsRequest } from 'pb/dburst/session/v1/service';
import type { PageLoad } from './$types';

export const load: PageLoad = async (e) => {
  const parent = await e.parent();
  const api = parent.api!;

  const req = ListConnectionsRequest.create();
  return api.session.listConnections(req).response;
};
