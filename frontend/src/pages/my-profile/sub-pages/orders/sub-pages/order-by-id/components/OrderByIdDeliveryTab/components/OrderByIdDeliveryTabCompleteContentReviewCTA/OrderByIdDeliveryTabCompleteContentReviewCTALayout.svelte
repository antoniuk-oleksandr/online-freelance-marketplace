<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div class="dark:bg-green-900 dark:border-green-700 bg-green-100 rounded-lg p-6 border border-green-300">
  {@render children()}
</div>
