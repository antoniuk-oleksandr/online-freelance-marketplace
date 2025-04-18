<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  type OrderByIdOverviewTabReceiptRowLayoutProps = LayoutProps & {
    styles?: string
  }

  const { children, styles }: OrderByIdOverviewTabReceiptRowLayoutProps = $props()
</script>

<div class="{styles} flex justify-between gap-3">
  {@render children()}
</div>
