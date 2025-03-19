<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div class="size-fit ">
  {@render children()}
</div>
