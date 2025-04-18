<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div
  class="flex text-lg font-semibold items-center gap-1 text-light-palette-text-primary dark:text-dark-palette-text-primary"
>
  {@render children()}
</div>
