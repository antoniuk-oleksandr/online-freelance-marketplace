<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div class="flex flex-col flex-1 min-h-0">
  {@render children()}
</div>
