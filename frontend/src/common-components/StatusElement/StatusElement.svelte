<script lang="ts">
  import { StatusEnum } from '@/types/StatusEnum'
  import { getStatusBgColor as getStatusAdditionalStyles } from '@/pages/order-submit-requirements/helpers'
  import { formatStatus } from './helpers'

  type StatusElementProps = {
    status: StatusEnum
    styles?: string
  }
  const { status, styles }: StatusElementProps = $props()
</script>

<span
  class="{styles} px-3 py-1 text-sm font-semibold rounded-md border {getStatusAdditionalStyles(
    status,
  )}"
>
  {formatStatus(status)}
</span>
