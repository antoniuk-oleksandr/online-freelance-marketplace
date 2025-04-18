<script lang="ts">
  import type { ServiceQA } from '@/types/ServiceQA'
  import OrderByIdRequirementsTabQuestion from '../OrderByIdRequirementsTabQuestion/OrderByIdRequirementsTabQuestion.svelte'
  import OrderByIdRequirementsTabQaLayout from './OrderByIdRequirementsTabQALayout.svelte'
    import OrderByIdRequirementsTabAnswer from '../OrderByIdRequirementsTabAnswer/OrderByIdRequirementsTabAnswer.svelte'

  type OrderByIdRequirementsTabQAProps = ServiceQA & {
    index: number
  }

  const props: OrderByIdRequirementsTabQAProps = $props()
</script>

<OrderByIdRequirementsTabQaLayout>
  <OrderByIdRequirementsTabQuestion {...props} />
  <OrderByIdRequirementsTabAnswer {...props} />
</OrderByIdRequirementsTabQaLayout>
