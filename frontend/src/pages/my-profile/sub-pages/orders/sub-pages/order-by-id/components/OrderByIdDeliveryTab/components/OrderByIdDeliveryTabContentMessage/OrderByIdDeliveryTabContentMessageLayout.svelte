<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div
  class="bg-light-palette-action-hover dark:bg-dark-palette-action-hover rounded-lg p-6 ring-1 ring-light-palette-divider dark:ring-dark-palette-divider"
>
  {@render children()}
</div>
