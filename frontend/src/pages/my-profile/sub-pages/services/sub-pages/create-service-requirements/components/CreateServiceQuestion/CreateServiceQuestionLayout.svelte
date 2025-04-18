<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div class="group p-6 bg-light-palette-action-hover dark:bg-dark-palette-action-hover relative flex items-start rounded-md ring-1 ring-light-palette-divider dark:ring-dark-palette-divider">
  {@render children()}  
</div>
