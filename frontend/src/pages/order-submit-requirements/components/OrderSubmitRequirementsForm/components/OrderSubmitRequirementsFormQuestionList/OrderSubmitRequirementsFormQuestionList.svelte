<script lang="ts">
  import type { ServiceQuestion } from '@/types/ServiceQuestion'
  import OrderSubmitRequirementsFormQuestionListLayout from './OrderSubmitRequirementsFormQuestionListLayout.svelte'
  import OrderSubmitRequirementsFormQuestion from '../OrderSubmitRequirementsFormQuestion/OrderSubmitRequirementsFormQuestion.svelte'
  import DividerElement from '@/common-components/DividerElement/DividerElement.svelte'
  import type { OrderSubmitRequirementsFormErrors } from '@/types/OrderSubmitRequirementsFormErrors'
  import type { OrderSubmitRequirementsFormData } from '@/types/OrderSubmitRequirementsFormData'

  type OrderSubmitRequirementsFormQuestionListProps = {
    freelanceQuestions: ServiceQuestion[]
    orderFormErrors: undefined | OrderSubmitRequirementsFormErrors
    orderFormData: undefined | OrderSubmitRequirementsFormData
    wasSubmitted: boolean
  }

  const {
    freelanceQuestions,
    orderFormErrors,
    wasSubmitted,
    orderFormData,
  }: OrderSubmitRequirementsFormQuestionListProps = $props()
</script>

{#if orderFormData}
  <OrderSubmitRequirementsFormQuestionListLayout>
    {#each freelanceQuestions as item, index}
      <OrderSubmitRequirementsFormQuestion
        value={orderFormData.answers[index].content}
        {wasSubmitted}
        placeholder="Answer"
        error={orderFormErrors?.answers[index]}
        content={item.content}
        id={`answers.${index}.content`}
        {index}
      />
      {#if index !== freelanceQuestions.length - 1}
        <DividerElement />
      {/if}
    {/each}
  </OrderSubmitRequirementsFormQuestionListLayout>
{/if}
