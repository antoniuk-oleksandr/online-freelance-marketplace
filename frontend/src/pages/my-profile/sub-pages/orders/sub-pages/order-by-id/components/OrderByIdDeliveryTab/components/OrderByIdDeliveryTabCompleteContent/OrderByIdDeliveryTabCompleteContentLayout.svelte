<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div class="h-full w-full flex flex-col">
  {@render children()}
</div>
