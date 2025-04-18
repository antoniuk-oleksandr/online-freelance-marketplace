<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div class="hidden md:block mt-6 rounded-md ring-1 ring-light-palette-divider dark:ring-dark-palette-divider">
  {@render children()}
</div>
