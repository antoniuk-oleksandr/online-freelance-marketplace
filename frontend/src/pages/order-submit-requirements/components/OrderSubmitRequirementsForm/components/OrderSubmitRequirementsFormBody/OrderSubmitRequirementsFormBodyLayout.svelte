<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div class="p-0 md:p-6 flex flex-col gap-6">
  {@render children()}
</div>
