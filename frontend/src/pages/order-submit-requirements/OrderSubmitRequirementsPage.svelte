<script lang="ts">
  import { onMount } from 'svelte'
  import type { StepperItem } from '@/types/StepperItem.ts'
  import OrderSubmitRequirementsPageLayout from '@/pages/order-submit-requirements/OrderSubmitRequirementsPageLayout.svelte'
  import Stepper from '@/common-components/Stepper/Stepper.svelte'
  import type { OrderSubmitRequirementsData } from '@/types/OrderSubmitRequirementsData.ts'
  import {
    getSubmitRequirementsSteps,
    makeOrderServiceQuestionsRequest,
  } from '@/pages/order-submit-requirements/helpers'
  import OrderSubmitRequirementsBody from '@/pages/order-submit-requirements/components/OrderSubmitRequirementsBody/OrderSubmitRequirementsBody.svelte'
  import SuccessOrderMessage from '@/pages/order-submit-requirements/components/SuccessOrderMessage/SuccessOrderMessage.svelte'

  type OrderSubmitRequirementsPageParams = {
    orderId: string
  }

  const { orderId }: OrderSubmitRequirementsPageParams = $props()

  let orderData = $state<OrderSubmitRequirementsData | undefined>()
  const setOrderData = (value: OrderSubmitRequirementsData) => (orderData = value)

  let steps: StepperItem[] | undefined = $derived(getSubmitRequirementsSteps(orderData))

  onMount(() => makeOrderServiceQuestionsRequest(orderId, setOrderData))
</script>

{#if steps && orderData}
  <OrderSubmitRequirementsPageLayout>
    <Stepper styles="col-span-2 capitalize flex h-fit justify-center" activeStepIndex={2} {steps} />
    <SuccessOrderMessage />
    <OrderSubmitRequirementsBody {orderId} {orderData} />
  </OrderSubmitRequirementsPageLayout>
{/if}
