<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps.ts'

  type NoContentMessageLayoutProps = LayoutProps & {
    styles?: string
  }

  const { children, styles }: NoContentMessageLayoutProps = $props()
</script>

<div
  class="{styles} text-center text-base text-light-palette-text-secondary dark:text-dark-palette-text-secondary"
>
  {@render children()}
</div>
