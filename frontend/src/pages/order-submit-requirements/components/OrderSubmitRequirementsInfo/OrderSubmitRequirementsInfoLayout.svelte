<script lang="ts">
  import PaperElement from '@/common-components/PaperElement/PaperElement.svelte'
  import type { LayoutProps } from '@/types/LayoutProps.ts'

  const { children }: LayoutProps = $props()
</script>

<PaperElement styles="h-fit overflow-hidden !p-0 md:!p-6 shadow-none md:shadow-md dark:shadow-none !ring-0 md:!ring-1 !bg-transparent md:!bg-light-palette-background-block md:dark:!bg-dark-palette-background-block">
  {@render children()}
</PaperElement>
