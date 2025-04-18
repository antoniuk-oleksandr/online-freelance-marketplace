<script lang="ts">
  import Number from '@/common-components/Number/Number.svelte'
  import OrderByIdRequirementsTabQuestionLayout from './OrderByIdRequirementsTabQuestionLayout.svelte'

  type OrderByIdRequirementsTabQuestionProps = {
    index: number
    question: string
  }

  const { index, question }: OrderByIdRequirementsTabQuestionProps = $props()
</script>

<OrderByIdRequirementsTabQuestionLayout>
  <Number value={index + 1} />
  <span class="font-semibold text-light-palette-text-primary dark:text-dark-palette-text-primary"
    >{question}
  </span>
</OrderByIdRequirementsTabQuestionLayout>
