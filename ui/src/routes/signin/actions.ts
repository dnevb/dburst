import { goto } from '$app/navigation';
import getServices from 'provider/services';
import type FormSubmitHandler from 'utils/handleSubmit';
import { SigninRequest } from '../../pb/dburst/session/v1/service';
import { SessionServiceClient } from 'pb/dburst/session/v1/service.client';
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport';
import { PUBLIC_API_URL } from '$env/static/public';

export const handleSubmit = async (data: any) => {
  const svc = new SessionServiceClient(
    new GrpcWebFetchTransport({
      baseUrl: PUBLIC_API_URL,
    }),
  );

  const req = SigninRequest.create(data);

  const { response } = await svc.signin(req);
  localStorage.setItem('api_token', response.token);
  await goto('/');
};
