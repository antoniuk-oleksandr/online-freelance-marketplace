<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div
  class="p-6 grid bg-light-palette-background-block dark:bg-dark-palette-background-block gap-6 grid-cols-order-overview-receipt border-x-2 border-light-palette-action-hover dark:border-dark-palette-action-hover"
>
  {@render children()}
</div>
