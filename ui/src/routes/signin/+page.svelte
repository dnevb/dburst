<script lang="ts">
  import SigninForm from './SigninForm.svelte';
</script>

<section id="login-page">
  <div
    class="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0"
  >
    <div
      class="flex items-center mb-6 text-2xl font-semibold text-gray-900"
    >
      <img class="h-16 mr-2" src="/logo-banner.svg" alt="banner" />
    </div>
    <div class="w-full rounded-2xl md:mt-0 sm:max-w-md xl:p-0">
      <div class="p-6 space-y-4 md:space-y-6 sm:p-8">
        <SigninForm />
      </div>
    </div>
  </div>
</section>
