<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div class="cols-span-1 flex flex-col gap-6 w-full">
  {@render children()}
</div>
