<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div class="relative flex items-center justify-center py-3">
  {@render children()}
</div>
