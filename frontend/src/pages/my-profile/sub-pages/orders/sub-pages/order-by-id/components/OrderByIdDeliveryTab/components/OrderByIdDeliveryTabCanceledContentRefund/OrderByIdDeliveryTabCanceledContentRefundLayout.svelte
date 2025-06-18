<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div
  class="bg-blue-100 dark:bg-blue-900 rounded-lg p-6 ring-1 border-blue-300 dark:border-blue-600"
>
  <div class="flex items-start gap-3">
    {@render children()}
  </div>
</div>
