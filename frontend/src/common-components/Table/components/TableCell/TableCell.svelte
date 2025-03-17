<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  type OrdersTableCellProps = LayoutProps & {
    styles?: string
    label?: string
  }

  const { children, styles, label }: OrdersTableCellProps = $props()
</script>

<div class="flex p-0 lg:p-3 h-fit {styles}">
  {#if label}
    <span class="lg:hidden font-semibold">{label}</span>
  {/if}
  {@render children()}
</div>
