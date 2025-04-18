<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div class="text-light-palette-text-primary dark:text-dark-palette-text-primary grid border-t text-base border-l border-light-palette-divider dark:border-dark-palette-divider">
  {@render children()}
</div>
