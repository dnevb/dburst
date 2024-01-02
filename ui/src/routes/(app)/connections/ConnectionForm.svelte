<script lang="ts">
  import Input from 'ui/input/input.svelte';
  import Label from 'ui/label/label.svelte';
  import { selectTriggerVariants } from 'ui/select/variants';
</script>

<div>
  <Label>Connection name*</Label>
  <Input name="name" required />
</div>

<div>
  <Label>Engine*</Label>
  <select
    class={selectTriggerVariants()}
    name="credentials.scheme"
    required
  >
    <option value=""></option>
    <option value="postgres">Postgres</option>
    <option value="mysql">Mysql</option>
  </select>
</div>

<div class="flex gap2">
  <div class="flex-1">
    <Label>Host*</Label>
    <Input name="credentials.host" required />
  </div>
  <div>
    <Label>Port</Label>
    <Input name="credentials.port" type="number" />
  </div>
</div>

<div>
  <Label>Database*</Label>
  <Input name="credentials.database" required />
</div>

<div class="flex gap-2">
  <div class="flex-1">
    <Label>Username*</Label>
    <Input name="credentials.username" required />
  </div>
  <div class="flex-1">
    <Label>Password*</Label>
    <Input name="credentials.password" type="password" required />
  </div>
</div>
