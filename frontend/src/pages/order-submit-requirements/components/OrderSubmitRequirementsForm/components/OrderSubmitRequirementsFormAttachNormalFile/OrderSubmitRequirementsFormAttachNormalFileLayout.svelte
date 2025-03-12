<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div
  class="w-full h-full flex flex-col gap-1 items-center justify-center bg-light-palette-action-hover dark:bg-dark-palette-action-hover rounded-xl p-3"
>
  {@render children()}
</div>
