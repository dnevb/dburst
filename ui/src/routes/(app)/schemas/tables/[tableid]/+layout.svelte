<script lang="ts">
  import type { PageData } from './$types';
  import PageHead from './components/PageHead.svelte';

  export let data: Required<PageData>;
</script>

<div class="flex flex-col h-full">
  <PageHead table={data.table} />
  <div class="p8 flex-1 overflow-auto overflow-x-auto">
    <slot />
  </div>
</div>
