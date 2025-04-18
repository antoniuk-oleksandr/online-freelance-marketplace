<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div
  class="md:hidden flex flex-col gap-1 bg-light-palette-action-hover dark:bg-dark-palette-action-hover p-6 -mx-6 mt-6"
>
  {@render children()}
</div>
