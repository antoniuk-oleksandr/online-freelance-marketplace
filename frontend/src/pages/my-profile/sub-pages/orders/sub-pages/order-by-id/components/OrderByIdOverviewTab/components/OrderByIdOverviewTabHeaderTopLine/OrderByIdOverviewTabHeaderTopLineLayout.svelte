<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div class="flex col-span-3 gap-3 items-center">
  {@render children()}
</div>
