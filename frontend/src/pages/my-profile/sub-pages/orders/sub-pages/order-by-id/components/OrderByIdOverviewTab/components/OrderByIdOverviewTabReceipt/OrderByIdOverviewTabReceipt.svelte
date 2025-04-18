<script lang="ts">
  import type { OrderByIdOverviewTabData } from '@/types/OrderByIdOverviewTabData'
  import OrderByIdOverviewTabReceiptHeader from '../OrderByIdOverviewTabReceiptHeader/OrderByIdOverviewTabReceiptHeader.svelte'
  import OrderByIdOverviewTabReceiptLayout from './OrderByIdOverviewTabReceiptLayout.svelte'
  import OrderByIdOverviewTabReceiptBody from '../OrderByIdOverviewTabReceiptBody/OrderByIdOverviewTabReceiptBody.svelte'
  import OrderByIdOverviewTabReceiptFooter from '../OrderByIdOverviewTabReceiptFooter/OrderByIdOverviewTabReceiptFooter.svelte'

  type OrderByIdOverviewTabReceiptProps = {
    order: OrderByIdOverviewTabData
  }

  const props: OrderByIdOverviewTabReceiptProps = $props()
</script>

<OrderByIdOverviewTabReceiptLayout>
  <OrderByIdOverviewTabReceiptHeader {...props} />
  <OrderByIdOverviewTabReceiptBody {...props} />
  <OrderByIdOverviewTabReceiptFooter {...props} />
</OrderByIdOverviewTabReceiptLayout>
