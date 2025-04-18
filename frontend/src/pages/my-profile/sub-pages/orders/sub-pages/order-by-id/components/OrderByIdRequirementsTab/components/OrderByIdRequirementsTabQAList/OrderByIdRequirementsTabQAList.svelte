<script lang="ts">
  import type { ServiceQA } from '@/types/ServiceQA'
  import OrderByIdRequirementsTabQa from '../OrderByIdRequirementsTabQA/OrderByIdRequirementsTabQA.svelte'

  type orderByIdRequirementsTabQAListProps = {
    serviceQAs: ServiceQA[]
  }

  const { serviceQAs }: orderByIdRequirementsTabQAListProps = $props()
</script>

{#each serviceQAs as { question, answer }, index}
  <OrderByIdRequirementsTabQa {question} {answer} {index} />
{/each}
