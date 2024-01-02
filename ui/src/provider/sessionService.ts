import { SessionServiceClient } from '../pb/dburst/session/v1/service.client';
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport';
import { PUBLIC_API_URL } from '$env/static/public';

const transport = new GrpcWebFetchTransport({
  baseUrl: PUBLIC_API_URL,
  meta: {},
});

export default new SessionServiceClient(transport);
