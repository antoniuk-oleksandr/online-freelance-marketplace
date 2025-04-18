<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div
  class="flex flex-col gap-6 rounded-t-md bg-light-palette-action-hover dark:bg-dark-palette-action-hover p-6"
>
  {@render children()}
</div>
