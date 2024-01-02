<script lang="ts">
  import type { PageData } from './$types';

  import { createForm } from 'felte';
  import Button from 'ui/button/button.svelte';
  import ConnectionForm from '../ConnectionForm.svelte';
  import actions from './actions';
  import Separator from 'ui/separator/separator.svelte';

  export let data: PageData;
  const connection = data.connection!;

  const { form } = createForm({
    initialValues: connection,
    onSubmit: actions.update,
  });
</script>

<div class="flex gap4 items-center">
  <a href="/connections">
    <i class="i-carbon:arrow-left"></i>
  </a>
  <h3 class="text-2xl">
    Edit {connection.name} connection
  </h3>
</div>

<Separator class="my4" />

<form class="flex flex-col gap-4" autocomplete="off" use:form>
  <ConnectionForm />
  <div class="flex gap-4">
    <Button class="flex-1" type="submit">Update</Button>

    <Button
      variant="danger"
      on:click={() => actions.delete(connection.id)}>Delete</Button
    >
  </div>
</form>
