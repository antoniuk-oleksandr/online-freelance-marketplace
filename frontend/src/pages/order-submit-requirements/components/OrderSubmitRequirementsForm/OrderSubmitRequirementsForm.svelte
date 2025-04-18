<script lang="ts">
  import type { OrderSubmitRequirementsData } from '@/types/OrderSubmitRequirementsData'
  import OrderSubmitRequirementsFormLayout from './OrderSubmitRequirementsFormLayout.svelte'
  import OrderSubmitRequirementsFormHeader from './components/OrderSubmitRequirementsFormHeader/OrderSubmitRequirementsFormHeader.svelte'
  import OrderSubmitRequirementsFormBody from './components/OrderSubmitRequirementsFormBody/OrderSubmitRequirementsFormBody.svelte'
  import type { OrderSubmitRequirementsFormErrors } from '@/types/OrderSubmitRequirementsFormErrors'
  import type { OrderSubmitRequirementsFormData } from '@/types/OrderSubmitRequirementsFormData'

  type OrderSubmitRequirementsFormProps = {
    orderData: OrderSubmitRequirementsData
    orderId: string
  }

  const { orderData, orderId }: OrderSubmitRequirementsFormProps = $props()

  let wasSubmitted = $state(false)
  let orderFormErrors = $state<undefined | OrderSubmitRequirementsFormErrors>()
  let orderFormData = $state<undefined | OrderSubmitRequirementsFormData>()
</script>

<OrderSubmitRequirementsFormLayout
  {orderId}
  bind:orderFormErrors
  bind:wasSubmitted
  bind:orderFormData
  freelanceQuestions={orderData.freelanceQuestions}
>
  <OrderSubmitRequirementsFormHeader />
  <OrderSubmitRequirementsFormBody {orderFormData} {wasSubmitted} {orderFormErrors} {orderData} />
</OrderSubmitRequirementsFormLayout>
