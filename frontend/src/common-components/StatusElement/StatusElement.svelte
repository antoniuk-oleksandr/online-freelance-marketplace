<script lang="ts">
  import { StatusEnum } from '@/types/StatusEnum'
  import { getStatusBgColor } from '@/pages/order-submit-requirements/helpers'
  import { formatStatus } from './helpers'

  type StatusElementProps = {
    status: StatusEnum
  }
  const { status }: StatusElementProps = $props()
</script>

<span class="rounded-sm px-1 uppercase {getStatusBgColor(status)} text-white font-semibold"
  >{formatStatus(status)}</span
>
