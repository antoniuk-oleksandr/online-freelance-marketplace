<script lang="ts">
  import type { ServiceTableData } from '@/types/ServiceTableData'
  import ServicesTableLayout from './ServicesTableLayout.svelte'
  import ServicesTableBody from './ServicesTableBody/ServicesTableBody.svelte'
  import ServicesTableHeader from './ServicesTableHeader/ServicesTableHeader.svelte'

  type OrdersTableProps = {
    services: ServiceTableData[]
  }

  const props: OrdersTableProps = $props()
</script>

<ServicesTableLayout>
  <ServicesTableHeader />
  <ServicesTableBody {...props} />
</ServicesTableLayout>
