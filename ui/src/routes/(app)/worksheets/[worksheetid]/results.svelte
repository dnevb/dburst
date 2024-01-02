<script lang="ts">
  import type { RunWorksheetResponse } from 'pb/dburst/worksheet/v1/service';
  import { Value } from 'pb/google/protobuf/struct';
  import { createEventDispatcher } from 'svelte';
  import Button from 'ui/button/button.svelte';
  import Separator from 'ui/separator/separator.svelte';
  import * as Table from 'ui/table';

  export let response: RunWorksheetResponse;
  const dispatch = createEventDispatcher();
</script>

<header class="flex justify-between items-center pxsm border-y">
  <div>
    <h1>Worksheet results</h1>
  </div>
  <div>
    <Button
      variant="ghost"
      size="icon"
      on:click={() => dispatch('close')}
    >
      <i class="i-carbon:close"></i>
    </Button>
  </div>
</header>
<div class="flex-1 max-w-700px min-w-full overflow-auto">
  <Table.Root class="relative!">
    <Table.Header>
      <Table.Row class="sticky! top-0!">
        {#each response.columns as col}
          <Table.Head>{col.name}</Table.Head>
        {/each}
      </Table.Row>
    </Table.Header>
    <Table.Body>
      {#each response.rows as row}
        <Table.Row>
          {#each row.values as value}
            <Table.Cell
              class="truncate"
              title={Value.toJson(value)?.toString()}
              >{Value.toJson(value)}</Table.Cell
            >
          {/each}
        </Table.Row>
      {/each}
    </Table.Body>
  </Table.Root>
</div>
