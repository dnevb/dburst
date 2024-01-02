<script lang="ts">
  import { createForm } from 'felte';
  import Button from 'ui/button/button.svelte';
  import Input from 'ui/input/input.svelte';
  import Label from 'ui/label/label.svelte';
  import { handleSubmit } from './actions';

  const { form, isSubmitting } = createForm({
    onSubmit: handleSubmit,
  });
</script>

<form
  class="space-y-4 md:space-y-6"
  autocomplete="off"
  method="post"
  use:form
>
  <div>
    <Label for="username">Email</Label>
    <Input id="username" name="username" type="email" required />
  </div>
  <div>
    <Label for="password">Password</Label>
    <Input
      id="password"
      name="password"
      type="password"
      placeholder="••••••••"
      required
    />
  </div>
  <Button type="submit">
    {#if $isSubmitting}
      <i class="i-svg-spinners:180-ring"></i>
    {/if}
    Sign in
  </Button>
</form>
