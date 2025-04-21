<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div class="flex col-span-2 items-end justify-end">
  {@render children()}
</div>
