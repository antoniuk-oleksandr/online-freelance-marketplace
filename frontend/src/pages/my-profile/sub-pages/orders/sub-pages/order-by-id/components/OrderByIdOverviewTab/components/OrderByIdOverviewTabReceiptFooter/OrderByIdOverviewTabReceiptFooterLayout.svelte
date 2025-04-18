<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div
  class="p-6 rounded-b-md font-semibold bg-light-palette-action-hover dark:bg-dark-palette-action-hover"
>
  {@render children()}
</div>
