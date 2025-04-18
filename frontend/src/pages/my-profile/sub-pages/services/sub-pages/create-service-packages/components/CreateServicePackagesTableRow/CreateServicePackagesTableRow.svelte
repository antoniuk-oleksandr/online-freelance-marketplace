<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div
  class="grid grid-cols-create-service-packages-table border-b border-light-palette-divider dark:border-dark-palette-divider"
>
  {@render children()}
</div>
