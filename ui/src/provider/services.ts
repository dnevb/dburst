import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport';
import { SessionServiceClient } from 'pb/dburst/session/v1/service.client';
import { get, writable } from 'svelte/store';
import { SqlClientServiceClient } from 'pb/dburst/sqlclient/v1/service.client';
import { WorksheetServiceClient } from 'pb/dburst/worksheet/v1/service.client';
import { PUBLIC_API_URL } from '$env/static/public';

export const services = writable<ReturnType<typeof createServices>>();

export function createServices(token?: string) {
  const transport = new GrpcWebFetchTransport({
    baseUrl: PUBLIC_API_URL,
    meta: { authorization: token || '' },
  });

  return {
    sqlclient: new SqlClientServiceClient(transport),
    session: new SessionServiceClient(transport),
    worksheet: new WorksheetServiceClient(transport),
  };
}
createServices();

function getServices() {
  return get(services);
}

export default getServices;
