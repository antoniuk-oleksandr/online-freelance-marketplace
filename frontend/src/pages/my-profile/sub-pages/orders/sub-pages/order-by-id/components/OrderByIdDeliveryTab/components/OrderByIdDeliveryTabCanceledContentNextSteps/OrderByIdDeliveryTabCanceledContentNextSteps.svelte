<script lang="ts">
  import Title from '@/common-components/Title/Title.svelte'
  import OrderByIdDeliveryTabCanceledContentNextStepsLayout from './OrderByIdDeliveryTabCanceledContentNextStepsLayout.svelte'
  import OrderByIdDeliveryTabCanceledContentNextStepList from '../OrderByIdDeliveryTabCanceledContentNextStepList/OrderByIdDeliveryTabCanceledContentNextStepList.svelte'
  import OrderByIdDeliveryTabCanceledContentStepsButtons from '../OrderByIdDeliveryTabCanceledContentStepsButtons/OrderByIdDeliveryTabCanceledContentStepsButtons.svelte'
</script>

<OrderByIdDeliveryTabCanceledContentNextStepsLayout>
  <Title text="Next Steps" />
  <OrderByIdDeliveryTabCanceledContentNextStepList />
  <OrderByIdDeliveryTabCanceledContentStepsButtons />
</OrderByIdDeliveryTabCanceledContentNextStepsLayout>
