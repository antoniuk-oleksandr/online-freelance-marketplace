<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div class="bg-red-600 dark:bg-red-900 h-22 flex items-center px-6 md:rounded-t-xl w-full">
  {@render children()}
</div>
