<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div class="bg-light-palette-background-block dark:bg-dark-palette-background-block rounded-full">
  {@render children()}
</div>
