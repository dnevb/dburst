<script lang="ts">
  import { selectTriggerVariants } from 'ui/select/variants';
  import type { PageData } from './$types';
  import SchemaItem from './schemaItem.svelte';
  import { invalidate, invalidateAll } from '$app/navigation';

  export let data: PageData;
  let conn = sessionStorage.getItem('conn');

  export const snapshot = {
    capture: () => sessionStorage.setItem('conn', conn || ''),
    restore: () => (conn = sessionStorage.getItem('conn')),
  };
</script>

<div class="flex h-full overflow-hidden">
  <div class="border-r w-64 p4 flex flex-col">
    <div>
      <select
        name="connection"
        class={selectTriggerVariants({})}
        bind:value={conn}
        on:change={async (e) => {
          conn = e.currentTarget.value;
          await invalidateAll();
        }}
      >
        <option value=""></option>
        {#each data.connections as conn}
          <option value={conn.id}>{conn.name}</option>
        {/each}
      </select>
    </div>
    <div class="overflow-y-auto flex-1">
      {#if conn && data.schemas.length}
        <ul class="ActionList">
          {#each data.schemas as schema}
            <li>
              <SchemaItem connection={conn} {schema} />
            </li>
          {/each}
        </ul>
      {/if}
    </div>
  </div>
  <div class="flex-1">
    <slot />
  </div>
</div>
