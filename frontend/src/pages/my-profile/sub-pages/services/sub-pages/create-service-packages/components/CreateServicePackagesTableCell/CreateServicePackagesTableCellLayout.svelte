<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div class="border-r border-light-palette-divider dark:border-dark-palette-divider text-base">
  {@render children()}
</div>
