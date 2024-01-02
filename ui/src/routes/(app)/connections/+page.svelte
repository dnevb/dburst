<script lang="ts">
  import { buttonVariants } from 'ui/button';
  import { Input } from 'ui/input';
  import Separator from 'ui/separator/separator.svelte';
  import type { PageData } from './$types';
  import a from 'ui/button/button.svelte';

  export let data: Required<PageData>;
</script>

<div class="space-y-1">
  <h2 class="text-2xl font-semibold">Connections</h2>
  <p class="text-sm">Manage how you connect to your databases</p>
</div>
<Separator class="my4" />
<div class="content">
  <div class="filters">
    <Input placeholder="Search..." />

    <a href="/connections/create" class={buttonVariants()}>
      <i class="i-carbon:add h-6 w-6"></i>
      New Connection
    </a>
  </div>
  <ul class="py-4">
    {#each data.connections as conn}
      <li>
        <a
          class={buttonVariants({
            variant: 'ghost',
            class: 'w-full gap3 justify-start h-auto',
          })}
          href={`/connections/${conn.id}`}
        >
          <img
            src={`/connector/${conn.credentials?.scheme}.png`}
            alt=""
            class="w-10 h-10"
          />
          <div>
            <h3 class="text-lg font-semibold">{conn.name}</h3>
            <span class="text-sm">{conn.credentials?.host}</span>
          </div>
        </a>
      </li>
    {/each}
  </ul>
</div>

<style>
  .content {
    --apply: 'pt-4';
  }
  .filters {
    --apply: 'flex items-center gap-4';
  }
</style>
