<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div class="uppercase font-semibold hidden md:block md:text-end">
  {@render children()}
</div>
