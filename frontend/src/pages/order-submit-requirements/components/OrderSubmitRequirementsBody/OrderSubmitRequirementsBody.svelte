<script lang="ts">
  import OrderSubmitRequirementsBodyLayout from '@/pages/order-submit-requirements/components/OrderSubmitRequirementsBody/OrderSubmitRequirementsBodyLayout.svelte'
  import type { OrderSubmitRequirementsData } from '@/types/OrderSubmitRequirementsData.ts'
  import OrderSubmitRequirementsInfo from '@/pages/order-submit-requirements/components/OrderSubmitRequirementsInfo/OrderSubmitRequirementsInfo.svelte'
  import OrderSubmitRequirementsForm from '../OrderSubmitRequirementsForm/OrderSubmitRequirementsForm.svelte'

  type OrderSubmitRequirementsBodyProps = {
    orderData: OrderSubmitRequirementsData
    orderId: string
  }

  const props: OrderSubmitRequirementsBodyProps = $props()
</script>

<OrderSubmitRequirementsBodyLayout>
  <OrderSubmitRequirementsForm {...props} />
  <OrderSubmitRequirementsInfo {...props} />
</OrderSubmitRequirementsBodyLayout>
