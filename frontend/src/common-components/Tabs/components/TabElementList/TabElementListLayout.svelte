<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div class="flex gap-6 text-lg font-medium justify-center md:justify-normal">
  {@render children()}
</div>
