<script lang='ts'>
  import Button from 'ui/button/button.svelte';
  import ConnectionForm from '../ConnectionForm.svelte';
  import handleSubmit from './actions';
  import Separator from 'ui/separator/separator.svelte';
  import { createForm } from 'felte';

  const {form} = createForm({onSubmit: handleSubmit})
</script>

<div class="flex gap4 items-center">
  <a href="/connections">
    <i class="i-carbon:arrow-left"></i>
  </a>
  <h3 class="text-2xl">Add new connection</h3>
</div>

<Separator class="my4" />

<form
  class="flex flex-col gap-4"
  autocomplete="off"
  use:form
>
  <ConnectionForm />
  <Button type="submit">Create</Button>
</form>
