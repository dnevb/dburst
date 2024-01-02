<script lang="ts">
  import { Value } from 'pb/google/protobuf/struct';
  import type { PageData } from './$types';
  import * as Table from 'ui/table';

  export let data: Required<PageData>;
</script>

<div class="TableOverflowWrapper max-w-700px min-w-full">
  <Table.Root class="Table">
    <Table.Header>
      <Table.Row>
        {#each data.columns as col}
          <Table.Head>{col.name}</Table.Head>
        {/each}
      </Table.Row>
    </Table.Header>
    <Table.Body>
      {#each data.rows as row}
        <Table.Row>
          {#each row.values as val}
            <Table.Cell>{Value.toJson(val)}</Table.Cell>
          {/each}
        </Table.Row>
      {/each}
    </Table.Body>
  </Table.Root>
</div>
