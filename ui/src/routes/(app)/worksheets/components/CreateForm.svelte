<script lang="ts">
  import { createForm } from 'felte';
  import Button from 'ui/button/button.svelte';
  import * as Dialog from 'ui/dialog';
  import Input from 'ui/input/input.svelte';
  import Label from 'ui/label/label.svelte';
  import { selectTriggerVariants } from 'ui/select/variants';
  import actions from '../actions';

  let open = false;

  const { form, setInitialValues } = createForm({
    onSubmit: (v) =>
      actions.createWorksheet(v).then(() => (open = false)),
  });
  export const openDialog = (v?: any) => {
    open = true;
    v && setInitialValues(v);
  };
</script>

<Dialog.Root bind:open>
  <Dialog.Content>
    <Dialog.Header>
      <Dialog.Title>Create worksheet</Dialog.Title>
    </Dialog.Header>

    <form class="flex flex-col gap4" use:form>
      <div>
        <Label>Name*</Label>
        <Input name="name" required />
      </div>

      <div>
        <Label>Type*</Label>
        <select class={selectTriggerVariants()} name="type" required>
          <option value="sql">Sql worksheet</option>
          <option value="python">Python worksheet</option>
          <option value="folder">Folder</option>
        </select>
      </div>
      <Button variant="outline" type="submit">Create</Button>
    </form>
  </Dialog.Content>
</Dialog.Root>
