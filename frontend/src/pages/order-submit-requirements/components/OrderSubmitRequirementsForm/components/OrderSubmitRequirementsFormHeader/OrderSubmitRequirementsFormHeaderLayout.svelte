<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div class="p-6 -mx-6 mb-6 md:m-0 rounded-none md:rounded-t-xl bg-light-palette-action-hover dark:bg-dark-palette-action-hover">
  {@render children()}
</div>
