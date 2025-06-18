<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div class="grid grid-cols-1 md:grid-cols-2 gap-4 flex-1 min-h-0 overflow-auto">
  {@render children()}
</div>
