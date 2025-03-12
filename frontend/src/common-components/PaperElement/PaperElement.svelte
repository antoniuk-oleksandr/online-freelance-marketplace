<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps.ts'

  type PaperElementProps = LayoutProps & {
    styles?: string
  }

  const { children, styles }: PaperElementProps = $props()
</script>

<div
  class={`${styles} p-6 shadow-md dark:shadow-none rounded-xl bg-light-palette-background-block dark:bg-dark-palette-background-block dark:ring-1 ring-light-palette-divider dark:ring-dark-palette-divider`}
>
  {@render children()}
</div>
