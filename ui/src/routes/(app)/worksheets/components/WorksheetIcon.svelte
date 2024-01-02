<script lang="ts">
  export let type: string;
</script>

<span>
  {#if type == 'sql'}
    <i class="i-carbon:sql" />
  {:else if type == 'python'}
    <i class="i-carbon:logo-python" />
  {:else}
    <i class="i-carbon:cube" />
  {/if}
</span>
