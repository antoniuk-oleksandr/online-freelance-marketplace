<script lang="ts">
  import PaperElement from '@/common-components/PaperElement/PaperElement.svelte'
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<PaperElement
  styles="h-fit sticky lg:top-22 lg:mt-19 !ring-0 lg:dark:!ring-1 !p-0 lg:!p-6 !shadow-none !bg-transparent lg:!bg-light-palette-background-block lg:dark:!bg-dark-palette-background-block  lg:!shadow-md dark:!shadow-none gap-3 flex flex-col flex-grow text-base text-light-palette-text-secondary dark:text-dark-palette-text-secondary"
>
  {@render children()}
</PaperElement>
