<script lang="ts">
  export let type: string;
</script>

{#if type.includes('int') || type.includes('number') || type.includes('numeric')}
  <i class="i-carbon:hashtag" />
{:else if type.includes('text') || type.includes('varchar')}
  <i class="i-carbon:letter-aa" />
{:else if type.includes('date') || type.includes('time')}
  <i class="i-carbon:time" />
{:else if type == 'bool'}
  <i class="i-carbon:boolean" />
{:else}
  <i class="i-carbon:unknown" />
{/if}
