<script lang="ts">
  import Icon from '@iconify/svelte'
  import OrderByIdDeliveryTabCanceledContentRefundLayout from './OrderByIdDeliveryTabCanceledContentRefundLayout.svelte'
  import type { DeliveryPayment } from '@/types/DeliveryPayment'

  type OrderByIdDeliveryTabCanceledContentRefundProps = {
    data: DeliveryPayment
  }

  const { data }: OrderByIdDeliveryTabCanceledContentRefundProps = $props()
</script>

<OrderByIdDeliveryTabCanceledContentRefundLayout>
  <Icon
    icon="heroicons:currency-dollar"
    height="64"
    width="64"
    class="w-8 h-8 text-blue-700 dark:text-blue-200 flex-shrink-0"
  />
  <div>
    <h3 class="text-blue-700 dark:text-blue-200 font-medium">Refund Status</h3>
    <p class="text-blue-600 dark:text-blue-300 mt-2">
      Your payment of <span class="font-semibold">{data.paidAmount}$</span> has been refunded to your
      original payment method.
    </p>
    <p class="text-blue-600 dark:text-blue-300 text-sm mt-2">
      Please allow 5-7 business days for the refund to appear in your account.
    </p>
  </div>
</OrderByIdDeliveryTabCanceledContentRefundLayout>
