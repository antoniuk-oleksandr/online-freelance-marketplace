<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div class="flex rounded-t-xl gap-2 flex-wrap p-6 bg-light-palette-action-hover dark:bg-dark-palette-action-hover">
  {@render children()}
</div>
