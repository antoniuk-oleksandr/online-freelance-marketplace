<script lang="ts">
  type OrderByIdOverviewTabOrderNumberProps = {
    id: number
  }

  const { id }: OrderByIdOverviewTabOrderNumberProps = $props()
</script>

<div>
  <span>Order number</span>
  <span class="font-semibold">#{id}</span>
</div>
