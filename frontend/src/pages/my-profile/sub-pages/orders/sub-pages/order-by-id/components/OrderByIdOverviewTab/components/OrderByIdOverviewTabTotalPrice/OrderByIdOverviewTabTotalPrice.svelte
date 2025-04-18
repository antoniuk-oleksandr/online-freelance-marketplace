<script lang="ts">
  import OrderByIdOverviewTabTotalPriceLayout from './OrderByIdOverviewTabTotalPriceLayout.svelte'

  type OrderByIdOverviewTabTotalPriceProps = {
    totalPrice: number
  }

  const { totalPrice }: OrderByIdOverviewTabTotalPriceProps = $props()
</script>

<OrderByIdOverviewTabTotalPriceLayout>
  <p>total price</p>
  <p class="text-xl">{totalPrice}$</p>
</OrderByIdOverviewTabTotalPriceLayout>
