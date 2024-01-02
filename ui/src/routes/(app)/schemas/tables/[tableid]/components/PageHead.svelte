<script lang="ts">
  import { goto } from '$app/navigation';
  import { page } from '$app/stores';
  import type { Table } from 'pb/dburst/sqlclient/v1/types';
  import ObjectIcon from 'routes/(app)/schemas/ObjectIcon.svelte';
  import Separator from 'ui/separator/separator.svelte';
  import * as Tabs from 'ui/tabs';

  export let table: Table;
  $: baseurl = `/schemas/tables/${table.id}`;
</script>

<div class="px-xl p-xl space-y-2">
  <div class="text-3xl font-semibold">
    <ObjectIcon type={table.type} />
    {table.schema}.{table.name}
  </div>
  <div>
    <p>{table?.comment}</p>
    <span>
      <i class="i-carbon:table"></i>
      {table?.type}
    </span>
    <span>
      <i class="i-carbon:chart-column-target"></i>
      {table?.rowCount}
    </span>
    <span>
      <i class="i-carbon:data-base"></i>
      {table?.totalSize}
    </span>
  </div>

  <Tabs.Root
    value={$page.url.pathname.endsWith('/preview') ? '/preview' : ''}
    onValueChange={(v) => goto(`${baseurl}${v}${$page.url.search}`)}
  >
    <Tabs.List>
      <Tabs.Trigger value="">Schema</Tabs.Trigger>
      <Tabs.Trigger value="/preview">Preview</Tabs.Trigger>
    </Tabs.List>
  </Tabs.Root>
</div>
<Separator class="" />
