<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div
  class="size-full flex items-center justify-center flex-col gap-1 text-light-palette-text-secondary dark:text-dark-palette-text-secondary"
>
  {@render children()}
</div>
