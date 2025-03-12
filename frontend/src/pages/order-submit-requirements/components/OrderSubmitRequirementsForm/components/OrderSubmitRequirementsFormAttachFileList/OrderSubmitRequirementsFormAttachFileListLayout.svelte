<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div class="grid grid-cols-4 gap-3">
  {@render children()}
</div>
