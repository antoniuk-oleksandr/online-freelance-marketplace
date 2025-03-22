<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  type DropdownLayoutProps = LayoutProps & {
    layoutStyles?: string
  }

  const { children, layoutStyles }: DropdownLayoutProps = $props()
</script>

<div class="{layoutStyles} relative flex flex-col">
  {@render children()}
</div>
