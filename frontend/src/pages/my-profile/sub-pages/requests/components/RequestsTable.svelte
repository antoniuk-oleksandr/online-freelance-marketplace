<script lang="ts">
  import RequestsTableHeader from './RequestsTableHeader/RequestsTableHeader.svelte'
  import type { RequestTableData } from '@/types/RequestTableData'
  import RequestsTableBody from './RequestsTableBody/RequestsTableBody.svelte'
  import RequestsTableLayout from './RequestsTableLayout.svelte'

  type OrdersTableProps = {
    requests: RequestTableData[]
  }

  const props: OrdersTableProps = $props()
</script>

<RequestsTableLayout>
  <RequestsTableHeader />
  <RequestsTableBody {...props} />
</RequestsTableLayout>
