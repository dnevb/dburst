import * as monaco from 'monaco-editor';

import editorWorker from 'monaco-editor/esm/vs/editor/editor.worker?worker';

self.MonacoEnvironment = {
  getWorker: function (_: string, label: string) {
    switch (label) {
      default:
        return new editorWorker();
    }
  },
};

monaco.editor.defineTheme('dark', {
  base: 'vs-dark',
  inherit: true,
  rules: [],
  colors: {
    'editor.background': '#0c0a09',
  },
});
monaco.editor.defineTheme('light', {
  base: 'vs',
  inherit: true,
  rules: [],
  colors: {
    'editor.background': '#ffffff',
  },
});
monaco.editor.setTheme(localStorage.getItem('theme') || 'dark');

export default monaco;
