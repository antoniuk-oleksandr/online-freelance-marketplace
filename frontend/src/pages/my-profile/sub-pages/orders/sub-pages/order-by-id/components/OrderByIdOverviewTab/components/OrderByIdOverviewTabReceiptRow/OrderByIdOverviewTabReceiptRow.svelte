<script lang="ts">
  import OrderByIdOverviewTabReceiptRowLayout from './OrderByIdOverviewTabReceiptRowLayout.svelte'

  type OrderByIdOverviewTabReceiptFooterRowProps = {
    firstText: string
    secondText: string
    styles?: string
  }

  const { firstText, secondText, styles }: OrderByIdOverviewTabReceiptFooterRowProps = $props()
</script>

<OrderByIdOverviewTabReceiptRowLayout {styles}>
  <span class="uppercase font-semibold">{firstText}</span>
  <span>{secondText}</span>
</OrderByIdOverviewTabReceiptRowLayout>
