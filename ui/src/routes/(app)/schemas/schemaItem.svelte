<!-- svelte-ignore a11y-click-events-have-key-events -->
<script lang="ts">
  import { page } from '$app/stores';
  import { ListTablesRequest } from 'pb/dburst/sqlclient/v1/service';
  import type { Schema, Table } from 'pb/dburst/sqlclient/v1/types';
  import getServices from 'provider/services';
  import ObjectIcon from './ObjectIcon.svelte';

  export let connection: string;
  export let schema: Schema;
  let open = false;

  let tables: Table[] = [];

  async function toggle() {
    if (open) {
      open = false;
      tables = [];
      return;
    }
    open = true;
    const svc = getServices();
    const req = ListTablesRequest.create({
      connection,
      schema: schema.id,
    });
    const res = await svc.sqlclient.listTables(req).response;
    tables = res.tables;
  }
</script>

<div
  class="sectionDivider cursor-pointer"
  role="button"
  tabindex="0"
  on:click={toggle}
>
  <span class="sectionDivider-title">
    <span class="">
      {#if open}
        <i class="i-carbon:caret-down"></i>
      {:else}
        <i class="i-carbon:caret-right"></i>
      {/if}
    </span>
    {schema.name}
  </span>
</div>

{#if open}
  <ul class=" pl-4">
    {#each tables as object}
      <li class="item">
        <a class="truncate" href={`/schemas/tables/${object.id}`}>
          <ObjectIcon type={object.type} />
          <span class="item-label truncate">{object.name}</span>
        </a>
      </li>
    {/each}
  </ul>
{/if}
