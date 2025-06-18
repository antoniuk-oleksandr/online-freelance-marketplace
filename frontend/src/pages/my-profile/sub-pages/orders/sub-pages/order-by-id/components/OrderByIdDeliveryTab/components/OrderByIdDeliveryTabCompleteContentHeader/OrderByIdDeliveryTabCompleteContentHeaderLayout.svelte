<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div class="bg-green-600 dark:bg-green-900 px-6 h-22 md:rounded-t-xl flex items-center w-full">
  {@render children()}
</div>
