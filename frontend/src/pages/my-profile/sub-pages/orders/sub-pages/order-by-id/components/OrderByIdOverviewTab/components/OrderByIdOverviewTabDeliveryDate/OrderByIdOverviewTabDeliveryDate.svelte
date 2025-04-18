<script lang="ts">
  import { formatFullOrderDate } from '@/pages/order-submit-requirements/helpers'

  type OrderByIdOverviewTabDeliveryDateProps = {
    deliveryDate: number
  }

  const { deliveryDate }: OrderByIdOverviewTabDeliveryDateProps = $props()
</script>

<div class="flex gap-1 flex-col md:flex-row">
  <span>Delivery date</span>
  <span class="font-semibold">{formatFullOrderDate(deliveryDate)}</span>
</div>
