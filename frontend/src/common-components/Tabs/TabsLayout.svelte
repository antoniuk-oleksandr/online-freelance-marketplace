<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div class="flex flex-col w-full gap-6">
  {@render children()}
</div>
