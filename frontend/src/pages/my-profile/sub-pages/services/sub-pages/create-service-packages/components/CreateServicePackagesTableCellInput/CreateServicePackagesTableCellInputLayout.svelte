<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div class="flex items-center pr-3 focus-within:ring-1 ring-cyan-500 ease-out duration-200">
  {@render children()}
</div>
