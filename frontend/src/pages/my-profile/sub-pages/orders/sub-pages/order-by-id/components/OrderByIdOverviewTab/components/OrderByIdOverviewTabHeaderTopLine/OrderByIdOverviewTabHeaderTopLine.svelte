<script lang="ts">
  import { formatOrderDate } from '@/pages/order-submit-requirements/helpers'
  import OrderByIdOverviewTabHeaderTopLineLayout from './OrderByIdOverviewTabHeaderTopLineLayout.svelte'
  import Icon from '@iconify/svelte'

  type OrderByIdOverviewTabHeaderTopLineProps = {
    createdAt: number
  }

  const { createdAt }: OrderByIdOverviewTabHeaderTopLineProps = $props()
</script>

<OrderByIdOverviewTabHeaderTopLineLayout>
  <Icon width="24" height="24" icon="hugeicons:invoice-02" />
  <span class="text-light-palette-text-primary dark:text-dark-palette-text-primary"
    >Order receipt</span
  >
  <span class="italic">{formatOrderDate(createdAt)}</span>
</OrderByIdOverviewTabHeaderTopLineLayout>
