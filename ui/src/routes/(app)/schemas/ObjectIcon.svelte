<script lang="ts">
  export let type: string;
</script>

{#if type == 'table'}
  <i class="i-carbon:data-table"></i>
{:else if type == 'view'}
  <i class="i-carbon:view"></i>
{:else if type == 'materialized_view'}
  <i class="i-carbon:cube"></i>
{/if}
