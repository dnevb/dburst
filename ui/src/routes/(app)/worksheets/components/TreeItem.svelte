<script lang="ts">
  import * as Dropdown from 'ui/dropdown-menu';
  import type { Worksheet } from 'pb/dburst/worksheet/v1/worksheet';
  import { createEventDispatcher } from 'svelte';
  import type { WithTree } from 'utils/createTree';
  import ObjectIcon from './WorksheetIcon.svelte';
  import WorksheetTree from './TreeList.svelte';
  import Button from 'ui/button/button.svelte';
  import { cn } from 'utils';
  import { page } from '$app/stores';

  export let wrk: WithTree<Worksheet>;
  export let level = 0;
  let open = false;
  let style = `padding-left: calc(${level}*1.7rem)`;
  let url = `/worksheets/${wrk.id}`;

  const dispatch = createEventDispatcher();
</script>

<li class="py0.5">
  {#if wrk.type != 'folder'}
    <Dropdown.Root>
      <div class="justify-between flex cursor-pointer">
        <a
          href={url}
          class={cn(
            'flex-1 truncate',
            $page.url.pathname.startsWith(url) && 'font-extrabold',
          )}
          title={wrk.name}
          {style}
        >
          <ObjectIcon type={wrk.type} />
          {wrk.name}
        </a>

        <Dropdown.Trigger asChild let:builder>
          <Button variant="ghost" size="none" builders={[builder]}>
            <i class="i-carbon:overflow-menu-horizontal" />
          </Button>
        </Dropdown.Trigger>
        <Dropdown.Content>
          <Dropdown.Group>
            <Dropdown.Item
              on:click={() => dispatch('openUpdate', wrk)}
              >Update</Dropdown.Item
            >
            <Dropdown.Item
              on:click={() => dispatch('openDelete', { id: wrk.id })}
              >Delete</Dropdown.Item
            >
          </Dropdown.Group>
        </Dropdown.Content>
      </div>
    </Dropdown.Root>
  {:else}
    <Dropdown.Root>
      <div class="justify-between flex cursor-pointer">
        <!-- svelte-ignore a11y-click-events-have-key-events -->
        <!-- svelte-ignore a11y-interactive-supports-focus -->
        <span
          role="button"
          on:click={() => (open = !open)}
          class="flex-1 truncate"
          title={wrk.name}
          {style}
        >
          {#if open}
            <i class="i-carbon:caret-down"></i>
          {:else}
            <i class="i-carbon:caret-right"></i>
          {/if}
          <i class="i-carbon:folder" />
          {wrk.name}
        </span>
        <Dropdown.Trigger asChild let:builder>
          <Button variant="ghost" size="none" builders={[builder]}>
            <i class="i-carbon:overflow-menu-horizontal" />
          </Button>
        </Dropdown.Trigger>
        <Dropdown.Content>
          <Dropdown.Group>
            <Dropdown.Item
              on:click={() =>
                dispatch('openCreate', { parentId: wrk.id })}
              >New child</Dropdown.Item
            >
            <Dropdown.Item
              on:click={() => dispatch('openUpdate', wrk)}
              >Update</Dropdown.Item
            >
            <Dropdown.Item
              on:click={() => dispatch('openDelete', { id: wrk.id })}
              >Delete</Dropdown.Item
            >
          </Dropdown.Group>
        </Dropdown.Content>
      </div>
    </Dropdown.Root>
    {#if open}
      <WorksheetTree
        on:openUpdate
        on:openDelete
        on:openCreate
        level={level + 1}
        worksheets={wrk._children}
      />
    {/if}
  {/if}
</li>
