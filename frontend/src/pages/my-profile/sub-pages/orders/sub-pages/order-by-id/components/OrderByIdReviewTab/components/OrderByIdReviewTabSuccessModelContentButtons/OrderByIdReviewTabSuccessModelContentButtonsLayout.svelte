<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div class="flex w-full gap-6 flex-col md:flex-row h-30 md:h-fit">
  {@render children()}
</div>
