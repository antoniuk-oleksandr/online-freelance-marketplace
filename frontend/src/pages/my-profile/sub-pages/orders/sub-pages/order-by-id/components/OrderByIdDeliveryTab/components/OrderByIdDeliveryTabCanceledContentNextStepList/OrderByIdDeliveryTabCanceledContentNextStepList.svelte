<script lang="ts">
  import type { OrderByIdDeliveryTabCanceledContentStepItem } from '@/types/OrderByIdDeliveryTabCanceledContentStepItem'
  import OrderByIdDeliveryTabCanceledContentNextStepListLayout from './OrderByIdDeliveryTabCanceledContentNextStepListLayout.svelte'
  import OrderByIdDeliveryTabCanceledContentStep from '../OrderByIdDeliveryTabCanceledContentStep/OrderByIdDeliveryTabCanceledContentStep.svelte'

  const steps: OrderByIdDeliveryTabCanceledContentStepItem[] = [
    {
      title: 'Find Another Freelancer',
      description: 'Browse similar services to find the perfect freelancer for your project.',
      icon: 'heroicons:arrow-path',
    },
    {
      icon: 'heroicons:question-mark-circle',
      title: 'Need Help?',
      description: 'Our support team is available 24/7 to assist with any questions.',
    },
  ]
</script>

<OrderByIdDeliveryTabCanceledContentNextStepListLayout>
  {#each steps as item}
    <OrderByIdDeliveryTabCanceledContentStep {...item} />
  {/each}
</OrderByIdDeliveryTabCanceledContentNextStepListLayout>
