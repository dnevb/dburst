import { ListConnectionsRequest } from 'pb/dburst/session/v1/service';
import type { LayoutLoad } from './$types';
import { ListSchemasRequest } from 'pb/dburst/sqlclient/v1/service';

export const load: LayoutLoad = async (e) => {
  const parent = await e.parent();
  const api = parent.api!;

  const connReq = ListConnectionsRequest.create();
  const connRes = await api.session.listConnections(connReq).response;

  const selectedConn = sessionStorage.getItem('conn');
  const connection = connRes.connections.find(
    (c) => c.id == selectedConn,
  );

  const schemasReq = ListSchemasRequest.create({
    connection: selectedConn!,
  });
  const schemasRes = await api.sqlclient
    .listSchemas(schemasReq)
    .response.catch((e) => ({ schemas: [] }));

  return {
    connections: connRes.connections,
    schemas: schemasRes.schemas,
    currentConn: selectedConn,
    connection,
  };
};
