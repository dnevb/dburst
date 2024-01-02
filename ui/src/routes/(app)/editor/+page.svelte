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
  import {
    QueryRequest,
    QueryResponse,
  } from 'pb/dburst/sqlclient/v1/service';

  let connection = '';
  let content = '';

  let response: QueryResponse | null = null;
  let editorRef: Editor;

  let height: number;
  let width: number;

  const onRun = async () => {
    const svc = getServices();
    const req = QueryRequest.create({
      connection,
      sql: content,
    });
    const res = await svc.sqlclient.query(req).response;

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
</script>

<div class="flex flex-col h-full">
  <div class="flex justify-between px-sm">
    <div>
      <h3 class="text-2xl">Sql Editor</h3>
      <p class="text-md">sql</p>
    </div>

    <Actions
      {connection}
      on:update={(v) => (connection = v.detail)}
      on:run={onRun}
    />
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
      {content}
      on:update={(v) => (content = v.detail)}
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
