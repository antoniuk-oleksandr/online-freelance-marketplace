<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div class="lg:hidden !h-13.5 bg-light-palette-background-block dark:bg-dark-palette-background-block w-full p-3 text-lg font-semibold border-b border-light-palette-divider dark:border-dark-palette-divider flex justify-between items-center">
  {@render children()}
</div>
