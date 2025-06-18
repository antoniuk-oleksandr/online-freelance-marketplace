<script lang="ts">
  import Icon from '@iconify/svelte'
  import OrderByIdDeliveryTabCanceledContentStepLayout from './OrderByIdDeliveryTabCanceledContentStepLayout.svelte'
  import type { OrderByIdDeliveryTabCanceledContentStepItem } from '@/types/OrderByIdDeliveryTabCanceledContentStepItem'

  type OrderByIdDeliveryTabCanceledContentStepProps = OrderByIdDeliveryTabCanceledContentStepItem

  const { title, description, icon }: OrderByIdDeliveryTabCanceledContentStepProps = $props()
</script>

<OrderByIdDeliveryTabCanceledContentStepLayout>
  <div class="bg-light-pa dark:bg-gray-700 rounded-full p-2 mt-1">
    <Icon {icon} class="w-5 h-5 text-gray-700 dark:text-gray-300" />
  </div>
  <div class="flex flex-col">
    <h4 class="font-medium text-light-palette-text-primary dark:text-dark-palette-text-primary">
      {title}
    </h4>
    <p class="text-light-palette-text-secondary dark:text-dark-palette-text-secondary">
      {description}
    </p>
  </div>
</OrderByIdDeliveryTabCanceledContentStepLayout>
