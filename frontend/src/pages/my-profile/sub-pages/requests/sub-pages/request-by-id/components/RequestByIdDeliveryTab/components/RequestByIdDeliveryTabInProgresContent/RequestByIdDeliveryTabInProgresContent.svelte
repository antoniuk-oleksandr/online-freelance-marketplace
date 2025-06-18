<script lang="ts">
  import type { DeliveryData } from '@/types/DeliveryData'
  import RequestByIdDeliveryTabInProgresContentForm from '../RequestByIdDeliveryTabInProgresContentForm/RequestByIdDeliveryTabInProgresContentForm.svelte'
  import RequestByIdDeliveryTabInProgresContentFormContent from '../RequestByIdDeliveryTabInProgresContentFormContent/RequestByIdDeliveryTabInProgresContentFormContent.svelte'
  import { StatusEnum } from '@/types/StatusEnum'

  type RequestByIdDeliveryTabInProgresContentProps = {
    orderId: string
    data: DeliveryData
    setNewDeliveryData: (data: DeliveryData) => void
  }

  const { orderId, data, setNewDeliveryData }: RequestByIdDeliveryTabInProgresContentProps =
    $props()
</script>

{#if data.status === StatusEnum.InProgress}
  <RequestByIdDeliveryTabInProgresContentForm {setNewDeliveryData} {orderId}>
    <RequestByIdDeliveryTabInProgresContentFormContent />
  </RequestByIdDeliveryTabInProgresContentForm>
{/if}
