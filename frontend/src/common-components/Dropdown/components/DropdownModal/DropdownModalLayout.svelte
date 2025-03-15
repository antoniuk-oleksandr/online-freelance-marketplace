<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  const { children }: LayoutProps = $props()
</script>

<div
  class="p-6 lg:p-0 flex-col fixed lg:static top-0 left-0 lg:top-auto lg:left-auto size-full lg:size-auto bg-black lg:bg-transparent bg-opacity-50 flex justify-center items-center"
>
  {@render children()}
</div>
