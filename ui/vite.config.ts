import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';
import unocss from 'unocss/vite';
import tsconfig from 'vite-tsconfig-paths';
import { nodePolyfills } from 'vite-plugin-node-polyfills';

export default defineConfig({
  plugins: [unocss(), nodePolyfills(), sveltekit(), tsconfig()],
  build: {
    rollupOptions: {
      external: ['knex'],
    },
  },
});
