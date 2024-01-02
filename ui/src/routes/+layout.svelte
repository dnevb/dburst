<script lang="ts">
  import '@unocss/reset/tailwind-compat.css';
  import 'uno.css';
  import '../styles/root.scss';
  import colorMode from 'stores/colorMode';

  $: {
    if ($colorMode == 'light') document.body.classList.remove('dark');
    else document.body.classList.add('dark');
  }
</script>

<slot />
