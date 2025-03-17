<script lang="ts">
  import type { OrderTableData } from '@/types/OrderTableData'
  import OrdersTableHeader from './components/OrdersTableHeader/OrdersTableHeader.svelte'
  import OrdersTableBody from './components/OrdersTableBody/OrdersTableBody.svelte'
  import OrdersTableLayout from './OrdersTableLayout.svelte'

  type OrdersTableProps = {
    orders: OrderTableData[]
  }

  const props: OrdersTableProps = $props()
</script>

<OrdersTableLayout>
  <OrdersTableHeader />
  <OrdersTableBody {...props} />
</OrdersTableLayout>
