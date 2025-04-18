<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div class="p-2  border border-light-palette-divider dark:border-dark-palette-divider">
  <div class="aspect-square flex flex-col justify-between">
    {@render children()}
  </div>
</div>
