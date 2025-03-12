<script lang="ts">
  import Button from '@/common-components/Button/Button.svelte'
  import DividerElement from '@/common-components/DividerElement/DividerElement.svelte'
  import Title from '@/common-components/Title/Title.svelte'
  import type { OrderSubmitRequirementsData } from '@/types/OrderSubmitRequirementsData'
  import type { OrderSubmitRequirementsFormErrors } from '@/types/OrderSubmitRequirementsFormErrors'
  import OrderSubmitRequirementsFormQuestion from '../OrderSubmitRequirementsFormQuestion/OrderSubmitRequirementsFormQuestion.svelte'
  import OrderSubmitRequirementsFormQuestionList from '../OrderSubmitRequirementsFormQuestionList/OrderSubmitRequirementsFormQuestionList.svelte'
  import OrderSubmitRequirementsFormBodyLayout from './OrderSubmitRequirementsFormBodyLayout.svelte'
  import type { OrderSubmitRequirementsFormData } from '@/types/OrderSubmitRequirementsFormData'
  import OrderSubmitRequirementsFormAttachFilesBlock from '../OrderSubmitRequirementsFormAttachFilesBlock/OrderSubmitRequirementsFormAttachFilesBlock.svelte'

  type OrderSubmitRequirementsFormBodyProps = {
    orderData: OrderSubmitRequirementsData
    orderFormErrors: undefined | OrderSubmitRequirementsFormErrors
    orderFormData: undefined | OrderSubmitRequirementsFormData
    wasSubmitted: boolean
  }

  const {
    orderData,
    orderFormErrors,
    wasSubmitted,
    orderFormData,
  }: OrderSubmitRequirementsFormBodyProps = $props()

  const title = 'The freelancer needs the following information to start working on your order:'
  const messageContent =
    'Provide details about your request, including requirements and preferences for the best outcome.'
</script>

{#if orderFormData}
  <OrderSubmitRequirementsFormBodyLayout>
    <Title capitalize={false} text={title} />
    <OrderSubmitRequirementsFormQuestionList
      {wasSubmitted}
      {orderFormErrors}
      {orderFormData}
      freelanceQuestions={orderData.freelanceQuestions}
    />
    <DividerElement />
    <OrderSubmitRequirementsFormQuestion
      value={orderFormData.customerMessage}
      {wasSubmitted}
      placeholder="Message"
      error={orderFormErrors?.customerMessage}
      content={messageContent}
      id="customerMessage"
    />
    <DividerElement />
    <OrderSubmitRequirementsFormAttachFilesBlock />
    <Button type="submit" styles="self-end w-full md:!w-48">Start Order</Button>
  </OrderSubmitRequirementsFormBodyLayout>
{/if}
