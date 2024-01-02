<script lang="ts">
  import type { PageData } from './$types';
  import ColumnIcon from './components/columnIcon.svelte';
  import * as Table from 'ui/table';

  export let data: Required<PageData>;
</script>

<div class="TableOverflowWrapper">
  <Table.Root class="Table">
    <Table.Header>
      <Table.Row>
        <Table.Head>Name</Table.Head>
        <Table.Head>Type</Table.Head>
        <Table.Head>Ordinal</Table.Head>
        <Table.Head>Nullable</Table.Head>
        <Table.Head>Comment</Table.Head>
      </Table.Row>
    </Table.Header>
    <Table.Body>
      {#each data.columns as col}
        <Table.Row>
          <Table.Cell>{col.name}</Table.Cell>
          <Table.Cell>
            <ColumnIcon type={col.dataType} />
            {col.dataType}
          </Table.Cell>
          <Table.Cell>{col.ordinalPosition}</Table.Cell>
          <Table.Cell>
            {#if col.isNullable}
              <i class="i-carbon:checkmark"></i>
            {/if}
          </Table.Cell>
          <Table.Cell>{col.comment}</Table.Cell>
        </Table.Row>
      {/each}
    </Table.Body>
  </Table.Root>
</div>
