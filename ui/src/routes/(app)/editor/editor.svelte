<script lang="ts">
  import type * as Monaco from 'monaco-editor/esm/vs/editor/editor.api';
  import type { Worksheet } from 'pb/dburst/worksheet/v1/worksheet';
  import monaco from 'provider/monaco';
  import { createEventDispatcher, onDestroy, onMount } from 'svelte';

  export let height: number;
  export let width: number;
  export let content: string;
  export let type: string = 'sql';

  let editor: Monaco.editor.IStandaloneCodeEditor;
  let model: monaco.editor.ITextModel;
  let editorContainer: HTMLElement;

  const dispatch = createEventDispatcher();

  export const layout = (dim: monaco.editor.IDimension) =>
    editor.layout(dim);

  export const reset = (value: string, type: string) => {
    if (model.isDisposed()) return;
    model?.setValue(value || '');
    monaco.editor?.setModelLanguage(model, type);
  };

  onMount(async () => {
    model = monaco.editor.createModel(content, type);
    editor = monaco.editor.create(editorContainer, {
      model,
      minimap: { enabled: false },
      dimension: { height, width },
    });

    editor.addCommand(
      monaco.KeyMod.CtrlCmd | monaco.KeyCode.Enter,
      () => dispatch('run'),
    );

    model.onDidChangeContent((e) =>
      dispatch('update', model.getValue()),
    );
  });

  onDestroy(() => {
    editor?.dispose();
    model?.dispose();
  });
</script>

<div bind:this={editorContainer}></div>
