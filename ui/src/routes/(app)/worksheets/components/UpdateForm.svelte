<script lang="ts">
  import { createForm } from 'felte';
  import Editor from 'routes/(app)/editor/editor.svelte';
  import Button from 'ui/button/button.svelte';
  import * as Dialog from 'ui/dialog';
  import Input from 'ui/input/input.svelte';
  import Label from 'ui/label/label.svelte';
  import { selectTriggerVariants } from 'ui/select/variants';
  import { parse, stringify } from 'yaml';
  import actions from '../actions';

  let open = false;

  const { form, setInitialValues, data, setData } = createForm({
    onSubmit: (values) =>
      actions
        .updateWorksheet({
          ...values,
          variables: parse(values.variables),
        })
        .then(() => (open = false)),
  });

  export const openDialog = (v: any) => {
    open = true;
    setInitialValues(v);
    console.log(v);
  };
  let width: number;
</script>

<Dialog.Root bind:open>
  <Dialog.Content>
    <Dialog.Header>
      <Dialog.Title>Update worksheet</Dialog.Title>
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
      <div bind:clientWidth={width}>
        <Label>Variables (yaml)</Label>
        <Editor
          height={200}
          type="yaml"
          {width}
          content={stringify($data.variables)}
          on:update={(e) => setData('variables', e.detail)}
        />
      </div>
      <Button variant="outline" type="submit">Update</Button>
    </form>
  </Dialog.Content>
</Dialog.Root>
