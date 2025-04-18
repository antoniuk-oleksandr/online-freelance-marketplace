<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'
  import { fade } from 'svelte/transition'

  type OrderSubmitRequirementsFormAttachFileCloseButtonLayoutProps = LayoutProps & {
    index: number
    removeFile: (index: number) => void
  }

  const {
    children,
    removeFile,
    index,
  }: OrderSubmitRequirementsFormAttachFileCloseButtonLayoutProps = $props()
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<div
  onclick={() => removeFile(index)}
  class="right-0 top-0 absolute -mr-[15px] -mt-[15px] bg-light-palette-background-block dark:bg-dark-palette-background-block rounded-full p-2 text-red-500 hover:bg-red-100 dark:hover:bg-red-500/10 hover:text-red-600 transition-all ease-in-out duration-200 cursor-pointer focus:outline-none focus:ring-2 focus:ring-red-500"
  transition:fade={{ duration: 200 }}
>
  {@render children()}
</div>
