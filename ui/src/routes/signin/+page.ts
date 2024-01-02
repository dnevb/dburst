import { redirect } from '@sveltejs/kit';
import type { PageLoad } from './$types';

export const ssr = false;

export const load: PageLoad = () => {
  const token = localStorage.getItem('api_token');
  if (token) throw redirect(307, '/');
};
