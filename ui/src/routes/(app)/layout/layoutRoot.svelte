<script>
  import colorMode from 'stores/colorMode';
  import LayoutSidebar from './layoutSidebar.svelte';
</script>

<div
  class="h-screen w-full overflow-hidden flex"
  data-color-mode={$colorMode}
>
  <LayoutSidebar />
  <div class="flex-1 min-h-full max-h-full">
    <slot />
  </div>
</div>
