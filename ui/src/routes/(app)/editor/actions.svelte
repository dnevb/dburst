<script lang="ts">
  import { invalidateAll } from '$app/navigation';
  import { ListConnectionsRequest } from 'pb/dburst/session/v1/service';
  import type { ConnectionInfo } from 'pb/dburst/session/v1/types';
  import { UpdateWorksheetRequest } from 'pb/dburst/worksheet/v1/service';
  import type { Worksheet } from 'pb/dburst/worksheet/v1/worksheet';
  import getServices from 'provider/services';
  import { createEventDispatcher, onMount } from 'svelte';
  import Button from 'ui/button/button.svelte';
  import { selectTriggerVariants } from 'ui/select/variants';

  export let connection: string;
  let connections: ConnectionInfo[] = [];

  const dispatch = createEventDispatcher();

  onMount(async () => {
    const svc = getServices();
    const req = ListConnectionsRequest.create({});
    const res = await svc.session.listConnections(req).response;

    connections = res.connections;
  });
</script>

<div class="flex items-center gap-2">
  <select
    class={selectTriggerVariants({ size: 'sm' })}
    name="connection"
    value={connection}
    on:change={(e) => dispatch('update', e.currentTarget.value)}
  >
    <option value=""></option>
    {#each connections as conn}
      <option value={conn.id}>{conn.name}</option>
    {/each}
  </select>
  <Button size="sm" on:click={() => dispatch('run')}>
    <i class="i-carbon:play" />
    Run
  </Button>
</div>
