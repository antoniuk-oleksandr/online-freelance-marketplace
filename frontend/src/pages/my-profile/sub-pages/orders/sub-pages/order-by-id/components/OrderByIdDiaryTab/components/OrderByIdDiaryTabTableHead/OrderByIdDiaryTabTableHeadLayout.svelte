<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div class="grid grid-cols-7">
  {@render children()}
</div>
