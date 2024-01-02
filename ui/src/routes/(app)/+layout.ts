import { redirect } from '@sveltejs/kit';
import { CurrentSessionRequest } from 'pb/dburst/session/v1/service';
import { createServices, services } from 'provider/services';
import type { LayoutLoad } from './$types';

export const ssr = false;

export const load: LayoutLoad = async () => {
  const token = localStorage.getItem('api_token')!;

  const api = createServices(token);
  services.set(api);

  const res = await api.session
    .currentSession(CurrentSessionRequest.create())
    .response.catch(() => {
      throw redirect(307, '/signin');
    });

  return {
    session: res.session!,
    api,
  };
};
