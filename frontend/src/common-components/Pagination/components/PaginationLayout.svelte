<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div class="flex items-center justify-center gap-3 text-base">
  {@render children()}
</div>
