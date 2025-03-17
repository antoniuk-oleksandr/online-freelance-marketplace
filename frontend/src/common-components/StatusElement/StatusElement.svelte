<script lang="ts">
  import { StatusEnum } from '@/types/StatusEnum'
  import { getStatusBgColor as getStatusAdditioanlStyles } from '@/pages/order-submit-requirements/helpers'
  import { formatStatus } from './helpers'

  type StatusElementProps = {
    status: StatusEnum
  }
  const { status }: StatusElementProps = $props()
</script>

<span class="px-3 py-1 text-sm font-semibold rounded-md border {getStatusAdditioanlStyles(status)}">
  {formatStatus(status)}
</span>
