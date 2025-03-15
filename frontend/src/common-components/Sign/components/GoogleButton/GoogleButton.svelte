<script lang="ts">
  import Icon from '@iconify/svelte'
  import { handleGoogleButtonClick } from '@/common-components/Sign/handlers'
  import Spinner from '@/common-components/Spinner/Spinner.svelte'
  import type { SignProps } from '@/types/SignProps.ts'

  type GoogleButtonProps = SignProps & {
    keepSignedIn: boolean
  }

  const { googleButtonText, keepSignedIn }: GoogleButtonProps = $props()

  const clientId = import.meta.env.VITE_GOOGLE_CLIENT_ID

  let loading = $state(false)
  const setLoading = (value: boolean) => (loading = value)
</script>

{#if googleButtonText !== undefined}
  <button
    disabled={loading}
    onclick={() => handleGoogleButtonClick(setLoading, clientId, keepSignedIn)}
    type="button"
    class="{loading
      ? 'opacity-70'
      : 'active:scale-95 hover:bg-light-palette-action-hover dark:hover:bg-dark-palette-action-hover'} flex !h-12 bg-light-palette-background-block dark:bg-dark-palette-background-block font-semibold duration-200 ease-out items-center justify-center border border-light-palette-divider dark:border-dark-palette-divider rounded-md gap-x-3 w-full"
  >
    {#if loading}
      <Spinner
        size="size-8"
        color="border-l-light-palette-text-primary dark:border-l-dark-palette-text-primary"
      />
    {:else}
      <Icon icon="flat-color-icons:google" width="24" height="24" />
      <span>{googleButtonText}</span>
    {/if}
  </button>
{/if}
