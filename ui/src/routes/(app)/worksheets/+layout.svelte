<script lang="ts">
  import createTree from 'utils/createTree';
  import type { LayoutData } from './$types';
  import CreateForm from './components/CreateForm.svelte';
  import WorksheetTree from './components/TreeList.svelte';
  import UpdateForm from './components/UpdateForm.svelte';
  import Button from 'ui/button/button.svelte';

  export let data: Required<LayoutData>;
  let formRef: CreateForm;
  let updateformRef: UpdateForm;

  $: items = createTree(data.worksheets, 'parent');
</script>

<UpdateForm bind:this={updateformRef} />
<CreateForm bind:this={formRef} />

<div class="flex h-full max-h-screen">
  <div class="border-r w-64 p4">
    <Button
      variant="outline"
      class="wfull"
      on:click={() => formRef?.openDialog()}
    >
      <i class="i-carbon:add"></i>
      Create worksheet
    </Button>
    <div class="py4">
      <WorksheetTree
        worksheets={items}
        on:openCreate={(e) => formRef?.openDialog(e.detail)}
        on:openUpdate={(e) => updateformRef?.openDialog(e.detail)}
        on:openDelete
      />
    </div>
  </div>

  <div class="flex-1 h-full">
    <slot />
  </div>
</div>
