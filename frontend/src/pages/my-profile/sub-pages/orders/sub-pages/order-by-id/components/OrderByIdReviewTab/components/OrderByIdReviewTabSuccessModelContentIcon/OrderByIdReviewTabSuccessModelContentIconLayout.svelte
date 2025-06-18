<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div
  class=" p-4 bg-green-100 text-green-700 border-green-300 dark:bg-green-900 dark:text-green-200 size-fit rounded-full"
>
  {@render children()}
</div>
