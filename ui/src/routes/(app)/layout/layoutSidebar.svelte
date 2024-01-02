<script lang="ts">
  import colorMode, { toggleColorMode } from 'stores/colorMode';
  import Button from 'ui/button/button.svelte';
  import navlinks from './navlinks.json';
  import { page } from '$app/stores';
  import { invalidateAll } from '$app/navigation';

  const logout = () => {
    localStorage.removeItem('api_token');
    invalidateAll();
  };
</script>

<aside>
  <nav>
    <a href="/" title="Home page">
      <img
        class="w-auto h-12"
        src="/logo-icon.svg"
        alt="sidebar logo"
      />
    </a>

    {#each navlinks as element}
      <Button
        variant={$page.url.pathname.startsWith(element.path)
          ? 'default'
          : 'ghost'}
        size="icon"
        href={element.path}
        title={element.title}
      >
        <i class={element.icon}></i>
      </Button>
    {/each}
  </nav>

  <div class="flex flex-col space-y-4">
    <Button variant="ghost" size="icon" on:click={toggleColorMode}>
      {#if $colorMode == 'dark'}
        <i class="i-carbon:sun"></i>
      {:else}
        <i class="i-carbon:moon"></i>
      {/if}
    </Button>
    <Button variant="ghost" size="icon" on:click={logout}>
      <i class="i-carbon:logout"></i>
    </Button>
  </div>
</aside>

<style lang="scss">
  aside {
    --at: flex flex-col items-center w-16 h-screen py-8
      overflow-y-auto border-r;
  }
  nav {
    --at: flex flex-col flex-1 space-y-4 items-center;
  }
</style>
