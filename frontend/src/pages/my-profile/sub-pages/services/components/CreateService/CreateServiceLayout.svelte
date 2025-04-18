<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div class="flex min-h-app flex-col gap-6">
  {@render children()}
</div>
