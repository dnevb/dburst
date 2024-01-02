import monaco from 'provider/monaco';
import { writable } from 'svelte/store';

const colorMode = writable(localStorage.getItem('theme') || 'dark');

export const toggleColorMode = () =>
  colorMode.update((v) => {
    const mode = v == 'dark' ? 'light' : 'dark';
    localStorage.setItem('theme', mode);
    monaco.editor.setTheme(mode);
    return mode;
  });

export default colorMode;
