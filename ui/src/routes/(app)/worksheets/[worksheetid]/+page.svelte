<script lang="ts">
  import { page } from '$app/stores';
  import {
    RunWorksheetRequest,
    RunWorksheetResponse,
  } from 'pb/dburst/worksheet/v1/service';
  import getServices from 'provider/services';
  import { onMount } from 'svelte';
  import Separator from 'ui/separator/separator.svelte';
  import type { PageData } from './$types';
  import Actions from './actions.svelte';
  import Editor from './editor.svelte';
  import Results from './results.svelte';
  import ObjectIcon from '../components/WorksheetIcon.svelte';
  import Button from 'ui/button/button.svelte';
  import UpdateForm from '../components/UpdateForm.svelte';

  export let data: Required<PageData>;

  let response: RunWorksheetResponse | null = null;
  let editorRef: Editor;

  let height: number;
  let width: number;

  onMount(() => {
    page.subscribe((p) => {
      editorRef?.reset(
        p.data.worksheet?.content,
        p.data.worksheet?.type,
      );
      onCloseResults();
    });
  });

  const onRun = async () => {
    const svc = getServices();
    const req = RunWorksheetRequest.create({
      worksheetId: data.worksheet.id,
    });
    const res = await svc.worksheet.runWorksheet(req).response;

    if (!response)
      editorRef?.layout({
        width,
        height: height * 0.6,
      });

    response = res;
  };
  const onCloseResults = () => {
    response = null;
    editorRef?.layout({ width, height });
  };
  let updateformRef: UpdateForm;
</script>

<UpdateForm bind:this={updateformRef} />
<div class="flex flex-col h-full">
  <div class="flex justify-between px-sm">
    <div>
      <h3 class="text-2xl">
        <ObjectIcon type={data.worksheet?.type} />
        {data.worksheet?.name}
        <Button
          variant="ghost"
          size="none"
          on:click={() => updateformRef?.openDialog(data.worksheet)}
        >
          <i class="i-carbon:edit" />
        </Button>
      </h3>
      <p class="text-md">
        variables:
        {#each Object.keys(data.worksheet.variables) as varkey}
          <span class="border-2 rounded-xl px1">
            {varkey}: {data.worksheet.variables[varkey]}
          </span>
        {/each}
      </p>
    </div>

    <Actions worksheet={data.worksheet} on:run={onRun} />
  </div>

  <Separator class="mt2" />

  <div
    class="flex-1"
    bind:clientHeight={height}
    bind:clientWidth={width}
  >
    <Editor
      on:run={onRun}
      bind:this={editorRef}
      worksheet={data.worksheet}
      {height}
      {width}
    />
    {#if response}
      <div class="flex flex-col" style={`height:${height * 0.4}px`}>
        <Results {response} on:close={onCloseResults} />
      </div>
    {/if}
  </div>
</div>
