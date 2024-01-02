<script lang="ts">
  import type { Worksheet } from 'pb/dburst/worksheet/v1/worksheet';
  import type { WithTree } from 'utils/createTree';
  import TreeItem from './TreeItem.svelte';

  export let worksheets: WithTree<Worksheet>[];
  export let level: number = 0;
</script>

<ul>
  {#each worksheets || [] as wrk}
    <TreeItem
      on:openUpdate
      on:openDelete
      on:openCreate
      {level}
      {wrk}
    />
  {/each}
</ul>
