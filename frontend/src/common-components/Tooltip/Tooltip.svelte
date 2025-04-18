<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  type ToolTipProps = LayoutProps & {
    text: string
  }

  const { children, text }: ToolTipProps = $props()
</script>

<div class="flex relative group">
  {@render children()}
  <div
    class="absolute whitespace-nowrap py-0.5 left-0 bottom-full bg-neutral-800 text-white px-2 text-sm rounded-md mb-1 opacity-0 group-hover:opacity-100 transition-opacity"
  >
    {text}
  </div>
</div>
