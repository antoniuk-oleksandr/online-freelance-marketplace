<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div class="min-h-app gap-6 grid grid-cols-1 lg:grid-cols-search-page">
  {@render children()}
</div>
